// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

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
	var (
		err  error
		node *Provincia
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvinciaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Provincia)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProvinciaMutation", v)
		}
		node = nv
	}
	return node, err
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
	return _node, nil
}

func (pc *ProvinciaCreate) createSpec() (*Provincia, *sqlgraph.CreateSpec) {
	var (
		_node = &Provincia{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: provincia.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: provincia.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: provincia.FieldName,
		})
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
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: municipio.FieldID,
				},
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
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: dpto.FieldID,
				},
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

// ProvinciaCreateBulk is the builder for creating many Provincia entities in bulk.
type ProvinciaCreateBulk struct {
	config
	builders []*ProvinciaCreate
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
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
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