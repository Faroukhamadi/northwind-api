package api

import (
	"context"
	"log"
	"net/http"
)

func (s *Server) GetOne(w http.ResponseWriter, r *http.Request) {
	s.l.Println("[DEBUG] Get one method")
	ctx := context.Background()

	countries, err := s.orm.Country.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	s.l.Println("this is the response: ", countries[0])
	w.Write([]byte(countries[0].Name))
}
