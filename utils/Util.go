package utils

import (
	"encoding/json"
	"net/http"
)

func Message(success bool, message string, error string) map[string]interface{} {
	return map[string]interface{}{"success": success, "message": message, "error": error}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
