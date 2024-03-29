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
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/score"
	"github.com/vmkevv/rigelapi/ent/student"
)

// StudentCreate is the builder for creating a Student entity.
type StudentCreate struct {
	config
	mutation *StudentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (sc *StudentCreate) SetName(s string) *StudentCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetLastName sets the "last_name" field.
func (sc *StudentCreate) SetLastName(s string) *StudentCreate {
	sc.mutation.SetLastName(s)
	return sc
}

// SetCi sets the "ci" field.
func (sc *StudentCreate) SetCi(s string) *StudentCreate {
	sc.mutation.SetCi(s)
	return sc
}

// SetID sets the "id" field.
func (sc *StudentCreate) SetID(s string) *StudentCreate {
	sc.mutation.SetID(s)
	return sc
}

// AddAttendanceIDs adds the "attendances" edge to the Attendance entity by IDs.
func (sc *StudentCreate) AddAttendanceIDs(ids ...string) *StudentCreate {
	sc.mutation.AddAttendanceIDs(ids...)
	return sc
}

// AddAttendances adds the "attendances" edges to the Attendance entity.
func (sc *StudentCreate) AddAttendances(a ...*Attendance) *StudentCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddAttendanceIDs(ids...)
}

// AddScoreIDs adds the "scores" edge to the Score entity by IDs.
func (sc *StudentCreate) AddScoreIDs(ids ...string) *StudentCreate {
	sc.mutation.AddScoreIDs(ids...)
	return sc
}

// AddScores adds the "scores" edges to the Score entity.
func (sc *StudentCreate) AddScores(s ...*Score) *StudentCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddScoreIDs(ids...)
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (sc *StudentCreate) SetClassID(id string) *StudentCreate {
	sc.mutation.SetClassID(id)
	return sc
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (sc *StudentCreate) SetNillableClassID(id *string) *StudentCreate {
	if id != nil {
		sc = sc.SetClassID(*id)
	}
	return sc
}

// SetClass sets the "class" edge to the Class entity.
func (sc *StudentCreate) SetClass(c *Class) *StudentCreate {
	return sc.SetClassID(c.ID)
}

// Mutation returns the StudentMutation object of the builder.
func (sc *StudentCreate) Mutation() *StudentMutation {
	return sc.mutation
}

// Save creates the Student in the database.
func (sc *StudentCreate) Save(ctx context.Context) (*Student, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StudentCreate) SaveX(ctx context.Context) *Student {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StudentCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StudentCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StudentCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Student.name"`)}
	}
	if _, ok := sc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Student.last_name"`)}
	}
	if _, ok := sc.mutation.Ci(); !ok {
		return &ValidationError{Name: "ci", err: errors.New(`ent: missing required field "Student.ci"`)}
	}
	return nil
}

