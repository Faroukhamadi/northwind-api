// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/Faroukhamadi/northwind-api/ent/country"
	"github.com/Faroukhamadi/northwind-api/ent/country_language"
)

// Country_language is the model entity for the Country_language schema.
type Country_language struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CountryCode holds the value of the "country_code" field.
	CountryCode string `json:"country_code,omitempty"`
	// Language holds the value of the "language" field.
	Language string `json:"language,omitempty"`
	// IsOfficial holds the value of the "is_official" field.
	IsOfficial bool `json:"is_official,omitempty"`
	// Percentage holds the value of the "percentage" field.
	Percentage float64 `json:"percentage,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Country_languageQuery when eager-loading is set.
	Edges            Country_languageEdges `json:"edges"`
	country_language *string
}

// Country_languageEdges holds the relations/edges for other nodes in the graph.
type Country_languageEdges struct {
	// Country holds the value of the country edge.
	Country *Country `json:"country,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CountryOrErr returns the Country value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Country_languageEdges) CountryOrErr() (*Country, error) {
	if e.loadedTypes[0] {
		if e.Country == nil {
			// The edge country was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: country.Label}
		}
		return e.Country, nil
	}
	return nil, &NotLoadedError{edge: "country"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Country_language) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case country_language.FieldIsOfficial:
			values[i] = new(sql.NullBool)
		case country_language.FieldPercentage:
			values[i] = new(sql.NullFloat64)
		case country_language.FieldID:
			values[i] = new(sql.NullInt64)
		case country_language.FieldCountryCode, country_language.FieldLanguage:
			values[i] = new(sql.NullString)
		case country_language.ForeignKeys[0]: // country_language
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Country_language", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Country_language fields.
func (cl *Country_language) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case country_language.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cl.ID = int(value.Int64)
		case country_language.FieldCountryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_code", values[i])
			} else if value.Valid {
				cl.CountryCode = value.String
			}
		case country_language.FieldLanguage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language", values[i])
			} else if value.Valid {
				cl.Language = value.String
			}
		case country_language.FieldIsOfficial:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_official", values[i])
			} else if value.Valid {
				cl.IsOfficial = value.Bool
			}
		case country_language.FieldPercentage:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field percentage", values[i])
			} else if value.Valid {
				cl.Percentage = value.Float64
			}
		case country_language.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_language", values[i])
			} else if value.Valid {
				cl.country_language = new(string)
				*cl.country_language = value.String
			}
		}
	}
	return nil
}

// QueryCountry queries the "country" edge of the Country_language entity.
func (cl *Country_language) QueryCountry() *CountryQuery {
	return (&Country_languageClient{config: cl.config}).QueryCountry(cl)
}

// Update returns a builder for updating this Country_language.
// Note that you need to call Country_language.Unwrap() before calling this method if this Country_language
// was returned from a transaction, and the transaction was committed or rolled back.
func (cl *Country_language) Update() *CountryLanguageUpdateOne {
	return (&Country_languageClient{config: cl.config}).UpdateOne(cl)
}

// Unwrap unwraps the Country_language entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cl *Country_language) Unwrap() *Country_language {
	tx, ok := cl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Country_language is not a transactional entity")
	}
	cl.config.driver = tx.drv
	return cl
}

// String implements the fmt.Stringer.
func (cl *Country_language) String() string {
	var builder strings.Builder
	builder.WriteString("Country_language(")
	builder.WriteString(fmt.Sprintf("id=%v", cl.ID))
	builder.WriteString(", country_code=")
	builder.WriteString(cl.CountryCode)
	builder.WriteString(", language=")
	builder.WriteString(cl.Language)
	builder.WriteString(", is_official=")
	builder.WriteString(fmt.Sprintf("%v", cl.IsOfficial))
	builder.WriteString(", percentage=")
	builder.WriteString(fmt.Sprintf("%v", cl.Percentage))
	builder.WriteByte(')')
	return builder.String()
}

// Country_languages is a parsable slice of Country_language.
type Country_languages []*Country_language

func (cl Country_languages) config(cfg config) {
	for _i := range cl {
		cl[_i].config = cfg
	}
}