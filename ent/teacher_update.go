// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/adminaction"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/teacher"
)

// TeacherUpdate is the builder for updating Teacher entities.
type TeacherUpdate struct {
	config
	hooks    []Hook
	mutation *TeacherMutation
}

// Where appends a list predicates to the TeacherUpdate builder.
func (tu *TeacherUpdate) Where(ps ...predicate.Teacher) *TeacherUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TeacherUpdate) SetName(s string) *TeacherUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetLastName sets the "last_name" field.
func (tu *TeacherUpdate) SetLastName(s string) *TeacherUpdate {
	tu.mutation.SetLastName(s)
	return tu
}

// SetEmail sets the "email" field.
func (tu *TeacherUpdate) SetEmail(s string) *TeacherUpdate {
	tu.mutation.SetEmail(s)
	return tu
}

// SetPassword sets the "password" field.
func (tu *TeacherUpdate) SetPassword(s string) *TeacherUpdate {
	tu.mutation.SetPassword(s)
	return tu
}

// SetIsAdmin sets the "is_admin" field.
func (tu *TeacherUpdate) SetIsAdmin(b bool) *TeacherUpdate {
	tu.mutation.SetIsAdmin(b)
	return tu
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableIsAdmin(b *bool) *TeacherUpdate {
	if b != nil {
		tu.SetIsAdmin(*b)
	}
	return tu
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (tu *TeacherUpdate) AddClassIDs(ids ...string) *TeacherUpdate {
	tu.mutation.AddClassIDs(ids...)
	return tu
}

// AddClasses adds the "classes" edges to the Class entity.
func (tu *TeacherUpdate) AddClasses(c ...*Class) *TeacherUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.AddClassIDs(ids...)
}

// AddActionIDs adds the "actions" edge to the AdminAction entity by IDs.
func (tu *TeacherUpdate) AddActionIDs(ids ...string) *TeacherUpdate {
	tu.mutation.AddActionIDs(ids...)
	return tu
}

// AddActions adds the "actions" edges to the AdminAction entity.
func (tu *TeacherUpdate) AddActions(a ...*AdminAction) *TeacherUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tu.AddActionIDs(ids...)
}

// Mutation returns the TeacherMutation object of the builder.
func (tu *TeacherUpdate) Mutation() *TeacherMutation {
	return tu.mutation
}

// ClearClasses clears all "classes" edges to the Class entity.
func (tu *TeacherUpdate) ClearClasses() *TeacherUpdate {
	tu.mutation.ClearClasses()
	return tu
}

// RemoveClassIDs removes the "classes" edge to Class entities by IDs.
func (tu *TeacherUpdate) RemoveClassIDs(ids ...string) *TeacherUpdate {
	tu.mutation.RemoveClassIDs(ids...)
	return tu
}

// RemoveClasses removes "classes" edges to Class entities.
func (tu *TeacherUpdate) RemoveClasses(c ...*Class) *TeacherUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.RemoveClassIDs(ids...)
}

// ClearActions clears all "actions" edges to the AdminAction entity.
func (tu *TeacherUpdate) ClearActions() *TeacherUpdate {
	tu.mutation.ClearActions()
	return tu
}

// RemoveActionIDs removes the "actions" edge to AdminAction entities by IDs.
func (tu *TeacherUpdate) RemoveActionIDs(ids ...string) *TeacherUpdate {
	tu.mutation.RemoveActionIDs(ids...)
	return tu
}

