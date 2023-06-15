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
	"github.com/vmkevv/rigelapi/ent/activity"
	"github.com/vmkevv/rigelapi/ent/score"
	"github.com/vmkevv/rigelapi/ent/student"
)

// ScoreCreate is the builder for creating a Score entity.
type ScoreCreate struct {
	config
	mutation *ScoreMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPoints sets the "points" field.
func (sc *ScoreCreate) SetPoints(i int) *ScoreCreate {
	sc.mutation.SetPoints(i)
	return sc
}

// SetID sets the "id" field.
func (sc *ScoreCreate) SetID(s string) *ScoreCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetActivityID sets the "activity" edge to the Activity entity by ID.
func (sc *ScoreCreate) SetActivityID(id string) *ScoreCreate {
	sc.mutation.SetActivityID(id)
	return sc
}

// SetNillableActivityID sets the "activity" edge to the Activity entity by ID if the given value is not nil.
func (sc *ScoreCreate) SetNillableActivityID(id *string) *ScoreCreate {
	if id != nil {
		sc = sc.SetActivityID(*id)
	}
	return sc
}

// SetActivity sets the "activity" edge to the Activity entity.
func (sc *ScoreCreate) SetActivity(a *Activity) *ScoreCreate {
	return sc.SetActivityID(a.ID)
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (sc *ScoreCreate) SetStudentID(id string) *ScoreCreate {
	sc.mutation.SetStudentID(id)
	return sc
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (sc *ScoreCreate) SetNillableStudentID(id *string) *ScoreCreate {
	if id != nil {
		sc = sc.SetStudentID(*id)
	}
	return sc
}

// SetStudent sets the "student" edge to the Student entity.
func (sc *ScoreCreate) SetStudent(s *Student) *ScoreCreate {
	return sc.SetStudentID(s.ID)
}

// Mutation returns the ScoreMutation object of the builder.
func (sc *ScoreCreate) Mutation() *ScoreMutation {
	return sc.mutation
}

// Save creates the Score in the database.
func (sc *ScoreCreate) Save(ctx context.Context) (*Score, error) {
	var (
		err  error
		node *Score
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScoreMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Score)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ScoreMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScoreCreate) SaveX(ctx context.Context) *Score {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScoreCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScoreCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScoreCreate) check() error {
	if _, ok := sc.mutation.Points(); !ok {
		return &ValidationError{Name: "points", err: errors.New(`ent: missing required field "Score.points"`)}
	}
	return nil
}

func (sc *ScoreCreate) sqlSave(ctx context.Context) (*Score, error) {
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
			return nil, fmt.Errorf("unexpected Score.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (sc *ScoreCreate) createSpec() (*Score, *sqlgraph.CreateSpec) {
	var (
		_node = &Score{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: score.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: score.FieldID,
			},
		}
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Points(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: score.FieldPoints,
		})
		_node.Points = value
	}
	if nodes := sc.mutation.ActivityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.ActivityTable,
			Columns: []string{score.ActivityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: activity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.activity_scores = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   score.StudentTable,
			Columns: []string{score.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: student.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.student_scores = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Score.Create().
//		SetPoints(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ScoreUpsert) {
//			SetPoints(v+v).
//		}).
//		Exec(ctx)
func (sc *ScoreCreate) OnConflict(opts ...sql.ConflictOption) *ScoreUpsertOne {
	sc.conflict = opts
	return &ScoreUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Score.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *ScoreCreate) OnConflictColumns(columns ...string) *ScoreUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &ScoreUpsertOne{
		create: sc,
	}
}

type (
	// ScoreUpsertOne is the builder for "upsert"-ing
	//  one Score node.
	ScoreUpsertOne struct {
		create *ScoreCreate
	}

	// ScoreUpsert is the "OnConflict" setter.
	ScoreUpsert struct {
		*sql.UpdateSet
	}
)

// SetPoints sets the "points" field.
func (u *ScoreUpsert) SetPoints(v int) *ScoreUpsert {
	u.Set(score.FieldPoints, v)
	return u
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *ScoreUpsert) UpdatePoints() *ScoreUpsert {
	u.SetExcluded(score.FieldPoints)
	return u
}

// AddPoints adds v to the "points" field.
func (u *ScoreUpsert) AddPoints(v int) *ScoreUpsert {
	u.Add(score.FieldPoints, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Score.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(score.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ScoreUpsertOne) UpdateNewValues() *ScoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(score.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Score.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ScoreUpsertOne) Ignore() *ScoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ScoreUpsertOne) DoNothing() *ScoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ScoreCreate.OnConflict
// documentation for more info.
func (u *ScoreUpsertOne) Update(set func(*ScoreUpsert)) *ScoreUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ScoreUpsert{UpdateSet: update})
	}))
	return u
}

