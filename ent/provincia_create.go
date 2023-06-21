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
	"github.com/vmkevv/rigelapi/ent/dpto"
	"github.com/vmkevv/rigelapi/ent/municipio"
	"github.com/vmkevv/rigelapi/ent/provincia"
)

// ProvinciaCreate is the builder for creating a Provincia entity.
type ProvinciaCreate struct {
	config
	mutation *ProvinciaMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (pc *ProvinciaCreate) SetName(s string) *ProvinciaCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetID sets the "id" field.
func (pc *ProvinciaCreate) SetID(s string) *ProvinciaCreate {
	pc.mutation.SetID(s)
	return pc
}

// AddMunicipioIDs adds the "municipios" edge to the Municipio entity by IDs.
func (pc *ProvinciaCreate) AddMunicipioIDs(ids ...string) *ProvinciaCreate {
	pc.mutation.AddMunicipioIDs(ids...)
	return pc
}

// AddMunicipios adds the "municipios" edges to the Municipio entity.
func (pc *ProvinciaCreate) AddMunicipios(m ...*Municipio) *ProvinciaCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pc.AddMunicipioIDs(ids...)
}

// SetDepartamentoID sets the "departamento" edge to the Dpto entity by ID.
func (pc *ProvinciaCreate) SetDepartamentoID(id string) *ProvinciaCreate {
	pc.mutation.SetDepartamentoID(id)
	return pc
}

// SetNillableDepartamentoID sets the "departamento" edge to the Dpto entity by ID if the given value is not nil.
func (pc *ProvinciaCreate) SetNillableDepartamentoID(id *string) *ProvinciaCreate {
	if id != nil {
		pc = pc.SetDepartamentoID(*id)
	}
	return pc
}

// SetDepartamento sets the "departamento" edge to the Dpto entity.
func (pc *ProvinciaCreate) SetDepartamento(d *Dpto) *ProvinciaCreate {
	return pc.SetDepartamentoID(d.ID)
}

// Mutation returns the ProvinciaMutation object of the builder.
func (pc *ProvinciaCreate) Mutation() *ProvinciaMutation {
	return pc.mutation
}

// Save creates the Provincia in the database.
func (pc *ProvinciaCreate) Save(ctx context.Context) (*Provincia, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProvinciaCreate) SaveX(ctx context.Context) *Provincia {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProvinciaCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProvinciaCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProvinciaCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Provincia.name"`)}
	}
	return nil
}

func (pc *ProvinciaCreate) sqlSave(ctx context.Context) (*Provincia, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Provincia.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProvinciaCreate) createSpec() (*Provincia, *sqlgraph.CreateSpec) {
	var (
		_node = &Provincia{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(provincia.Table, sqlgraph.NewFieldSpec(provincia.FieldID, field.TypeString))
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(provincia.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := pc.mutation.MunicipiosIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.DepartamentoIDs(); len(nodes) > 0 {
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
		_node.dpto_provincias = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Provincia.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProvinciaUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pc *ProvinciaCreate) OnConflict(opts ...sql.ConflictOption) *ProvinciaUpsertOne {
	pc.conflict = opts
	return &ProvinciaUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Provincia.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *ProvinciaCreate) OnConflictColumns(columns ...string) *ProvinciaUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &ProvinciaUpsertOne{
		create: pc,
	}
}

type (
	// ProvinciaUpsertOne is the builder for "upsert"-ing
	//  one Provincia node.
	ProvinciaUpsertOne struct {
		create *ProvinciaCreate
	}

	// ProvinciaUpsert is the "OnConflict" setter.
	ProvinciaUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *ProvinciaUpsert) SetName(v string) *ProvinciaUpsert {
	u.Set(provincia.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProvinciaUpsert) UpdateName() *ProvinciaUpsert {
	u.SetExcluded(provincia.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Provincia.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(provincia.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProvinciaUpsertOne) UpdateNewValues() *ProvinciaUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(provincia.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Provincia.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProvinciaUpsertOne) Ignore() *ProvinciaUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProvinciaUpsertOne) DoNothing() *ProvinciaUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProvinciaCreate.OnConflict
// documentation for more info.
func (u *ProvinciaUpsertOne) Update(set func(*ProvinciaUpsert)) *ProvinciaUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProvinciaUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ProvinciaUpsertOne) SetName(v string) *ProvinciaUpsertOne {
	return u.Update(func(s *ProvinciaUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProvinciaUpsertOne) UpdateName() *ProvinciaUpsertOne {
	return u.Update(func(s *ProvinciaUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *ProvinciaUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProvinciaCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProvinciaUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProvinciaUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ProvinciaUpsertOne.ID is not supported by MySQL driver. Use ProvinciaUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProvinciaUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProvinciaCreateBulk is the builder for creating many Provincia entities in bulk.
type ProvinciaCreateBulk struct {
	config
	builders []*ProvinciaCreate
	conflict []sql.ConflictOption
}

// Save creates the Provincia entities in the database.
func (pcb *ProvinciaCreateBulk) Save(ctx context.Context) ([]*Provincia, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Provincia, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProvinciaMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProvinciaCreateBulk) SaveX(ctx context.Context) []*Provincia {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProvinciaCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProvinciaCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Provincia.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProvinciaUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pcb *ProvinciaCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProvinciaUpsertBulk {
	pcb.conflict = opts
	return &ProvinciaUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Provincia.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *ProvinciaCreateBulk) OnConflictColumns(columns ...string) *ProvinciaUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &ProvinciaUpsertBulk{
		create: pcb,
	}
}

// ProvinciaUpsertBulk is the builder for "upsert"-ing
// a bulk of Provincia nodes.
type ProvinciaUpsertBulk struct {
	create *ProvinciaCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Provincia.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(provincia.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProvinciaUpsertBulk) UpdateNewValues() *ProvinciaUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(provincia.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Provincia.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProvinciaUpsertBulk) Ignore() *ProvinciaUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProvinciaUpsertBulk) DoNothing() *ProvinciaUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProvinciaCreateBulk.OnConflict
// documentation for more info.
func (u *ProvinciaUpsertBulk) Update(set func(*ProvinciaUpsert)) *ProvinciaUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProvinciaUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ProvinciaUpsertBulk) SetName(v string) *ProvinciaUpsertBulk {
	return u.Update(func(s *ProvinciaUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ProvinciaUpsertBulk) UpdateName() *ProvinciaUpsertBulk {
	return u.Update(func(s *ProvinciaUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *ProvinciaUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProvinciaCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProvinciaCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProvinciaUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
