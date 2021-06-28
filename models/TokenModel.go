package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type TokenModel struct {
	UserId uint
	jwt.StandardClaims
}

type TokenAccountModel struct {
	gorm.Model
	Email      string `json:"email"`
	Password   string `json:"password"`
	TokenModel string `json:"token" sql:"-"`
}

func (b *TokenModel) TableName() string {
	return "token"
}

func (b *TokenAccountModel) TableName() string {
	return "token_account"
}
