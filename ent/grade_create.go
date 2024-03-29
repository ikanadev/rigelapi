// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/grade"
)

// GradeCreate is the builder for creating a Grade entity.
type GradeCreate struct {
	config
	mutation *GradeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (gc *GradeCreate) SetName(s string) *GradeCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetID sets the "id" field.
func (gc *GradeCreate) SetID(s string) *GradeCreate {
	gc.mutation.SetID(s)
	return gc
}

// AddClassIDs adds the "classes" edge to the Class entity by IDs.
func (gc *GradeCreate) AddClassIDs(ids ...string) *GradeCreate {
	gc.mutation.AddClassIDs(ids...)
	return gc
}

// AddClasses adds the "classes" edges to the Class entity.
func (gc *GradeCreate) AddClasses(c ...*Class) *GradeCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return gc.AddClassIDs(ids...)
}

// Mutation returns the GradeMutation object of the builder.
func (gc *GradeCreate) Mutation() *GradeMutation {
	return gc.mutation
}

// Save creates the Grade in the database.
func (gc *GradeCreate) Save(ctx context.Context) (*Grade, error) {
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GradeCreate) SaveX(ctx context.Context) *Grade {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GradeCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GradeCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GradeCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Grade.name"`)}
	}
	return nil
}

func (gc *GradeCreate) sqlSave(ctx context.Context) (*Grade, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Grade.ID type: %T", _spec.ID.Value)
		}
	}
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GradeCreate) createSpec() (*Grade, *sqlgraph.CreateSpec) {
	var (
		_node = &Grade{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(grade.Table, sqlgraph.NewFieldSpec(grade.FieldID, field.TypeString))
	)
	_spec.OnConflict = gc.conflict
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(grade.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := gc.mutation.ClassesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Grade.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GradeUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (gc *GradeCreate) OnConflict(opts ...sql.ConflictOption) *GradeUpsertOne {
	gc.conflict = opts
	return &GradeUpsertOne{
		create: gc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Grade.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gc *GradeCreate) OnConflictColumns(columns ...string) *GradeUpsertOne {
	gc.conflict = append(gc.conflict, sql.ConflictColumns(columns...))
	return &GradeUpsertOne{
		create: gc,
	}
}

type (
	// GradeUpsertOne is the builder for "upsert"-ing
	//  one Grade node.
	GradeUpsertOne struct {
		create *GradeCreate
	}

	// GradeUpsert is the "OnConflict" setter.
	GradeUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *GradeUpsert) SetName(v string) *GradeUpsert {
	u.Set(grade.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GradeUpsert) UpdateName() *GradeUpsert {
	u.SetExcluded(grade.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Grade.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(grade.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GradeUpsertOne) UpdateNewValues() *GradeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(grade.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Grade.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GradeUpsertOne) Ignore() *GradeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GradeUpsertOne) DoNothing() *GradeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GradeCreate.OnConflict
// documentation for more info.
func (u *GradeUpsertOne) Update(set func(*GradeUpsert)) *GradeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GradeUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *GradeUpsertOne) SetName(v string) *GradeUpsertOne {
	return u.Update(func(s *GradeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GradeUpsertOne) UpdateName() *GradeUpsertOne {
	return u.Update(func(s *GradeUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *GradeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GradeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GradeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GradeUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GradeUpsertOne.ID is not supported by MySQL driver. Use GradeUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GradeUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GradeCreateBulk is the builder for creating many Grade entities in bulk.
type GradeCreateBulk struct {
	config
	builders []*GradeCreate
	conflict []sql.ConflictOption
}

// Save creates the Grade entities in the database.
func (gcb *GradeCreateBulk) Save(ctx context.Context) ([]*Grade, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Grade, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GradeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GradeCreateBulk) SaveX(ctx context.Context) []*Grade {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GradeCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GradeCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Grade.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GradeUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (gcb *GradeCreateBulk) OnConflict(opts ...sql.ConflictOption) *GradeUpsertBulk {
	gcb.conflict = opts
	return &GradeUpsertBulk{
		create: gcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Grade.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gcb *GradeCreateBulk) OnConflictColumns(columns ...string) *GradeUpsertBulk {
	gcb.conflict = append(gcb.conflict, sql.ConflictColumns(columns...))
	return &GradeUpsertBulk{
		create: gcb,
	}
}

// GradeUpsertBulk is the builder for "upsert"-ing
// a bulk of Grade nodes.
type GradeUpsertBulk struct {
	create *GradeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Grade.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(grade.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GradeUpsertBulk) UpdateNewValues() *GradeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(grade.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Grade.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GradeUpsertBulk) Ignore() *GradeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GradeUpsertBulk) DoNothing() *GradeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GradeCreateBulk.OnConflict
// documentation for more info.
func (u *GradeUpsertBulk) Update(set func(*GradeUpsert)) *GradeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GradeUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *GradeUpsertBulk) SetName(v string) *GradeUpsertBulk {
	return u.Update(func(s *GradeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GradeUpsertBulk) UpdateName() *GradeUpsertBulk {
	return u.Update(func(s *GradeUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *GradeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GradeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GradeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GradeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
