package user

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/match-a-player/server/database"
	"github.com/rs/zerolog/log"
)

func Crud(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	switch r.Method {
	case http.MethodGet:
		users, err := getAllUsers(database.GlobalDB.ReadTimeout, &database.GlobalDB)
		if err != nil {
			log.Error().Msgf("Crud: %v", err)
			w.WriteHeader(500)
			return
		}
		body, err := json.Marshal(users)
		if err != nil {
			log.Error().Msgf("Crud: %v", err)
			w.WriteHeader(500)
			return
		}
		w.Write(body)
		return
	case http.MethodPost:
		// Create a new record.
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		// Give an error message.
	}
}
