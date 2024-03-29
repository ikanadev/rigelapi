// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/dpto"
	"github.com/vmkevv/rigelapi/ent/municipio"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/provincia"
)

// ProvinciaUpdate is the builder for updating Provincia entities.
type ProvinciaUpdate struct {
	config
	hooks    []Hook
	mutation *ProvinciaMutation
}

// Where appends a list predicates to the ProvinciaUpdate builder.
func (pu *ProvinciaUpdate) Where(ps ...predicate.Provincia) *ProvinciaUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProvinciaUpdate) SetName(s string) *ProvinciaUpdate {
	pu.mutation.SetName(s)
	return pu
}

// AddMunicipioIDs adds the "municipios" edge to the Municipio entity by IDs.
func (pu *ProvinciaUpdate) AddMunicipioIDs(ids ...string) *ProvinciaUpdate {
	pu.mutation.AddMunicipioIDs(ids...)
	return pu
}

// AddMunicipios adds the "municipios" edges to the Municipio entity.
func (pu *ProvinciaUpdate) AddMunicipios(m ...*Municipio) *ProvinciaUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pu.AddMunicipioIDs(ids...)
}

// SetDepartamentoID sets the "departamento" edge to the Dpto entity by ID.
func (pu *ProvinciaUpdate) SetDepartamentoID(id string) *ProvinciaUpdate {
	pu.mutation.SetDepartamentoID(id)
	return pu
}

// SetNillableDepartamentoID sets the "departamento" edge to the Dpto entity by ID if the given value is not nil.
func (pu *ProvinciaUpdate) SetNillableDepartamentoID(id *string) *ProvinciaUpdate {
	if id != nil {
		pu = pu.SetDepartamentoID(*id)
	}
	return pu
}

// SetDepartamento sets the "departamento" edge to the Dpto entity.
func (pu *ProvinciaUpdate) SetDepartamento(d *Dpto) *ProvinciaUpdate {
	return pu.SetDepartamentoID(d.ID)
}

// Mutation returns the ProvinciaMutation object of the builder.
func (pu *ProvinciaUpdate) Mutation() *ProvinciaMutation {
	return pu.mutation
}

// ClearMunicipios clears all "municipios" edges to the Municipio entity.
func (pu *ProvinciaUpdate) ClearMunicipios() *ProvinciaUpdate {
	pu.mutation.ClearMunicipios()
	return pu
}

// RemoveMunicipioIDs removes the "municipios" edge to Municipio entities by IDs.
func (pu *ProvinciaUpdate) RemoveMunicipioIDs(ids ...string) *ProvinciaUpdate {
	pu.mutation.RemoveMunicipioIDs(ids...)
	return pu
}

// RemoveMunicipios removes "municipios" edges to Municipio entities.
func (pu *ProvinciaUpdate) RemoveMunicipios(m ...*Municipio) *ProvinciaUpdate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pu.RemoveMunicipioIDs(ids...)
}

