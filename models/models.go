package models

type Country struct {
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Region holds the value of the "region" field.
	Region string `json:"region,omitempty"`
	// SurfaceArea holds the value of the "surface_area" field.
	SurfaceArea float64 `json:"surface_area,omitempty"`
	// IndepYear holds the value of the "indep_year" field.
	IndepYear int16 `json:"indep_year,omitempty"`
	// Population holds the value of the "population" field.
	Population int `json:"population,omitempty"`
	// LifeExpectancy holds the value of the "life_expectancy" field.
	LifeExpectancy float64 `json:"life_expectancy,omitempty"`
	// Gnp holds the value of the "gnp" field.
	Gnp float64 `json:"gnp,omitempty"`
	// GnpOld holds the value of the "gnp_old" field.
	GnpOld float64 `json:"gnp_old,omitempty"`
	// LocalName holds the value of the "local_name" field.
	LocalName string `json:"local_name,omitempty"`
	// GovernmentForm holds the value of the "government_form" field.
	GovernmentForm string `json:"government_form,omitempty"`
	// HeadOfState holds the value of the "head_of_state" field.
	HeadOfState *string `json:"head_of_state,omitempty"`
	// Code2 holds the value of the "code2" field.
	Code2 string `json:"code2,omitempty"`
}