// SetPoints sets the "points" field.
func (u *ScoreUpsertOne) SetPoints(v int) *ScoreUpsertOne {
	return u.Update(func(s *ScoreUpsert) {
		s.SetPoints(v)
	})
}

// AddPoints adds v to the "points" field.
func (u *ScoreUpsertOne) AddPoints(v int) *ScoreUpsertOne {
	return u.Update(func(s *ScoreUpsert) {
		s.AddPoints(v)
	})
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *ScoreUpsertOne) UpdatePoints() *ScoreUpsertOne {
	return u.Update(func(s *ScoreUpsert) {
		s.UpdatePoints()
	})
}

// Exec executes the query.
func (u *ScoreUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ScoreCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ScoreUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ScoreUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ScoreUpsertOne.ID is not supported by MySQL driver. Use ScoreUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ScoreUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ScoreCreateBulk is the builder for creating many Score entities in bulk.
type ScoreCreateBulk struct {
	config
	builders []*ScoreCreate
	conflict []sql.ConflictOption
}

// Save creates the Score entities in the database.
func (scb *ScoreCreateBulk) Save(ctx context.Context) ([]*Score, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Score, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScoreMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
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
func (scb *ScoreCreateBulk) SaveX(ctx context.Context) []*Score {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScoreCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScoreCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Score.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ScoreUpsert) {
//			SetPoints(v+v).
//		}).
//		Exec(ctx)
func (scb *ScoreCreateBulk) OnConflict(opts ...sql.ConflictOption) *ScoreUpsertBulk {
	scb.conflict = opts
	return &ScoreUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Score.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *ScoreCreateBulk) OnConflictColumns(columns ...string) *ScoreUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &ScoreUpsertBulk{
		create: scb,
	}
}

// ScoreUpsertBulk is the builder for "upsert"-ing
// a bulk of Score nodes.
type ScoreUpsertBulk struct {
	create *ScoreCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Score.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(score.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ScoreUpsertBulk) UpdateNewValues() *ScoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(score.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Score.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ScoreUpsertBulk) Ignore() *ScoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ScoreUpsertBulk) DoNothing() *ScoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ScoreCreateBulk.OnConflict
// documentation for more info.
func (u *ScoreUpsertBulk) Update(set func(*ScoreUpsert)) *ScoreUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ScoreUpsert{UpdateSet: update})
	}))
	return u
}

// SetPoints sets the "points" field.
func (u *ScoreUpsertBulk) SetPoints(v int) *ScoreUpsertBulk {
	return u.Update(func(s *ScoreUpsert) {
		s.SetPoints(v)
	})
}

// AddPoints adds v to the "points" field.
func (u *ScoreUpsertBulk) AddPoints(v int) *ScoreUpsertBulk {
	return u.Update(func(s *ScoreUpsert) {
		s.AddPoints(v)
	})
}

// UpdatePoints sets the "points" field to the value that was provided on create.
func (u *ScoreUpsertBulk) UpdatePoints() *ScoreUpsertBulk {
	return u.Update(func(s *ScoreUpsert) {
		s.UpdatePoints()
	})
}

// Exec executes the query.
func (u *ScoreUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ScoreCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ScoreCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ScoreUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
