// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/apperror"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// AppErrorDelete is the builder for deleting a AppError entity.
type AppErrorDelete struct {
	config
	hooks    []Hook
	mutation *AppErrorMutation
}

// Where appends a list predicates to the AppErrorDelete builder.
func (aed *AppErrorDelete) Where(ps ...predicate.AppError) *AppErrorDelete {
	aed.mutation.Where(ps...)
	return aed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aed *AppErrorDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, aed.sqlExec, aed.mutation, aed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (aed *AppErrorDelete) ExecX(ctx context.Context) int {
	n, err := aed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aed *AppErrorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(apperror.Table, sqlgraph.NewFieldSpec(apperror.FieldID, field.TypeString))
	if ps := aed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, aed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	aed.mutation.done = true
	return affected, err
}

// AppErrorDeleteOne is the builder for deleting a single AppError entity.
type AppErrorDeleteOne struct {
	aed *AppErrorDelete
}

// Where appends a list predicates to the AppErrorDelete builder.
func (aedo *AppErrorDeleteOne) Where(ps ...predicate.AppError) *AppErrorDeleteOne {
	aedo.aed.mutation.Where(ps...)
	return aedo
}

// Exec executes the deletion query.
func (aedo *AppErrorDeleteOne) Exec(ctx context.Context) error {
	n, err := aedo.aed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{apperror.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aedo *AppErrorDeleteOne) ExecX(ctx context.Context) {
	if err := aedo.Exec(ctx); err != nil {
		panic(err)
	}
}
