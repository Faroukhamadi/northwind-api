// Code generated by entc, DO NOT EDIT.

package country

const (
	// Label holds the string label denoting the country type in the database.
	Label = "country"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRegion holds the string denoting the region field in the database.
	FieldRegion = "region"
	// FieldSurfaceArea holds the string denoting the surface_area field in the database.
	FieldSurfaceArea = "surface_area"
	// FieldIndepYear holds the string denoting the indep_year field in the database.
	FieldIndepYear = "indep_year"
	// FieldPopulation holds the string denoting the population field in the database.
	FieldPopulation = "population"
	// FieldLifeExpectancy holds the string denoting the life_expectancy field in the database.
	FieldLifeExpectancy = "life_expectancy"
	// FieldGnp holds the string denoting the gnp field in the database.
	FieldGnp = "gnp"
	// FieldGnpOld holds the string denoting the gnp_old field in the database.
	FieldGnpOld = "gnp_old"
	// FieldLocalName holds the string denoting the local_name field in the database.
	FieldLocalName = "local_name"
	// FieldGovernmentForm holds the string denoting the government_form field in the database.
	FieldGovernmentForm = "government_form"
	// FieldHeadOfState holds the string denoting the head_of_state field in the database.
	FieldHeadOfState = "head_of_state"
	// FieldCode2 holds the string denoting the code2 field in the database.
	FieldCode2 = "code2"
	// EdgeCapital holds the string denoting the capital edge name in mutations.
	EdgeCapital = "capital"
	// EdgeLanguage holds the string denoting the language edge name in mutations.
	EdgeLanguage = "language"
	// CityFieldID holds the string denoting the ID field of the City.
	CityFieldID = "id"
	// Country_languageFieldID holds the string denoting the ID field of the Country_language.
	Country_languageFieldID = "id"
	// Table holds the table name of the country in the database.
	Table = "country"
	// CapitalTable is the table that holds the capital relation/edge.
	CapitalTable = "city"
	// CapitalInverseTable is the table name for the City entity.
	// It exists in this package in order to avoid circular dependency with the "city" package.
	CapitalInverseTable = "city"
	// CapitalColumn is the table column denoting the capital relation/edge.
	CapitalColumn = "country_capital"
	// LanguageTable is the table that holds the language relation/edge.
	LanguageTable = "country_language"
	// LanguageInverseTable is the table name for the Country_language entity.
	// It exists in this package in order to avoid circular dependency with the "country_language" package.
	LanguageInverseTable = "country_language"
	// LanguageColumn is the table column denoting the language relation/edge.
	LanguageColumn = "country_language"
)

// Columns holds all SQL columns for country fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldRegion,
	FieldSurfaceArea,
	FieldIndepYear,
	FieldPopulation,
	FieldLifeExpectancy,
	FieldGnp,
	FieldGnpOld,
	FieldLocalName,
	FieldGovernmentForm,
	FieldHeadOfState,
	FieldCode2,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// Code2Validator is a validator for the "code2" field. It is called by the builders before save.
	Code2Validator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)
