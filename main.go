package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("hello world")

	router := httprouter.New()
	router.GET("/health", HealthHandler)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Panic().Msg(err.Error())
	}
}
