package api

import (
	"context"
	"net/http"

	"github.com/Faroukhamadi/northwind-api/ent"
	"github.com/gorilla/mux"
)

func (s *Server) GetAll(w http.ResponseWriter, r *http.Request) {
	s.l.Println("[DEBUG] GET ALL METHOD")
	ctx := context.Background()

	countries, err := s.orm.Country.Query().All(ctx)
	if err != nil {
		s.l.Fatal("[ERROR]", err)
	}
	s.respond(w, r, countries, http.StatusOK)
}

func (s *Server) GetOne(w http.ResponseWriter, r *http.Request) {
	s.l.Println("[DEBUG] GET ONE METHOD")
	vars := mux.Vars(r)

	ctx := context.Background()
	country, err := s.orm.Country.Get(ctx, vars["id"])
	if err != nil {
		s.l.Fatal("[ERROR]", err)
	}
	s.respond(w, r, country, http.StatusOK)
}

func (s *Server) UpdateOne(w http.ResponseWriter, r *http.Request) {
	s.l.Println("[DEBUG] UPDATE ONE METHOD")
	ctx := context.Background()
	vars := mux.Vars(r)

	var c ent.Country

	err := s.decode(w, r, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	updatedCountry, err := s.orm.Country.UpdateOneID(vars["id"]).SetPopulation(c.Population).Save(ctx)
	s.l.Println("updated country:", updatedCountry)

	if err != nil {
		s.l.Fatal("[ERROR]", err)
	}
}
