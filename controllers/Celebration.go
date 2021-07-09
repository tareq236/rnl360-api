package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rnl360-api/entity"
	"rnl360-api/models"
	u "rnl360-api/utils"
	"strconv"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

var GetAllCelebrationList = func(w http.ResponseWriter, r *http.Request) {

	celebration := &models.CelebrationModel{}
	err := json.NewDecoder(r.Body).Decode(celebration)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", err.Error()))
		return
	}

	if celebration.CelebrateStatus == "1" {
		var celebrationList []models.CelebrationModel
		err_list := entity.GetAllCelebrationListRM(&celebrationList, celebration.RequestWorkArea)
		if err_list != nil {
			u.Respond(w, u.Message(false, "Celebration list error !", err_list.Error()))
			return
		} else {
			if len(celebrationList) == 0 {
				u.Respond(w, u.Message(false, "Celebration list empty !", ""))
				return
			}
		}
		resp := u.Message(true, "Celebration list found!", "")
		resp["results"] = celebrationList
		u.Respond(w, resp)
	} else if celebration.CelebrateStatus == "2" {
		var celebrationList []models.CelebrationModel
		err_list := entity.GetAllCelebrationListWA(&celebrationList, celebration.WorkArea)
		if err_list != nil {
			u.Respond(w, u.Message(false, "Celebration list error !", err_list.Error()))
			return
		} else {
			if len(celebrationList) == 0 {
				u.Respond(w, u.Message(false, "Celebration list empty !", ""))
				return
			}
		}
		resp := u.Message(true, "Celebration list found!", "")
		resp["results"] = celebrationList
		u.Respond(w, resp)
	}

}

var DoNotCelebration = func(w http.ResponseWriter, r *http.Request) {

	celebration := &models.CelebrationModel{}
	err := json.NewDecoder(r.Body).Decode(celebration)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", err.Error()))
		return
	}
	celebration.CelebrationStatus = 0
	celebration.PermissionRequestDateTime = time.Now()

	resp := entity.SaveCelebration(celebration)
	u.Respond(w, resp)

}

var AskCelebration = func(w http.ResponseWriter, r *http.Request) {

	celebration := &models.CelebrationModel{}
	err := json.NewDecoder(r.Body).Decode(celebration)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", err.Error()))
		return
	}
	celebration.CelebrationType = 1
	celebration.CelebrationStatus = 1
	celebration.PermissionRequestDateTime = time.Now()
	celebration.PermissionResponseType = 1

	var userDetails []models.UserModel
	err_user := entity.CheckUser(&userDetails, celebration.RequestWorkArea)
	if err_user != nil {
		u.Respond(w, u.Message(false, "User information not found", err_user.Error()))
		return
	}
	title := "MIO send a request"
	details := "Please acccept this request"
	activity := "celebration_request_list"
	err_push := sendPushNotification(
		title,
		details,
		activity,
		userDetails[0], "")
	if err_push != nil {
		u.Respond(w, u.Message(false, "Push notification  error", err_push.Error()))
		return
	}

	resp := entity.SaveCelebration(celebration)
	u.Respond(w, resp)

}

var CheckCelebrationPermission = func(w http.ResponseWriter, r *http.Request) {

	celebration := &models.CelebrationModel{}
	err := json.NewDecoder(r.Body).Decode(celebration)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", err.Error()))
		return
	}
	if celebration.WorkArea == "" {
		u.Respond(w, u.Message(false, "SPA ID is empty", ""))
		return
	}
	if celebration.ChamberTerritoryID == "" {
		u.Respond(w, u.Message(false, "Chamber Territory ID is empty", ""))
		return
	}
	if celebration.DrChildID == "" {
		u.Respond(w, u.Message(false, "Dr Child ID is empty", ""))
		return
	}

	err_db := entity.CheckCelebrationPermission(celebration)
	if err_db != nil {
		u.Respond(w, u.Message(false, "Data not found!", err_db.Error()))
		return
	}

	resp := u.Message(true, "Data found!", "")
	resp["result"] = celebration
	u.Respond(w, resp)

}

