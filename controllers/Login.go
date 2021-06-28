package controllers

import (
	"encoding/json"
	"net/http"
	"rnl360-api/entity"
	"rnl360-api/models"
	u "rnl360-api/utils"
)

var Login = func(w http.ResponseWriter, r *http.Request) {

	user := &models.LoginModel{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", err.Error()))
		return
	}

	if user.WorkArea == "" {
		u.Respond(w, u.Message(false, "Please enter SPA ID", ""))
		return
	}

	if user.Password == "" {
		u.Respond(w, u.Message(false, "Please enter password", ""))
		return
	}

	var userResult []entity.LoginResult
	err_db := entity.Login(user.WorkArea, user.Password, &userResult)
	if err_db != nil {
		u.Respond(w, u.Message(false, "Database Server error!", err_db.Error()))
		return
	} else {
		if len(userResult) == 0 {
			u.Respond(w, u.Message(false, "Please check your SAP id and password", ""))
			return
		} else {

			var checkUser []models.UserModel
			err_user := entity.CheckUser(&checkUser, userResult[0].WorkArea)
			if err_user != nil {
				u.Respond(w, u.Message(false, "User information not save", err_user.Error()))
				return
			} else {
				if len(checkUser) == 0 {
					userSave := &models.UserModel{}
					userSave.WorkArea = userResult[0].WorkArea
					userSave.UserName = userResult[0].UserName
					userSave.UserDesignationID = userResult[0].UserDesignationID
					userSave.Team = userResult[0].Team
					userSave.PushKey = user.PushKey
					userSave.DeviceType = user.DeviceType
					err_user_save := entity.SaveUser(userSave)
					if err_user_save != nil {
						u.Respond(w, u.Message(false, "User save error !", err_user_save.Error()))
						return
					}
				} else {
					userSave := &models.UserModel{}
					userSave.WorkArea = userResult[0].WorkArea
					userSave.UserName = userResult[0].UserName
					userSave.UserDesignationID = userResult[0].UserDesignationID
					userSave.Team = userResult[0].Team
					userSave.PushKey = user.PushKey
					userSave.DeviceType = user.DeviceType
					err_user_update := entity.UpdateUser(userSave)
					if err_user_update != nil {
						u.Respond(w, u.Message(false, "User update error !", err_user_update.Error()))
						return
					}
				}
			}

			resp := u.Message(true, "Data found!", "")
			resp["result"] = userResult[0]
			u.Respond(w, resp)
		}
	}
}
