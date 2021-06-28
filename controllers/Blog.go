package controllers

import (
	"net/http"
	"rnl360-api/entity"
	"rnl360-api/models"
	u "rnl360-api/utils"
)

var GetAllBlog = func(w http.ResponseWriter, r *http.Request) {
	var success bool
	var message string
	var error string
	var blogs []models.BlogModel
	err := entity.GetAllBlogFeature(&blogs)
	if err != nil {
		success = false
		error = err.Error()
		message = "Database Server error!"
	} else {
		if len(blogs) == 0 {
			success = true
			message = "Data not found!"
		} else {
			success = true
			message = "Data found!"
		}
	}

	resp := u.Message(success, message, error)
	resp["results"] = blogs
	u.Respond(w, resp)
}
