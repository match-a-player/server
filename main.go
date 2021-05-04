package main

import (
	"github.com/match-a-player/server/config"
	"github.com/match-a-player/server/database"
	"github.com/match-a-player/server/web"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if err := config.InitialiseConfig(); err != nil {
		log.Panic().Msgf("Config could not be initialized due to %v", err)
	}
	log.Info().Msgf("Config is %v", config.AppConfig)
	if err := database.GlobalDB.InstantiateClient(config.AppConfig.DatabaseURI,
		config.AppConfig.DatabaseUserName, config.AppConfig.DatabasePassword, config.AppConfig.DatabaseName,
		config.AppConfig.DatabaseAuthSource,
		config.AppConfig.DatabaseReadTimeout, config.AppConfig.DatabaseWriteTimeout); err != nil {
		log.Panic().Msgf("Database could not be initialized due to %v", err)
	}
	if err := database.GlobalDB.CheckConnection(); err != nil {
		log.Panic().Msgf("Database could not be connected due to %v", err)
	}
	defer database.GlobalDB.DisconnectClient()
	web.StartRouter()
}