func (sc *StudentCreate) sqlSave(ctx context.Context) (*Student, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Student.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StudentCreate) createSpec() (*Student, *sqlgraph.CreateSpec) {
	var (
		_node = &Student{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(student.Table, sqlgraph.NewFieldSpec(student.FieldID, field.TypeString))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(student.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.LastName(); ok {
		_spec.SetField(student.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := sc.mutation.Ci(); ok {
		_spec.SetField(student.FieldCi, field.TypeString, value)
		_node.Ci = value
	}
	if nodes := sc.mutation.AttendancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.AttendancesTable,
			Columns: []string{student.AttendancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendance.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ScoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   student.ScoresTable,
			Columns: []string{student.ScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(score.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   student.ClassTable,
			Columns: []string{student.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.class_students = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Student.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StudentUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (sc *StudentCreate) OnConflict(opts ...sql.ConflictOption) *StudentUpsertOne {
	sc.conflict = opts
	return &StudentUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Student.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *StudentCreate) OnConflictColumns(columns ...string) *StudentUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &StudentUpsertOne{
		create: sc,
	}
}

type (
	// StudentUpsertOne is the builder for "upsert"-ing
	//  one Student node.
	StudentUpsertOne struct {
		create *StudentCreate
	}

	// StudentUpsert is the "OnConflict" setter.
	StudentUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *StudentUpsert) SetName(v string) *StudentUpsert {
	u.Set(student.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StudentUpsert) UpdateName() *StudentUpsert {
	u.SetExcluded(student.FieldName)
	return u
}

// SetLastName sets the "last_name" field.
func (u *StudentUpsert) SetLastName(v string) *StudentUpsert {
	u.Set(student.FieldLastName, v)
	return u
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *StudentUpsert) UpdateLastName() *StudentUpsert {
	u.SetExcluded(student.FieldLastName)
	return u
}

// SetCi sets the "ci" field.
func (u *StudentUpsert) SetCi(v string) *StudentUpsert {
	u.Set(student.FieldCi, v)
	return u
}

// UpdateCi sets the "ci" field to the value that was provided on create.
func (u *StudentUpsert) UpdateCi() *StudentUpsert {
	u.SetExcluded(student.FieldCi)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Student.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(student.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StudentUpsertOne) UpdateNewValues() *StudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(student.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Student.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *StudentUpsertOne) Ignore() *StudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StudentUpsertOne) DoNothing() *StudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StudentCreate.OnConflict
// documentation for more info.
func (u *StudentUpsertOne) Update(set func(*StudentUpsert)) *StudentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StudentUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StudentUpsertOne) SetName(v string) *StudentUpsertOne {
	return u.Update(func(s *StudentUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StudentUpsertOne) UpdateName() *StudentUpsertOne {
	return u.Update(func(s *StudentUpsert) {
		s.UpdateName()
	})
}

// SetLastName sets the "last_name" field.
func (u *StudentUpsertOne) SetLastName(v string) *StudentUpsertOne {
	return u.Update(func(s *StudentUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *StudentUpsertOne) UpdateLastName() *StudentUpsertOne {
	return u.Update(func(s *StudentUpsert) {
		s.UpdateLastName()
	})
}

// SetCi sets the "ci" field.
func (u *StudentUpsertOne) SetCi(v string) *StudentUpsertOne {
	return u.Update(func(s *StudentUpsert) {
		s.SetCi(v)
	})
}

// UpdateCi sets the "ci" field to the value that was provided on create.
func (u *StudentUpsertOne) UpdateCi() *StudentUpsertOne {
	return u.Update(func(s *StudentUpsert) {
		s.UpdateCi()
	})
}

// Exec executes the query.
func (u *StudentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StudentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StudentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StudentUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: StudentUpsertOne.ID is not supported by MySQL driver. Use StudentUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StudentUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StudentCreateBulk is the builder for creating many Student entities in bulk.
type StudentCreateBulk struct {
	config
	builders []*StudentCreate
	conflict []sql.ConflictOption
}

// Save creates the Student entities in the database.
func (scb *StudentCreateBulk) Save(ctx context.Context) ([]*Student, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Student, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudentMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StudentCreateBulk) SaveX(ctx context.Context) []*Student {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StudentCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StudentCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Student.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StudentUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (scb *StudentCreateBulk) OnConflict(opts ...sql.ConflictOption) *StudentUpsertBulk {
	scb.conflict = opts
	return &StudentUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Student.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *StudentCreateBulk) OnConflictColumns(columns ...string) *StudentUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &StudentUpsertBulk{
		create: scb,
	}
}

// StudentUpsertBulk is the builder for "upsert"-ing
// a bulk of Student nodes.
type StudentUpsertBulk struct {
	create *StudentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Student.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(student.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StudentUpsertBulk) UpdateNewValues() *StudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(student.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Student.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *StudentUpsertBulk) Ignore() *StudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StudentUpsertBulk) DoNothing() *StudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StudentCreateBulk.OnConflict
// documentation for more info.
func (u *StudentUpsertBulk) Update(set func(*StudentUpsert)) *StudentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StudentUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StudentUpsertBulk) SetName(v string) *StudentUpsertBulk {
	return u.Update(func(s *StudentUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StudentUpsertBulk) UpdateName() *StudentUpsertBulk {
	return u.Update(func(s *StudentUpsert) {
		s.UpdateName()
	})
}

// SetLastName sets the "last_name" field.
func (u *StudentUpsertBulk) SetLastName(v string) *StudentUpsertBulk {
	return u.Update(func(s *StudentUpsert) {
		s.SetLastName(v)
	})
}

// UpdateLastName sets the "last_name" field to the value that was provided on create.
func (u *StudentUpsertBulk) UpdateLastName() *StudentUpsertBulk {
	return u.Update(func(s *StudentUpsert) {
		s.UpdateLastName()
	})
}

// SetCi sets the "ci" field.
func (u *StudentUpsertBulk) SetCi(v string) *StudentUpsertBulk {
	return u.Update(func(s *StudentUpsert) {
		s.SetCi(v)
	})
}

// UpdateCi sets the "ci" field to the value that was provided on create.
func (u *StudentUpsertBulk) UpdateCi() *StudentUpsertBulk {
	return u.Update(func(s *StudentUpsert) {
		s.UpdateCi()
	})
}

// Exec executes the query.
func (u *StudentUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StudentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StudentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StudentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
