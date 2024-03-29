// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/grade"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// GradeUpdate is the builder for updating Grade entities.
type GradeUpdate struct {
	config
	hooks    []Hook
	mutation *GradeMutation
}

// Where appends a list predicates to the GradeUpdate builder.
func (gu *GradeUpdate) Where(ps ...predicate.Grade) *GradeUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GradeUpdate) SetName(s string) *GradeUpdate {
	gu.mutation.SetName(s)
	return gu
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (gu *GradeUpdate) AddClassIDs(ids ...string) *GradeUpdate {
	gu.mutation.AddClassIDs(ids...)
	return gu
}

// AddClasses adds the "classes" edges to the Class entity.
func (gu *GradeUpdate) AddClasses(c ...*Class) *GradeUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return gu.AddClassIDs(ids...)
}

// Mutation returns the GradeMutation object of the builder.
func (gu *GradeUpdate) Mutation() *GradeMutation {
	return gu.mutation
}

// ClearClasses clears all "classes" edges to the Class entity.
func (gu *GradeUpdate) ClearClasses() *GradeUpdate {
	gu.mutation.ClearClasses()
	return gu
}

// RemoveClassIDs removes the "classes" edge to Class entities by IDs.
func (gu *GradeUpdate) RemoveClassIDs(ids ...string) *GradeUpdate {
	gu.mutation.RemoveClassIDs(ids...)
	return gu
}

// RemoveClasses removes "classes" edges to Class entities.
func (gu *GradeUpdate) RemoveClasses(c ...*Class) *GradeUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return gu.RemoveClassIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GradeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GradeUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GradeUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GradeUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GradeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(grade.Table, grade.Columns, sqlgraph.NewFieldSpec(grade.FieldID, field.TypeString))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(grade.FieldName, field.TypeString, value)
	}
	if gu.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   grade.ClassesTable,
			Columns: []string{grade.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedClassesIDs(); len(nodes) > 0 && !gu.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   grade.ClassesTable,
			Columns: []string{grade.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   grade.ClassesTable,
			Columns: []string{grade.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grade.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GradeUpdateOne is the builder for updating a single Grade entity.
type GradeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GradeMutation
}

// SetName sets the "name" field.
func (guo *GradeUpdateOne) SetName(s string) *GradeUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (guo *GradeUpdateOne) AddClassIDs(ids ...string) *GradeUpdateOne {
	guo.mutation.AddClassIDs(ids...)
	return guo
}

// AddClasses adds the "classes" edges to the Class entity.
func (guo *GradeUpdateOne) AddClasses(c ...*Class) *GradeUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return guo.AddClassIDs(ids...)
}

// Mutation returns the GradeMutation object of the builder.
func (guo *GradeUpdateOne) Mutation() *GradeMutation {
	return guo.mutation
}

// ClearClasses clears all "classes" edges to the Class entity.
func (guo *GradeUpdateOne) ClearClasses() *GradeUpdateOne {
	guo.mutation.ClearClasses()
	return guo
}

// RemoveClassIDs removes the "classes" edge to Class entities by IDs.
func (guo *GradeUpdateOne) RemoveClassIDs(ids ...string) *GradeUpdateOne {
	guo.mutation.RemoveClassIDs(ids...)
	return guo
}

// RemoveClasses removes "classes" edges to Class entities.
func (guo *GradeUpdateOne) RemoveClasses(c ...*Class) *GradeUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return guo.RemoveClassIDs(ids...)
}

// Where appends a list predicates to the GradeUpdate builder.
func (guo *GradeUpdateOne) Where(ps ...predicate.Grade) *GradeUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GradeUpdateOne) Select(field string, fields ...string) *GradeUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Grade entity.
func (guo *GradeUpdateOne) Save(ctx context.Context) (*Grade, error) {
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GradeUpdateOne) SaveX(ctx context.Context) *Grade {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GradeUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GradeUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GradeUpdateOne) sqlSave(ctx context.Context) (_node *Grade, err error) {
	_spec := sqlgraph.NewUpdateSpec(grade.Table, grade.Columns, sqlgraph.NewFieldSpec(grade.FieldID, field.TypeString))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Grade.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grade.FieldID)
		for _, f := range fields {
			if !grade.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != grade.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(grade.FieldName, field.TypeString, value)
	}
	if guo.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   grade.ClassesTable,
			Columns: []string{grade.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedClassesIDs(); len(nodes) > 0 && !guo.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   grade.ClassesTable,
			Columns: []string{grade.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   grade.ClassesTable,
			Columns: []string{grade.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Grade{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grade.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
