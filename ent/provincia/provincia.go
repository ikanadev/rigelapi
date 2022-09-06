// Code generated by ent, DO NOT EDIT.

package provincia

const (
	// Label holds the string label denoting the provincia type in the database.
	Label = "provincia"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeMunicipios holds the string denoting the municipios edge name in mutations.
	EdgeMunicipios = "municipios"
	// EdgeDepartamento holds the string denoting the departamento edge name in mutations.
	EdgeDepartamento = "departamento"
	// Table holds the table name of the provincia in the database.
	Table = "provincia"
	// MunicipiosTable is the table that holds the municipios relation/edge.
	MunicipiosTable = "municipios"
	// MunicipiosInverseTable is the table name for the Municipio entity.
	// It exists in this package in order to avoid circular dependency with the "municipio" package.
	MunicipiosInverseTable = "municipios"
	// MunicipiosColumn is the table column denoting the municipios relation/edge.
	MunicipiosColumn = "provincia_municipios"
	// DepartamentoTable is the table that holds the departamento relation/edge.
	DepartamentoTable = "provincia"
	// DepartamentoInverseTable is the table name for the Dpto entity.
	// It exists in this package in order to avoid circular dependency with the "dpto" package.
	DepartamentoInverseTable = "dptos"
	// DepartamentoColumn is the table column denoting the departamento relation/edge.
	DepartamentoColumn = "dpto_provincias"
)

// Columns holds all SQL columns for provincia fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "provincia"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"dpto_provincias",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}