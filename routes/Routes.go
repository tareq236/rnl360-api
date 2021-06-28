package routes

import (
	"net/http"
	"rnl360-api/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	router_version_1 := router.PathPrefix("/api/v1").Subrouter()

	router_version_1.HandleFunc("/create_token_account", controllers.CreateTokenAccount).Methods("POST")
	router_version_1.HandleFunc("/get_token", controllers.Authenticate).Methods("POST")

	router_version_1.HandleFunc("/login", controllers.Login).Methods("POST")
	router_version_1.HandleFunc("/activity/{work_area}", controllers.GetActivities).Methods("GET")
	router_version_1.HandleFunc("/communication", controllers.GetAllCommunication).Methods("GET")

	router_version_1.HandleFunc("/response_type", controllers.GetAllResponseType).Methods("GET")

	router_version_1.HandleFunc("/donot_celebration", controllers.DoNotCelebration).Methods("POST")
	router_version_1.HandleFunc("/ask_celebration", controllers.AskCelebration).Methods("POST")
	router_version_1.HandleFunc("/check_celebration_permission", controllers.CheckCelebrationPermission).Methods("POST")
	router_version_1.HandleFunc("/update_celebration", controllers.UpdateCelebration).Methods("POST")
	router_version_1.HandleFunc("/update_celebration_with_photo", controllers.UpdateCelebrationWithPhoto).Methods("POST")
	router_version_1.HandleFunc("/celebration_list_pending_request/{work_area}", controllers.GetAllCelebrationPendingRequest).Methods("GET")
	router_version_1.HandleFunc("/celebration_list_permission_response/{work_area}", controllers.GetAllCelebrationPermissionResponse).Methods("GET")

	router_version_1.HandleFunc("/permission_response_notification", controllers.PermissionResponseNotification).Methods("POST")

	router_version_1.PathPrefix("/images").Handler(http.StripPrefix("/api/v1/images", http.FileServer(http.Dir("./public/images"))))

	// router_version_1.Use(auth.JwtAuthentication) //attach JWT auth middleware

	return router_version_1
}
