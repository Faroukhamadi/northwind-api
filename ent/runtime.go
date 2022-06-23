// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/Faroukhamadi/northwind-api/ent/country"
	"github.com/Faroukhamadi/northwind-api/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	countryFields := schema.Country{}.Fields()
	_ = countryFields
	// countryDescCode2 is the schema descriptor for code2 field.
	countryDescCode2 := countryFields[12].Descriptor()
	// country.Code2Validator is a validator for the "code2" field. It is called by the builders before save.
	country.Code2Validator = func() func(string) error {
		validators := countryDescCode2.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(code2 string) error {
			for _, fn := range fns {
				if err := fn(code2); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// countryDescID is the schema descriptor for id field.
	countryDescID := countryFields[0].Descriptor()
	// country.IDValidator is a validator for the "id" field. It is called by the builders before save.
	country.IDValidator = func() func(string) error {
		validators := countryDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}
