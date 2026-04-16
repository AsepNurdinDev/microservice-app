package utils

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func Success(w http.ResponseWriter, message string) {
	JSON(w, map[string]string{
		"message": message,
	})
}

func Error(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	JSON(w, map[string]string{
		"error": message,
	})
}