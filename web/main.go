package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/match-a-player/server/user"
	"github.com/rs/zerolog/log"
)

func StartRouter() {
	router := httprouter.New()
	router.GET("/health", HealthHandler)
	router.GET("/v1/user", authMiddleware(user.Crud))
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Panic().Msg(err.Error())
	}
}
