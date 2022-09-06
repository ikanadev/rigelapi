// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/activity"
	"github.com/vmkevv/rigelapi/ent/area"
	"github.com/vmkevv/rigelapi/ent/year"
)

// AreaCreate is the builder for creating a Area entity.
type AreaCreate struct {
	config
	mutation *AreaMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ac *AreaCreate) SetName(s string) *AreaCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetPoints sets the "points" field.
func (ac *AreaCreate) SetPoints(i int) *AreaCreate {
	ac.mutation.SetPoints(i)
	return ac
}

// SetID sets the "id" field.
func (ac *AreaCreate) SetID(s string) *AreaCreate {
	ac.mutation.SetID(s)
	return ac
}

// AddActivityIDs adds the "activities" edge to the Activity entity by IDs.
func (ac *AreaCreate) AddActivityIDs(ids ...string) *AreaCreate {
	ac.mutation.AddActivityIDs(ids...)
	return ac
}

// AddActivities adds the "activities" edges to the Activity entity.
func (ac *AreaCreate) AddActivities(a ...*Activity) *AreaCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddActivityIDs(ids...)
}

// SetYearID sets the "year" edge to the Year entity by ID.
func (ac *AreaCreate) SetYearID(id string) *AreaCreate {
	ac.mutation.SetYearID(id)
	return ac
}

// SetNillableYearID sets the "year" edge to the Year entity by ID if the given value is not nil.
func (ac *AreaCreate) SetNillableYearID(id *string) *AreaCreate {
	if id != nil {
		ac = ac.SetYearID(*id)
	}
	return ac
}

// SetYear sets the "year" edge to the Year entity.
func (ac *AreaCreate) SetYear(y *Year) *AreaCreate {
	return ac.SetYearID(y.ID)
}

// Mutation returns the AreaMutation object of the builder.
func (ac *AreaCreate) Mutation() *AreaMutation {
	return ac.mutation
}

// Save creates the Area in the database.
func (ac *AreaCreate) Save(ctx context.Context) (*Area, error) {
	var (
		err  error
		node *Area
	)
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AreaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Area)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AreaMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AreaCreate) SaveX(ctx context.Context) *Area {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AreaCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AreaCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AreaCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Area.name"`)}
	}
	if _, ok := ac.mutation.Points(); !ok {
		return &ValidationError{Name: "points", err: errors.New(`ent: missing required field "Area.points"`)}
	}
	return nil
}

func (ac *AreaCreate) sqlSave(ctx context.Context) (*Area, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Area.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (ac *AreaCreate) createSpec() (*Area, *sqlgraph.CreateSpec) {
	var (
		_node = &Area{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: area.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: area.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: area.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Points(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: area.FieldPoints,
		})
		_node.Points = value
	}
	if nodes := ac.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   area.ActivitiesTable,
			Columns: []string{area.ActivitiesColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.YearIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   area.YearTable,
			Columns: []string{area.YearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: year.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.year_areas = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AreaCreateBulk is the builder for creating many Area entities in bulk.
type AreaCreateBulk struct {
	config
	builders []*AreaCreate
}

// Save creates the Area entities in the database.
func (acb *AreaCreateBulk) Save(ctx context.Context) ([]*Area, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Area, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AreaMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AreaCreateBulk) SaveX(ctx context.Context) []*Area {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AreaCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AreaCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}