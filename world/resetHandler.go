package world

import (
	"github.com/Thomas-Webber/webcraft-backend/security"
	"net/http"
)

// ResetHandler -
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	security.EnableCors(&w)
	World = make(map[int32]int32)
	World[EncodeXYZtoInt(15, 5, 5)] = 0xffffff
	World[EncodeXYZtoInt(5, 5, 5)] = 0xff0000
	w.WriteHeader(http.StatusNoContent)
}
