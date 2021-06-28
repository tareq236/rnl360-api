package entity

import (
	"os"
	DB "rnl360-api/database"
	"rnl360-api/models"
	u "rnl360-api/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Validate incoming user details...
func ValidateToken(account *models.TokenAccountModel) (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required", ""), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required", ""), false
	}

	//Email must be unique
	temp := &models.TokenAccountModel{}

	//check for errors and duplicate emails
	err := DB.GetDB().Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry", ""), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user.", ""), false
	}

	return u.Message(false, "Requirement passed", ""), true
}

func CreateToken(account *models.TokenAccountModel) map[string]interface{} {

	if resp, ok := ValidateToken(account); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	err := DB.GetDB().Create(account).Error
	if err != nil {
		return u.Message(false, "Failed to create account.", err.Error())
	}

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.", "")
	}

	//Create new JWT token for the newly registered account
	tk := &models.TokenModel{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.TokenModel = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created", "")
	response["account"] = account
	return response
}

func GetToken(email, password string) map[string]interface{} {

	account := &models.TokenAccountModel{}
	err := DB.GetDB().Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found", "")
		}
		return u.Message(false, "Connection error. Please retry", "")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again", "")
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &models.TokenModel{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.TokenModel = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In", "")
	resp["account"] = account
	return resp

}
