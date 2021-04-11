package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HealthHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Ok"))
}
