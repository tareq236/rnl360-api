package controllers

import (
	"net/http"
	"rnl360-api/entity"
	"rnl360-api/models"
	u "rnl360-api/utils"
)

var GetAllCommunication = func(w http.ResponseWriter, r *http.Request) {
	var success bool
	var message string
	var error string
	var communication []models.CommunicationModel
	err := entity.GetAllCommunication(&communication)
	if err != nil {
		success = false
		error = err.Error()
		message = "Database Server error!"
	} else {
		if len(communication) == 0 {
			success = false
			message = "Data not found!"
		} else {
			success = true
			message = "Data found!"
		}
	}

	resp := u.Message(success, message, error)
	resp["results"] = communication
	u.Respond(w, resp)
}
