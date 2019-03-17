package world

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetHandler -
func GetHandler(w http.ResponseWriter, r *http.Request) {
	jsonString, err := json.Marshal(World)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonString)
}
