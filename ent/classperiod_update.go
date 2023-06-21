// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/activity"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/class"
	"github.com/vmkevv/rigelapi/ent/classperiod"
	"github.com/vmkevv/rigelapi/ent/period"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// ClassPeriodUpdate is the builder for updating ClassPeriod entities.
type ClassPeriodUpdate struct {
	config
	hooks    []Hook
	mutation *ClassPeriodMutation
}

// Where appends a list predicates to the ClassPeriodUpdate builder.
func (cpu *ClassPeriodUpdate) Where(ps ...predicate.ClassPeriod) *ClassPeriodUpdate {
	cpu.mutation.Where(ps...)
	return cpu
}

// SetStart sets the "start" field.
func (cpu *ClassPeriodUpdate) SetStart(t time.Time) *ClassPeriodUpdate {
	cpu.mutation.SetStart(t)
	return cpu
}

// SetEnd sets the "end" field.
func (cpu *ClassPeriodUpdate) SetEnd(t time.Time) *ClassPeriodUpdate {
	cpu.mutation.SetEnd(t)
	return cpu
}

// SetFinished sets the "finished" field.
func (cpu *ClassPeriodUpdate) SetFinished(b bool) *ClassPeriodUpdate {
	cpu.mutation.SetFinished(b)
	return cpu
}

// AddAttendanceDayIDs adds the "attendanceDays" edge to the AttendanceDay entity by IDs.
func (cpu *ClassPeriodUpdate) AddAttendanceDayIDs(ids ...string) *ClassPeriodUpdate {
	cpu.mutation.AddAttendanceDayIDs(ids...)
	return cpu
}

// AddAttendanceDays adds the "attendanceDays" edges to the AttendanceDay entity.
func (cpu *ClassPeriodUpdate) AddAttendanceDays(a ...*AttendanceDay) *ClassPeriodUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpu.AddAttendanceDayIDs(ids...)
}

// AddActivityIDs adds the "activities" edge to the Activity entity by IDs.
func (cpu *ClassPeriodUpdate) AddActivityIDs(ids ...string) *ClassPeriodUpdate {
	cpu.mutation.AddActivityIDs(ids...)
	return cpu
}

// AddActivities adds the "activities" edges to the Activity entity.
func (cpu *ClassPeriodUpdate) AddActivities(a ...*Activity) *ClassPeriodUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpu.AddActivityIDs(ids...)
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (cpu *ClassPeriodUpdate) SetClassID(id string) *ClassPeriodUpdate {
	cpu.mutation.SetClassID(id)
	return cpu
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (cpu *ClassPeriodUpdate) SetNillableClassID(id *string) *ClassPeriodUpdate {
	if id != nil {
		cpu = cpu.SetClassID(*id)
	}
	return cpu
}

// SetClass sets the "class" edge to the Class entity.
func (cpu *ClassPeriodUpdate) SetClass(c *Class) *ClassPeriodUpdate {
	return cpu.SetClassID(c.ID)
}

// SetPeriodID sets the "period" edge to the Period entity by ID.
func (cpu *ClassPeriodUpdate) SetPeriodID(id string) *ClassPeriodUpdate {
	cpu.mutation.SetPeriodID(id)
	return cpu
}

// SetNillablePeriodID sets the "period" edge to the Period entity by ID if the given value is not nil.
func (cpu *ClassPeriodUpdate) SetNillablePeriodID(id *string) *ClassPeriodUpdate {
	if id != nil {
		cpu = cpu.SetPeriodID(*id)
	}
	return cpu
}

// SetPeriod sets the "period" edge to the Period entity.
func (cpu *ClassPeriodUpdate) SetPeriod(p *Period) *ClassPeriodUpdate {
	return cpu.SetPeriodID(p.ID)
}

// Mutation returns the ClassPeriodMutation object of the builder.
func (cpu *ClassPeriodUpdate) Mutation() *ClassPeriodMutation {
	return cpu.mutation
}

// ClearAttendanceDays clears all "attendanceDays" edges to the AttendanceDay entity.
func (cpu *ClassPeriodUpdate) ClearAttendanceDays() *ClassPeriodUpdate {
	cpu.mutation.ClearAttendanceDays()
	return cpu
}

// RemoveAttendanceDayIDs removes the "attendanceDays" edge to AttendanceDay entities by IDs.
func (cpu *ClassPeriodUpdate) RemoveAttendanceDayIDs(ids ...string) *ClassPeriodUpdate {
	cpu.mutation.RemoveAttendanceDayIDs(ids...)
	return cpu
}

// RemoveAttendanceDays removes "attendanceDays" edges to AttendanceDay entities.
func (cpu *ClassPeriodUpdate) RemoveAttendanceDays(a ...*AttendanceDay) *ClassPeriodUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpu.RemoveAttendanceDayIDs(ids...)
}