// RemoveActions removes "actions" edges to AdminAction entities.
func (tu *TeacherUpdate) RemoveActions(a ...*AdminAction) *TeacherUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tu.RemoveActionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeacherUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeacherMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeacherUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeacherUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeacherUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TeacherUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teacher.Table,
			Columns: teacher.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teacher.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldName,
		})
	}
	if value, ok := tu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldLastName,
		})
	}
	if value, ok := tu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldEmail,
		})
	}
	if value, ok := tu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldPassword,
		})
	}
	if value, ok := tu.mutation.IsAdmin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: teacher.FieldIsAdmin,
		})
	}
	if tu.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ClassesTable,
			Columns: []string{teacher.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedClassesIDs(); len(nodes) > 0 && !tu.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ClassesTable,
			Columns: []string{teacher.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ClassesTable,
			Columns: []string{teacher.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ActionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ActionsTable,
			Columns: []string{teacher.ActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: adminaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedActionsIDs(); len(nodes) > 0 && !tu.mutation.ActionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ActionsTable,
			Columns: []string{teacher.ActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: adminaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ActionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ActionsTable,
			Columns: []string{teacher.ActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: adminaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teacher.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TeacherUpdateOne is the builder for updating a single Teacher entity.
type TeacherUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeacherMutation
}

// SetName sets the "name" field.
func (tuo *TeacherUpdateOne) SetName(s string) *TeacherUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetLastName sets the "last_name" field.
func (tuo *TeacherUpdateOne) SetLastName(s string) *TeacherUpdateOne {
	tuo.mutation.SetLastName(s)
	return tuo
}

// SetEmail sets the "email" field.
func (tuo *TeacherUpdateOne) SetEmail(s string) *TeacherUpdateOne {
	tuo.mutation.SetEmail(s)
	return tuo
}

// SetPassword sets the "password" field.
func (tuo *TeacherUpdateOne) SetPassword(s string) *TeacherUpdateOne {
	tuo.mutation.SetPassword(s)
	return tuo
}

// SetIsAdmin sets the "is_admin" field.
func (tuo *TeacherUpdateOne) SetIsAdmin(b bool) *TeacherUpdateOne {
	tuo.mutation.SetIsAdmin(b)
	return tuo
}

// SetNillableIsAdmin sets the "is_admin" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableIsAdmin(b *bool) *TeacherUpdateOne {
	if b != nil {
		tuo.SetIsAdmin(*b)
	}
	return tuo
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (tuo *TeacherUpdateOne) AddClassIDs(ids ...string) *TeacherUpdateOne {
	tuo.mutation.AddClassIDs(ids...)
	return tuo
}

// AddClasses adds the "classes" edges to the Class entity.
func (tuo *TeacherUpdateOne) AddClasses(c ...*Class) *TeacherUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.AddClassIDs(ids...)
}

// AddActionIDs adds the "actions" edge to the AdminAction entity by IDs.
func (tuo *TeacherUpdateOne) AddActionIDs(ids ...string) *TeacherUpdateOne {
	tuo.mutation.AddActionIDs(ids...)
	return tuo
}

// AddActions adds the "actions" edges to the AdminAction entity.
func (tuo *TeacherUpdateOne) AddActions(a ...*AdminAction) *TeacherUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tuo.AddActionIDs(ids...)
}

// Mutation returns the TeacherMutation object of the builder.
func (tuo *TeacherUpdateOne) Mutation() *TeacherMutation {
	return tuo.mutation
}

// ClearClasses clears all "classes" edges to the Class entity.
func (tuo *TeacherUpdateOne) ClearClasses() *TeacherUpdateOne {
	tuo.mutation.ClearClasses()
	return tuo
}

// RemoveClassIDs removes the "classes" edge to Class entities by IDs.
func (tuo *TeacherUpdateOne) RemoveClassIDs(ids ...string) *TeacherUpdateOne {
	tuo.mutation.RemoveClassIDs(ids...)
	return tuo
}

// RemoveClasses removes "classes" edges to Class entities.
func (tuo *TeacherUpdateOne) RemoveClasses(c ...*Class) *TeacherUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.RemoveClassIDs(ids...)
}

// ClearActions clears all "actions" edges to the AdminAction entity.
func (tuo *TeacherUpdateOne) ClearActions() *TeacherUpdateOne {
	tuo.mutation.ClearActions()
	return tuo
}

// RemoveActionIDs removes the "actions" edge to AdminAction entities by IDs.
func (tuo *TeacherUpdateOne) RemoveActionIDs(ids ...string) *TeacherUpdateOne {
	tuo.mutation.RemoveActionIDs(ids...)
	return tuo
}

// RemoveActions removes "actions" edges to AdminAction entities.
func (tuo *TeacherUpdateOne) RemoveActions(a ...*AdminAction) *TeacherUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return tuo.RemoveActionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeacherUpdateOne) Select(field string, fields ...string) *TeacherUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Teacher entity.
func (tuo *TeacherUpdateOne) Save(ctx context.Context) (*Teacher, error) {
	var (
		err  error
		node *Teacher
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TeacherMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Teacher)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TeacherMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeacherUpdateOne) SaveX(ctx context.Context) *Teacher {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeacherUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeacherUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TeacherUpdateOne) sqlSave(ctx context.Context) (_node *Teacher, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teacher.Table,
			Columns: teacher.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: teacher.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Teacher.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teacher.FieldID)
		for _, f := range fields {
			if !teacher.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != teacher.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldName,
		})
	}
	if value, ok := tuo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldLastName,
		})
	}
	if value, ok := tuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldEmail,
		})
	}
	if value, ok := tuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: teacher.FieldPassword,
		})
	}
	if value, ok := tuo.mutation.IsAdmin(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: teacher.FieldIsAdmin,
		})
	}
	if tuo.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ClassesTable,
			Columns: []string{teacher.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedClassesIDs(); len(nodes) > 0 && !tuo.mutation.ClassesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ClassesTable,
			Columns: []string{teacher.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ClassesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ClassesTable,
			Columns: []string{teacher.ClassesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: class.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ActionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ActionsTable,
			Columns: []string{teacher.ActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: adminaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedActionsIDs(); len(nodes) > 0 && !tuo.mutation.ActionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ActionsTable,
			Columns: []string{teacher.ActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: adminaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ActionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.ActionsTable,
			Columns: []string{teacher.ActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: adminaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Teacher{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teacher.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
