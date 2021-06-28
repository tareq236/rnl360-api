package controllers

import (
	"encoding/json"
	"net/http"
	"rnl360-api/entity"
	"rnl360-api/models"
	u "rnl360-api/utils"
)

var CreateTokenAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.TokenAccountModel{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", ""))
		return
	}

	resp := entity.CreateToken(account) //Create account
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.TokenAccountModel{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", ""))
		return
	}

	resp := entity.GetToken(account.Email, account.Password)
	u.Respond(w, resp)

}
