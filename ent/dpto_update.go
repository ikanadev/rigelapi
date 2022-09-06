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
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/provincia"
)

// DptoUpdate is the builder for updating Dpto entities.
type DptoUpdate struct {
	config
	hooks    []Hook
	mutation *DptoMutation
}

// Where appends a list predicates to the DptoUpdate builder.
func (du *DptoUpdate) Where(ps ...predicate.Dpto) *DptoUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetName sets the "name" field.
func (du *DptoUpdate) SetName(s string) *DptoUpdate {
	du.mutation.SetName(s)
	return du
}

// AddProvinciaIDs adds the "provincias" edge to the Provincia entity by IDs.
func (du *DptoUpdate) AddProvinciaIDs(ids ...string) *DptoUpdate {
	du.mutation.AddProvinciaIDs(ids...)
	return du
}

// AddProvincias adds the "provincias" edges to the Provincia entity.
func (du *DptoUpdate) AddProvincias(p ...*Provincia) *DptoUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.AddProvinciaIDs(ids...)
}

// Mutation returns the DptoMutation object of the builder.
func (du *DptoUpdate) Mutation() *DptoMutation {
	return du.mutation
}

// ClearProvincias clears all "provincias" edges to the Provincia entity.
func (du *DptoUpdate) ClearProvincias() *DptoUpdate {
	du.mutation.ClearProvincias()
	return du
}

// RemoveProvinciaIDs removes the "provincias" edge to Provincia entities by IDs.
func (du *DptoUpdate) RemoveProvinciaIDs(ids ...string) *DptoUpdate {
	du.mutation.RemoveProvinciaIDs(ids...)
	return du
}

// RemoveProvincias removes "provincias" edges to Provincia entities.
func (du *DptoUpdate) RemoveProvincias(p ...*Provincia) *DptoUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.RemoveProvinciaIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DptoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DptoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DptoUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DptoUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DptoUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DptoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dpto.Table,
			Columns: dpto.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: dpto.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dpto.FieldName,
		})
	}
	if du.mutation.ProvinciasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dpto.ProvinciasTable,
			Columns: []string{dpto.ProvinciasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: provincia.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedProvinciasIDs(); len(nodes) > 0 && !du.mutation.ProvinciasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dpto.ProvinciasTable,
			Columns: []string{dpto.ProvinciasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: provincia.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ProvinciasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dpto.ProvinciasTable,
			Columns: []string{dpto.ProvinciasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: provincia.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dpto.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DptoUpdateOne is the builder for updating a single Dpto entity.
type DptoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DptoMutation
}

// SetName sets the "name" field.
func (duo *DptoUpdateOne) SetName(s string) *DptoUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// AddProvinciaIDs adds the "provincias" edge to the Provincia entity by IDs.
func (duo *DptoUpdateOne) AddProvinciaIDs(ids ...string) *DptoUpdateOne {
	duo.mutation.AddProvinciaIDs(ids...)
	return duo
}

// AddProvincias adds the "provincias" edges to the Provincia entity.
func (duo *DptoUpdateOne) AddProvincias(p ...*Provincia) *DptoUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.AddProvinciaIDs(ids...)
}

// Mutation returns the DptoMutation object of the builder.
func (duo *DptoUpdateOne) Mutation() *DptoMutation {
	return duo.mutation
}

// ClearProvincias clears all "provincias" edges to the Provincia entity.
func (duo *DptoUpdateOne) ClearProvincias() *DptoUpdateOne {
	duo.mutation.ClearProvincias()
	return duo
}

// RemoveProvinciaIDs removes the "provincias" edge to Provincia entities by IDs.
func (duo *DptoUpdateOne) RemoveProvinciaIDs(ids ...string) *DptoUpdateOne {
	duo.mutation.RemoveProvinciaIDs(ids...)
	return duo
}

// RemoveProvincias removes "provincias" edges to Provincia entities.
func (duo *DptoUpdateOne) RemoveProvincias(p ...*Provincia) *DptoUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.RemoveProvinciaIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DptoUpdateOne) Select(field string, fields ...string) *DptoUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dpto entity.
func (duo *DptoUpdateOne) Save(ctx context.Context) (*Dpto, error) {
	var (
		err  error
		node *Dpto
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DptoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Dpto)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DptoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DptoUpdateOne) SaveX(ctx context.Context) *Dpto {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DptoUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DptoUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DptoUpdateOne) sqlSave(ctx context.Context) (_node *Dpto, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dpto.Table,
			Columns: dpto.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: dpto.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Dpto.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dpto.FieldID)
		for _, f := range fields {
			if !dpto.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dpto.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dpto.FieldName,
		})
	}
	if duo.mutation.ProvinciasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dpto.ProvinciasTable,
			Columns: []string{dpto.ProvinciasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: provincia.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedProvinciasIDs(); len(nodes) > 0 && !duo.mutation.ProvinciasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dpto.ProvinciasTable,
			Columns: []string{dpto.ProvinciasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: provincia.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ProvinciasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dpto.ProvinciasTable,
			Columns: []string{dpto.ProvinciasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: provincia.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Dpto{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dpto.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}