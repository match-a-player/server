package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

func StartRouter() {
	router := httprouter.New()
	router.GET("/health", HealthHandler)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Panic().Msg(err.Error())
	}
}