var UpdateCelebration = func(w http.ResponseWriter, r *http.Request) {

	celebration := &models.CelebrationModel{}
	err := json.NewDecoder(r.Body).Decode(celebration)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", err.Error()))
		return
	}

	if celebration.ResponseTypeInput == "sms" {
		celebration.ResponseType = 1
		celebration.ResponseDateTime = time.Now()
	}

	err_update := entity.UpdateCelebration(celebration)
	if err_update != nil {
		u.Respond(w, u.Message(false, "Update error !", err_update.Error()))
		return
	}

	resp := u.Message(true, "Update successfully!", "")
	resp["result"] = celebration
	u.Respond(w, resp)

}

var UpdateCelebrationWithPhoto = func(w http.ResponseWriter, r *http.Request) {

	celebration := &models.CelebrationModel{}
	celebrationID, err_CID := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err_CID != nil {
		u.Respond(w, u.Message(false, "ID number not format.", err_CID.Error()))
		return
	}
	celebration.ID = uint(celebrationID)
	celebration.ResponseType = 1
	celebration.ResponseDateTime = time.Now()
	celebration.Feedback = r.FormValue("feedback")

	// LIMIT 10MB
	r.ParseMultipartForm(10 * 1024 * 1024)
	file, _, err := r.FormFile("image")
	switch err {
	case nil:
		defer file.Close()
		// Upload file
		tempFile, err2 := ioutil.TempFile("public/upload/celebration", "celebration-*.jpg")
		if err2 != nil {
			u.Respond(w, u.Message(false, "I request", err2.Error()))
			return
		}
		defer tempFile.Close()

		fileBytes, err3 := ioutil.ReadAll(file)
		if err3 != nil {
			u.Respond(w, u.Message(false, "I request", err3.Error()))
			return
		}
		tempFile.Write(fileBytes)
		fileName := strings.Split(tempFile.Name(), "/")
		celebration.Picture = fileName[3]
	case http.ErrMissingFile:
		log.Println("no file")
	default:
		defer file.Close()
		u.Respond(w, u.Message(false, "I request", err.Error()))
		return
	}

	err_update := entity.UpdateCelebration(celebration)
	if err_update != nil {
		u.Respond(w, u.Message(false, "Update error !", err_update.Error()))
		return
	}

	resp := u.Message(true, "Update successfully!", "")
	resp["result"] = celebration
	u.Respond(w, resp)

}

var GetAllCelebrationPermissionResponse = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	workArea := vars["work_area"]

	var celebrationList []models.CelebrationModel
	err_list := entity.GetAllCelebrationPermissionResponse(&celebrationList, workArea)
	if err_list != nil {
		u.Respond(w, u.Message(false, "Celebration list error !", err_list.Error()))
		return
	} else {
		if len(celebrationList) == 0 {
			u.Respond(w, u.Message(false, "Celebration list empty !", ""))
			return
		}
	}
	resp := u.Message(true, "Update successfully!", "")
	resp["results"] = celebrationList
	u.Respond(w, resp)

}

var GetAllCelebrationPendingRequest = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	workArea := vars["work_area"]

	var celebrationList []models.CelebrationModel
	err_list := entity.GetAllCelebrationPendingRequest(&celebrationList, workArea)
	if err_list != nil {
		u.Respond(w, u.Message(false, "Celebration list error !", err_list.Error()))
		return
	} else {
		if len(celebrationList) == 0 {
			u.Respond(w, u.Message(false, "Celebration list empty !", ""))
			return
		}
	}
	resp := u.Message(true, "Update successfully!", "")
	resp["results"] = celebrationList
	u.Respond(w, resp)

}