// ClearActivities clears all "activities" edges to the Activity entity.
func (cpu *ClassPeriodUpdate) ClearActivities() *ClassPeriodUpdate {
	cpu.mutation.ClearActivities()
	return cpu
}

// RemoveActivityIDs removes the "activities" edge to Activity entities by IDs.
func (cpu *ClassPeriodUpdate) RemoveActivityIDs(ids ...string) *ClassPeriodUpdate {
	cpu.mutation.RemoveActivityIDs(ids...)
	return cpu
}

// RemoveActivities removes "activities" edges to Activity entities.
func (cpu *ClassPeriodUpdate) RemoveActivities(a ...*Activity) *ClassPeriodUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpu.RemoveActivityIDs(ids...)
}

// ClearClass clears the "class" edge to the Class entity.
func (cpu *ClassPeriodUpdate) ClearClass() *ClassPeriodUpdate {
	cpu.mutation.ClearClass()
	return cpu
}

// ClearPeriod clears the "period" edge to the Period entity.
func (cpu *ClassPeriodUpdate) ClearPeriod() *ClassPeriodUpdate {
	cpu.mutation.ClearPeriod()
	return cpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cpu *ClassPeriodUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cpu.sqlSave, cpu.mutation, cpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cpu *ClassPeriodUpdate) SaveX(ctx context.Context) int {
	affected, err := cpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cpu *ClassPeriodUpdate) Exec(ctx context.Context) error {
	_, err := cpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpu *ClassPeriodUpdate) ExecX(ctx context.Context) {
	if err := cpu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cpu *ClassPeriodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(classperiod.Table, classperiod.Columns, sqlgraph.NewFieldSpec(classperiod.FieldID, field.TypeString))
	if ps := cpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpu.mutation.Start(); ok {
		_spec.SetField(classperiod.FieldStart, field.TypeTime, value)
	}
	if value, ok := cpu.mutation.End(); ok {
		_spec.SetField(classperiod.FieldEnd, field.TypeTime, value)
	}
	if value, ok := cpu.mutation.Finished(); ok {
		_spec.SetField(classperiod.FieldFinished, field.TypeBool, value)
	}
	if cpu.mutation.AttendanceDaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.AttendanceDaysTable,
			Columns: []string{classperiod.AttendanceDaysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendanceday.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.RemovedAttendanceDaysIDs(); len(nodes) > 0 && !cpu.mutation.AttendanceDaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.AttendanceDaysTable,
			Columns: []string{classperiod.AttendanceDaysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendanceday.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.AttendanceDaysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.AttendanceDaysTable,
			Columns: []string{classperiod.AttendanceDaysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendanceday.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cpu.mutation.ActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.ActivitiesTable,
			Columns: []string{classperiod.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.RemovedActivitiesIDs(); len(nodes) > 0 && !cpu.mutation.ActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.ActivitiesTable,
			Columns: []string{classperiod.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.ActivitiesTable,
			Columns: []string{classperiod.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cpu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classperiod.ClassTable,
			Columns: []string{classperiod.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classperiod.ClassTable,
			Columns: []string{classperiod.ClassColumn},
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
	if cpu.mutation.PeriodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classperiod.PeriodTable,
			Columns: []string{classperiod.PeriodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(period.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.PeriodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classperiod.PeriodTable,
			Columns: []string{classperiod.PeriodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(period.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{classperiod.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cpu.mutation.done = true
	return n, nil
}

// ClassPeriodUpdateOne is the builder for updating a single ClassPeriod entity.
type ClassPeriodUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClassPeriodMutation
}

// SetStart sets the "start" field.
func (cpuo *ClassPeriodUpdateOne) SetStart(t time.Time) *ClassPeriodUpdateOne {
	cpuo.mutation.SetStart(t)
	return cpuo
}

// SetEnd sets the "end" field.
func (cpuo *ClassPeriodUpdateOne) SetEnd(t time.Time) *ClassPeriodUpdateOne {
	cpuo.mutation.SetEnd(t)
	return cpuo
}

// SetFinished sets the "finished" field.
func (cpuo *ClassPeriodUpdateOne) SetFinished(b bool) *ClassPeriodUpdateOne {
	cpuo.mutation.SetFinished(b)
	return cpuo
}

// AddAttendanceDayIDs adds the "attendanceDays" edge to the AttendanceDay entity by IDs.
func (cpuo *ClassPeriodUpdateOne) AddAttendanceDayIDs(ids ...string) *ClassPeriodUpdateOne {
	cpuo.mutation.AddAttendanceDayIDs(ids...)
	return cpuo
}

// AddAttendanceDays adds the "attendanceDays" edges to the AttendanceDay entity.
func (cpuo *ClassPeriodUpdateOne) AddAttendanceDays(a ...*AttendanceDay) *ClassPeriodUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpuo.AddAttendanceDayIDs(ids...)
}

// AddActivityIDs adds the "activities" edge to the Activity entity by IDs.
func (cpuo *ClassPeriodUpdateOne) AddActivityIDs(ids ...string) *ClassPeriodUpdateOne {
	cpuo.mutation.AddActivityIDs(ids...)
	return cpuo
}

// AddActivities adds the "activities" edges to the Activity entity.
func (cpuo *ClassPeriodUpdateOne) AddActivities(a ...*Activity) *ClassPeriodUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpuo.AddActivityIDs(ids...)
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (cpuo *ClassPeriodUpdateOne) SetClassID(id string) *ClassPeriodUpdateOne {
	cpuo.mutation.SetClassID(id)
	return cpuo
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (cpuo *ClassPeriodUpdateOne) SetNillableClassID(id *string) *ClassPeriodUpdateOne {
	if id != nil {
		cpuo = cpuo.SetClassID(*id)
	}
	return cpuo
}

// SetClass sets the "class" edge to the Class entity.
func (cpuo *ClassPeriodUpdateOne) SetClass(c *Class) *ClassPeriodUpdateOne {
	return cpuo.SetClassID(c.ID)
}

// SetPeriodID sets the "period" edge to the Period entity by ID.
func (cpuo *ClassPeriodUpdateOne) SetPeriodID(id string) *ClassPeriodUpdateOne {
	cpuo.mutation.SetPeriodID(id)
	return cpuo
}

// SetNillablePeriodID sets the "period" edge to the Period entity by ID if the given value is not nil.
func (cpuo *ClassPeriodUpdateOne) SetNillablePeriodID(id *string) *ClassPeriodUpdateOne {
	if id != nil {
		cpuo = cpuo.SetPeriodID(*id)
	}
	return cpuo
}

// SetPeriod sets the "period" edge to the Period entity.
func (cpuo *ClassPeriodUpdateOne) SetPeriod(p *Period) *ClassPeriodUpdateOne {
	return cpuo.SetPeriodID(p.ID)
}

// Mutation returns the ClassPeriodMutation object of the builder.
func (cpuo *ClassPeriodUpdateOne) Mutation() *ClassPeriodMutation {
	return cpuo.mutation
}

// ClearAttendanceDays clears all "attendanceDays" edges to the AttendanceDay entity.
func (cpuo *ClassPeriodUpdateOne) ClearAttendanceDays() *ClassPeriodUpdateOne {
	cpuo.mutation.ClearAttendanceDays()
	return cpuo
}

// RemoveAttendanceDayIDs removes the "attendanceDays" edge to AttendanceDay entities by IDs.
func (cpuo *ClassPeriodUpdateOne) RemoveAttendanceDayIDs(ids ...string) *ClassPeriodUpdateOne {
	cpuo.mutation.RemoveAttendanceDayIDs(ids...)
	return cpuo
}

// RemoveAttendanceDays removes "attendanceDays" edges to AttendanceDay entities.
func (cpuo *ClassPeriodUpdateOne) RemoveAttendanceDays(a ...*AttendanceDay) *ClassPeriodUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpuo.RemoveAttendanceDayIDs(ids...)
}

// ClearActivities clears all "activities" edges to the Activity entity.
func (cpuo *ClassPeriodUpdateOne) ClearActivities() *ClassPeriodUpdateOne {
	cpuo.mutation.ClearActivities()
	return cpuo
}

// RemoveActivityIDs removes the "activities" edge to Activity entities by IDs.
func (cpuo *ClassPeriodUpdateOne) RemoveActivityIDs(ids ...string) *ClassPeriodUpdateOne {
	cpuo.mutation.RemoveActivityIDs(ids...)
	return cpuo
}

// RemoveActivities removes "activities" edges to Activity entities.
func (cpuo *ClassPeriodUpdateOne) RemoveActivities(a ...*Activity) *ClassPeriodUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpuo.RemoveActivityIDs(ids...)
}

// ClearClass clears the "class" edge to the Class entity.
func (cpuo *ClassPeriodUpdateOne) ClearClass() *ClassPeriodUpdateOne {
	cpuo.mutation.ClearClass()
	return cpuo
}

// ClearPeriod clears the "period" edge to the Period entity.
func (cpuo *ClassPeriodUpdateOne) ClearPeriod() *ClassPeriodUpdateOne {
	cpuo.mutation.ClearPeriod()
	return cpuo
}

// Where appends a list predicates to the ClassPeriodUpdate builder.
func (cpuo *ClassPeriodUpdateOne) Where(ps ...predicate.ClassPeriod) *ClassPeriodUpdateOne {
	cpuo.mutation.Where(ps...)
	return cpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cpuo *ClassPeriodUpdateOne) Select(field string, fields ...string) *ClassPeriodUpdateOne {
	cpuo.fields = append([]string{field}, fields...)
	return cpuo
}

// Save executes the query and returns the updated ClassPeriod entity.
func (cpuo *ClassPeriodUpdateOne) Save(ctx context.Context) (*ClassPeriod, error) {
	return withHooks(ctx, cpuo.sqlSave, cpuo.mutation, cpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cpuo *ClassPeriodUpdateOne) SaveX(ctx context.Context) *ClassPeriod {
	node, err := cpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cpuo *ClassPeriodUpdateOne) Exec(ctx context.Context) error {
	_, err := cpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpuo *ClassPeriodUpdateOne) ExecX(ctx context.Context) {
	if err := cpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cpuo *ClassPeriodUpdateOne) sqlSave(ctx context.Context) (_node *ClassPeriod, err error) {
	_spec := sqlgraph.NewUpdateSpec(classperiod.Table, classperiod.Columns, sqlgraph.NewFieldSpec(classperiod.FieldID, field.TypeString))
	id, ok := cpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ClassPeriod.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, classperiod.FieldID)
		for _, f := range fields {
			if !classperiod.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != classperiod.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpuo.mutation.Start(); ok {
		_spec.SetField(classperiod.FieldStart, field.TypeTime, value)
	}
	if value, ok := cpuo.mutation.End(); ok {
		_spec.SetField(classperiod.FieldEnd, field.TypeTime, value)
	}
	if value, ok := cpuo.mutation.Finished(); ok {
		_spec.SetField(classperiod.FieldFinished, field.TypeBool, value)
	}
	if cpuo.mutation.AttendanceDaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.AttendanceDaysTable,
			Columns: []string{classperiod.AttendanceDaysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendanceday.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.RemovedAttendanceDaysIDs(); len(nodes) > 0 && !cpuo.mutation.AttendanceDaysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.AttendanceDaysTable,
			Columns: []string{classperiod.AttendanceDaysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendanceday.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.AttendanceDaysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.AttendanceDaysTable,
			Columns: []string{classperiod.AttendanceDaysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendanceday.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cpuo.mutation.ActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.ActivitiesTable,
			Columns: []string{classperiod.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.RemovedActivitiesIDs(); len(nodes) > 0 && !cpuo.mutation.ActivitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.ActivitiesTable,
			Columns: []string{classperiod.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.ActivitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classperiod.ActivitiesTable,
			Columns: []string{classperiod.ActivitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activity.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cpuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classperiod.ClassTable,
			Columns: []string{classperiod.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classperiod.ClassTable,
			Columns: []string{classperiod.ClassColumn},
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
	if cpuo.mutation.PeriodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classperiod.PeriodTable,
			Columns: []string{classperiod.PeriodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(period.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.PeriodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classperiod.PeriodTable,
			Columns: []string{classperiod.PeriodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(period.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ClassPeriod{config: cpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{classperiod.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cpuo.mutation.done = true
	return _node, nil
}
