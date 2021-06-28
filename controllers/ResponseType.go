package controllers

import (
	"net/http"
	"rnl360-api/entity"
	"rnl360-api/models"
	u "rnl360-api/utils"
)

var GetAllResponseType = func(w http.ResponseWriter, r *http.Request) {

	var rsponseTypeList []models.ResponseTypeModel
	err_list := entity.GetAllResponseType(&rsponseTypeList)
	if err_list != nil {
		u.Respond(w, u.Message(false, "Rsponse type list error !", err_list.Error()))
		return
	} else {
		if len(rsponseTypeList) == 0 {
			u.Respond(w, u.Message(false, "Rsponse type list empty !", ""))
			return
		}
	}
	resp := u.Message(true, "Rsponse type list found!", "")
	resp["results"] = rsponseTypeList
	u.Respond(w, resp)

}
