package controllers

import (
	"net/http"
	"rnl360-api/entity"
	"rnl360-api/models"
	u "rnl360-api/utils"

	"github.com/gorilla/mux"
)

var GetActivities = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	workArea := vars["work_area"]

	var todayActivitiesResult []entity.ActivitiesResult
	err_db := entity.GetTodayActivities(workArea, &todayActivitiesResult)
	if err_db != nil {
		u.Respond(w, u.Message(false, "Database Server error!", err_db.Error()))
		return
	}

	var upcomingActivitiesResult []entity.ActivitiesResult
	err_db1 := entity.GetUpcomingActivities(workArea, &upcomingActivitiesResult)
	if err_db1 != nil {
		u.Respond(w, u.Message(false, "Database Server error!", err_db1.Error()))
		return
	}

	var blogs []models.BlogModel
	err_db2 := entity.GetAllBlogFeature(&blogs)
	if err_db2 != nil {
		u.Respond(w, u.Message(false, "Database Server error!", err_db2.Error()))
		return
	}

	resp := u.Message(true, "Data found!", "")

	var todayActivitiesResultNew []entity.ActivitiesResult
	for _, element := range todayActivitiesResult {
		if !entity.IsCelebrationComplet(workArea, element.DrChildID) {
			todayActivitiesResultNew = append(todayActivitiesResultNew, element)
		}
	}
	var upcomingActivitiesResultNew []entity.ActivitiesResult
	for _, element := range upcomingActivitiesResult {
		if !entity.IsCelebrationComplet(workArea, element.DrChildID) {
			upcomingActivitiesResultNew = append(upcomingActivitiesResultNew, element)
		}
	}

	if len(blogs) == 0 {
		resp["blogs"] = make([]int, 0)
	} else {
		resp["blogs"] = blogs
	}
	if len(todayActivitiesResultNew) == 0 {
		resp["today"] = make([]int, 0)
	} else {
		resp["today"] = todayActivitiesResultNew
	}
	if len(upcomingActivitiesResultNew) == 0 {
		resp["upcoming"] = make([]int, 0)
	} else {
		resp["upcoming"] = upcomingActivitiesResultNew
	}
	u.Respond(w, resp)
}
