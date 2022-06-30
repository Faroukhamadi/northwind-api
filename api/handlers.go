package api

import (
	"context"
	"net/http"

	"github.com/Faroukhamadi/northwind-api/ent"
	models "github.com/Faroukhamadi/northwind-api/models"
	"github.com/gorilla/mux"
)

// GetAll godoc
// @Summary Get all countries details
// @Description get all details
// @ID get-countries-details
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Country
// @Failure 500 {string} string "Internal server error"
// @Router /countries [get]
func (s *Server) GetAll(w http.ResponseWriter, r *http.Request) {

	var c models.Country
	c.Code2 = "2"

	s.l.Println("[DEBUG] GET ALL METHOD")
	ctx := context.Background()

	countries, err := s.orm.Country.Query().All(ctx)
	if err != nil {
		s.l.Fatal("[ERROR]", err)
		http.Error(w, "failed querying countries from db", http.StatusInternalServerError)
	}
	s.respond(w, r, countries, http.StatusOK)
}

// GetOne godoc
// @Summary Get country details
// @Description get all details
// @ID get-country-details
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Country
// @Failure 500 {string} string "Internal server error"
// @Router /countries/{id} [get]
func (s *Server) GetOne(w http.ResponseWriter, r *http.Request) {
	s.l.Println("[DEBUG] GET ONE METHOD")
	vars := mux.Vars(r)

	ctx := context.Background()
	country, err := s.orm.Country.Get(ctx, vars["id"])
	if err != nil {
		s.l.Fatal("[ERROR]", err)
		http.Error(w, "failed querying country by id from db", http.StatusInternalServerError)
	}
	s.respond(w, r, country, http.StatusOK)
}

// UpdateOne godoc
// @Summary Update country based on parameter
// @Description Update country
// @ID update-country-details
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Country
// @Failure 500 {string} string "Internal server error"
// @Router /countries/{id} [put]
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
