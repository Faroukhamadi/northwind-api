package api

import (
	"encoding/json"
	"net/http"
	// "github.com/Faroukhamadi/northwind-api/utils"
)

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "[ERROR]Unable to encode data", http.StatusInternalServerError)
		}
	}
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}