// ClearDepartamento clears the "departamento" edge to the Dpto entity.
func (pu *ProvinciaUpdate) ClearDepartamento() *ProvinciaUpdate {
	pu.mutation.ClearDepartamento()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProvinciaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProvinciaUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProvinciaUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProvinciaUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProvinciaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(provincia.Table, provincia.Columns, sqlgraph.NewFieldSpec(provincia.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(provincia.FieldName, field.TypeString, value)
	}
	if pu.mutation.MunicipiosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provincia.MunicipiosTable,
			Columns: []string{provincia.MunicipiosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(municipio.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedMunicipiosIDs(); len(nodes) > 0 && !pu.mutation.MunicipiosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provincia.MunicipiosTable,
			Columns: []string{provincia.MunicipiosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(municipio.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.MunicipiosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provincia.MunicipiosTable,
			Columns: []string{provincia.MunicipiosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(municipio.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.DepartamentoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provincia.DepartamentoTable,
			Columns: []string{provincia.DepartamentoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dpto.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DepartamentoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provincia.DepartamentoTable,
			Columns: []string{provincia.DepartamentoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dpto.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provincia.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProvinciaUpdateOne is the builder for updating a single Provincia entity.
type ProvinciaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProvinciaMutation
}

// SetName sets the "name" field.
func (puo *ProvinciaUpdateOne) SetName(s string) *ProvinciaUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// AddMunicipioIDs adds the "municipios" edge to the Municipio entity by IDs.
func (puo *ProvinciaUpdateOne) AddMunicipioIDs(ids ...string) *ProvinciaUpdateOne {
	puo.mutation.AddMunicipioIDs(ids...)
	return puo
}

// AddMunicipios adds the "municipios" edges to the Municipio entity.
func (puo *ProvinciaUpdateOne) AddMunicipios(m ...*Municipio) *ProvinciaUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return puo.AddMunicipioIDs(ids...)
}

// SetDepartamentoID sets the "departamento" edge to the Dpto entity by ID.
func (puo *ProvinciaUpdateOne) SetDepartamentoID(id string) *ProvinciaUpdateOne {
	puo.mutation.SetDepartamentoID(id)
	return puo
}

// SetNillableDepartamentoID sets the "departamento" edge to the Dpto entity by ID if the given value is not nil.
func (puo *ProvinciaUpdateOne) SetNillableDepartamentoID(id *string) *ProvinciaUpdateOne {
	if id != nil {
		puo = puo.SetDepartamentoID(*id)
	}
	return puo
}

// SetDepartamento sets the "departamento" edge to the Dpto entity.
func (puo *ProvinciaUpdateOne) SetDepartamento(d *Dpto) *ProvinciaUpdateOne {
	return puo.SetDepartamentoID(d.ID)
}

// Mutation returns the ProvinciaMutation object of the builder.
func (puo *ProvinciaUpdateOne) Mutation() *ProvinciaMutation {
	return puo.mutation
}

// ClearMunicipios clears all "municipios" edges to the Municipio entity.
func (puo *ProvinciaUpdateOne) ClearMunicipios() *ProvinciaUpdateOne {
	puo.mutation.ClearMunicipios()
	return puo
}

// RemoveMunicipioIDs removes the "municipios" edge to Municipio entities by IDs.
func (puo *ProvinciaUpdateOne) RemoveMunicipioIDs(ids ...string) *ProvinciaUpdateOne {
	puo.mutation.RemoveMunicipioIDs(ids...)
	return puo
}

// RemoveMunicipios removes "municipios" edges to Municipio entities.
func (puo *ProvinciaUpdateOne) RemoveMunicipios(m ...*Municipio) *ProvinciaUpdateOne {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return puo.RemoveMunicipioIDs(ids...)
}

// ClearDepartamento clears the "departamento" edge to the Dpto entity.
func (puo *ProvinciaUpdateOne) ClearDepartamento() *ProvinciaUpdateOne {
	puo.mutation.ClearDepartamento()
	return puo
}

// Where appends a list predicates to the ProvinciaUpdate builder.
func (puo *ProvinciaUpdateOne) Where(ps ...predicate.Provincia) *ProvinciaUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProvinciaUpdateOne) Select(field string, fields ...string) *ProvinciaUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Provincia entity.
func (puo *ProvinciaUpdateOne) Save(ctx context.Context) (*Provincia, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProvinciaUpdateOne) SaveX(ctx context.Context) *Provincia {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProvinciaUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProvinciaUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProvinciaUpdateOne) sqlSave(ctx context.Context) (_node *Provincia, err error) {
	_spec := sqlgraph.NewUpdateSpec(provincia.Table, provincia.Columns, sqlgraph.NewFieldSpec(provincia.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Provincia.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, provincia.FieldID)
		for _, f := range fields {
			if !provincia.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != provincia.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(provincia.FieldName, field.TypeString, value)
	}
	if puo.mutation.MunicipiosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provincia.MunicipiosTable,
			Columns: []string{provincia.MunicipiosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(municipio.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedMunicipiosIDs(); len(nodes) > 0 && !puo.mutation.MunicipiosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provincia.MunicipiosTable,
			Columns: []string{provincia.MunicipiosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(municipio.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.MunicipiosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   provincia.MunicipiosTable,
			Columns: []string{provincia.MunicipiosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(municipio.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.DepartamentoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provincia.DepartamentoTable,
			Columns: []string{provincia.DepartamentoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dpto.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DepartamentoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provincia.DepartamentoTable,
			Columns: []string{provincia.DepartamentoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dpto.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Provincia{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{provincia.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