var PermissionResponseNotification = func(w http.ResponseWriter, r *http.Request) {

	celebration := &models.CelebrationModel{}
	err := json.NewDecoder(r.Body).Decode(celebration)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request", err.Error()))
		return
	}
	if entity.IsCelebrationChcekComplet(int(celebration.ID)) {
		u.Respond(w, u.Message(false, "Celebration completed by mio", "Celebration completed by mio"))
		return
	}

	if string(rune(celebration.ID)) == "" {
		u.Respond(w, u.Message(false, "Celebration ID is empty", ""))
		return
	}

	celebrationUpdate := &models.CelebrationModel{}
	celebrationUpdate.ID = celebration.ID

	if celebration.CelebrationStatus == 0 {
		celebrationUpdate.PermissionResponseDateTime = time.Now()
		celebrationUpdate.ResponseType = 1
	} else {

		if string(rune(celebration.PermissionStatus)) == "" {
			u.Respond(w, u.Message(false, "Permission status is empty", ""))
			return
		}
		if celebration.PermissionResponseTypeText == "" {
			u.Respond(w, u.Message(false, "Permission response type is empty", ""))
			return
		}

		celebrationUpdate.CelebrationStatus = 1
		celebrationUpdate.PermissionStatus = celebration.PermissionStatus

		responseType := &models.ResponseTypeModel{}
		err_rtn := entity.GetResponseTypeName(responseType, celebration.PermissionResponseTypeText)
		if err_rtn != nil {
			u.Respond(w, u.Message(false, "Response type not found", ""))
			return
		}
		celebrationUpdate.PermissionResponseType = responseType.ID

		celebrationUpdate.PermissionResponseDateTime = time.Now()
		celebrationUpdate.PermissionResponseText = celebration.PermissionResponseText
		if responseType.ID == 3 {
			celebrationUpdate.TextMessageID = 1
		} else if responseType.ID == 4 {
			celebrationUpdate.TextMessageID = 1
		} else if responseType.ID == 6 {
			celebrationUpdate.PermissionStatus = 3
			celebrationUpdate.ResponseType = 1
		}
	}

	err_update := entity.UpdateCelebration(celebrationUpdate)
	if err_update != nil {
		u.Respond(w, u.Message(false, "Update error !", err_update.Error()))
		return
	}
	var userDetails []models.UserModel
	err_user := entity.CheckUser(&userDetails, celebration.WorkArea)
	if err_user != nil {
		u.Respond(w, u.Message(false, "User information not found", err_user.Error()))
		return
	}
	celebrationDetails := &models.CelebrationModel{}
	celebrationDetails.ID = celebration.ID
	err_details := entity.CelebrationDetails(celebrationDetails)
	if err_user != nil {
		u.Respond(w, u.Message(false, "Celebration information not found", err_details.Error()))
		return
	}
	JSONout, err_json := json.Marshal(celebrationDetails)
	if err != nil {
		u.Respond(w, u.Message(false, "JSON convert error", err_json.Error()))
		return
	}

	title := "RM accept your request"
	details := "You get some gift and take a picture"
	activity := "celebration_details"
	err_push := sendPushNotification(
		title,
		details,
		activity,
		userDetails[0],
		string(JSONout))
	if err_push != nil {
		u.Respond(w, u.Message(false, "Push notification  error", err_push.Error()))
		return
	}

	resp := u.Message(true, "Push notification send successfully!", "")
	u.Respond(w, resp)

}

func sendPushNotification(title string, details string, activity string, userDetails models.UserModel, JSONDetails string) (err error) {
	opt := option.WithCredentialsFile("service_account_key.json")
	config := &firebase.Config{ProjectID: "rnl360-project"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return err
	}

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}
	registrationToken := userDetails.PushKey

	oneHour := time.Duration(1) * time.Hour
	message := &messaging.Message{
		Android: &messaging.AndroidConfig{
			TTL:      &oneHour,
			Priority: "normal",
			Notification: &messaging.AndroidNotification{
				Title: title,
				Body:  details,
				// Icon:  "stock_ticker_update",
				// Color: "#f45342",
			},
		},
		Data: map[string]string{
			"result":   JSONDetails,
			"activity": activity,
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		return err
	}
	fmt.Println("Successfully sent message:", response)
	return nil
}
