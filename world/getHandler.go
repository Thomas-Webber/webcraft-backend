package world

import (
	"encoding/json"
	"github.com/Thomas-Webber/webcraft-backend/security"
	"log"
	"net/http"
)

// GetHandler -
func GetHandler(w http.ResponseWriter, r *http.Request) {
	security.EnableCors(&w)

	jsonString, err := json.Marshal(World)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonString)
}
