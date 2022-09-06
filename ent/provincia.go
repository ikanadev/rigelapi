// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/vmkevv/rigelapi/ent/dpto"
	"github.com/vmkevv/rigelapi/ent/provincia"
)

// Provincia is the model entity for the Provincia schema.
type Provincia struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProvinciaQuery when eager-loading is set.
	Edges           ProvinciaEdges `json:"edges"`
	dpto_provincias *string
}

// ProvinciaEdges holds the relations/edges for other nodes in the graph.
type ProvinciaEdges struct {
	// Municipios holds the value of the municipios edge.
	Municipios []*Municipio `json:"municipios,omitempty"`
	// Departamento holds the value of the departamento edge.
	Departamento *Dpto `json:"departamento,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MunicipiosOrErr returns the Municipios value or an error if the edge
// was not loaded in eager-loading.
func (e ProvinciaEdges) MunicipiosOrErr() ([]*Municipio, error) {
	if e.loadedTypes[0] {
		return e.Municipios, nil
	}
	return nil, &NotLoadedError{edge: "municipios"}
}

// DepartamentoOrErr returns the Departamento value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProvinciaEdges) DepartamentoOrErr() (*Dpto, error) {
	if e.loadedTypes[1] {
		if e.Departamento == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: dpto.Label}
		}
		return e.Departamento, nil
	}
	return nil, &NotLoadedError{edge: "departamento"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Provincia) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case provincia.FieldID, provincia.FieldName:
			values[i] = new(sql.NullString)
		case provincia.ForeignKeys[0]: // dpto_provincias
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Provincia", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Provincia fields.
func (pr *Provincia) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provincia.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case provincia.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case provincia.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dpto_provincias", values[i])
			} else if value.Valid {
				pr.dpto_provincias = new(string)
				*pr.dpto_provincias = value.String
			}
		}
	}
	return nil
}

// QueryMunicipios queries the "municipios" edge of the Provincia entity.
func (pr *Provincia) QueryMunicipios() *MunicipioQuery {
	return (&ProvinciaClient{config: pr.config}).QueryMunicipios(pr)
}

// QueryDepartamento queries the "departamento" edge of the Provincia entity.
func (pr *Provincia) QueryDepartamento() *DptoQuery {
	return (&ProvinciaClient{config: pr.config}).QueryDepartamento(pr)
}

// Update returns a builder for updating this Provincia.
// Note that you need to call Provincia.Unwrap() before calling this method if this Provincia
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Provincia) Update() *ProvinciaUpdateOne {
	return (&ProvinciaClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Provincia entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Provincia) Unwrap() *Provincia {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Provincia is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Provincia) String() string {
	var builder strings.Builder
	builder.WriteString("Provincia(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// ProvinciaSlice is a parsable slice of Provincia.
type ProvinciaSlice []*Provincia

func (pr ProvinciaSlice) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}