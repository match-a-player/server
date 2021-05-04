package web

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/match-a-player/server/config"
)

func HealthHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Ok"))
}

func authMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if strings.Compare(r.Header.Get("Authorization"), config.AppConfig.Token) != 0 {
			w.WriteHeader(401)
			return
		}
		next(w, r, ps)
	}
}
