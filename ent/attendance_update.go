// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vmkevv/rigelapi/ent/attendance"
	"github.com/vmkevv/rigelapi/ent/attendanceday"
	"github.com/vmkevv/rigelapi/ent/predicate"
	"github.com/vmkevv/rigelapi/ent/student"
)

// AttendanceUpdate is the builder for updating Attendance entities.
type AttendanceUpdate struct {
	config
	hooks    []Hook
	mutation *AttendanceMutation
}

// Where appends a list predicates to the AttendanceUpdate builder.
func (au *AttendanceUpdate) Where(ps ...predicate.Attendance) *AttendanceUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetValue sets the "value" field.
func (au *AttendanceUpdate) SetValue(a attendance.Value) *AttendanceUpdate {
	au.mutation.SetValue(a)
	return au
}

// SetAttendanceDayID sets the "attendanceDay" edge to the AttendanceDay entity by ID.
func (au *AttendanceUpdate) SetAttendanceDayID(id string) *AttendanceUpdate {
	au.mutation.SetAttendanceDayID(id)
	return au
}

// SetNillableAttendanceDayID sets the "attendanceDay" edge to the AttendanceDay entity by ID if the given value is not nil.
func (au *AttendanceUpdate) SetNillableAttendanceDayID(id *string) *AttendanceUpdate {
	if id != nil {
		au = au.SetAttendanceDayID(*id)
	}
	return au
}

// SetAttendanceDay sets the "attendanceDay" edge to the AttendanceDay entity.
func (au *AttendanceUpdate) SetAttendanceDay(a *AttendanceDay) *AttendanceUpdate {
	return au.SetAttendanceDayID(a.ID)
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (au *AttendanceUpdate) SetStudentID(id string) *AttendanceUpdate {
	au.mutation.SetStudentID(id)
	return au
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (au *AttendanceUpdate) SetNillableStudentID(id *string) *AttendanceUpdate {
	if id != nil {
		au = au.SetStudentID(*id)
	}
	return au
}

// SetStudent sets the "student" edge to the Student entity.
func (au *AttendanceUpdate) SetStudent(s *Student) *AttendanceUpdate {
	return au.SetStudentID(s.ID)
}

// Mutation returns the AttendanceMutation object of the builder.
func (au *AttendanceUpdate) Mutation() *AttendanceMutation {
	return au.mutation
}

// ClearAttendanceDay clears the "attendanceDay" edge to the AttendanceDay entity.
func (au *AttendanceUpdate) ClearAttendanceDay() *AttendanceUpdate {
	au.mutation.ClearAttendanceDay()
	return au
}

// ClearStudent clears the "student" edge to the Student entity.
func (au *AttendanceUpdate) ClearStudent() *AttendanceUpdate {
	au.mutation.ClearStudent()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttendanceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttendanceUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttendanceUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttendanceUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AttendanceUpdate) check() error {
	if v, ok := au.mutation.Value(); ok {
		if err := attendance.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Attendance.value": %w`, err)}
		}
	}
	return nil
}

func (au *AttendanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(attendance.Table, attendance.Columns, sqlgraph.NewFieldSpec(attendance.FieldID, field.TypeString))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Value(); ok {
		_spec.SetField(attendance.FieldValue, field.TypeEnum, value)
	}
	if au.mutation.AttendanceDayCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendance.AttendanceDayTable,
			Columns: []string{attendance.AttendanceDayColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendanceday.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AttendanceDayIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendance.AttendanceDayTable,
			Columns: []string{attendance.AttendanceDayColumn},
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
	if au.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendance.StudentTable,
			Columns: []string{attendance.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendance.StudentTable,
			Columns: []string{attendance.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attendance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AttendanceUpdateOne is the builder for updating a single Attendance entity.
type AttendanceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttendanceMutation
}

// SetValue sets the "value" field.
func (auo *AttendanceUpdateOne) SetValue(a attendance.Value) *AttendanceUpdateOne {
	auo.mutation.SetValue(a)
	return auo
}

// SetAttendanceDayID sets the "attendanceDay" edge to the AttendanceDay entity by ID.
func (auo *AttendanceUpdateOne) SetAttendanceDayID(id string) *AttendanceUpdateOne {
	auo.mutation.SetAttendanceDayID(id)
	return auo
}

// SetNillableAttendanceDayID sets the "attendanceDay" edge to the AttendanceDay entity by ID if the given value is not nil.
func (auo *AttendanceUpdateOne) SetNillableAttendanceDayID(id *string) *AttendanceUpdateOne {
	if id != nil {
		auo = auo.SetAttendanceDayID(*id)
	}
	return auo
}

// SetAttendanceDay sets the "attendanceDay" edge to the AttendanceDay entity.
func (auo *AttendanceUpdateOne) SetAttendanceDay(a *AttendanceDay) *AttendanceUpdateOne {
	return auo.SetAttendanceDayID(a.ID)
}

// SetStudentID sets the "student" edge to the Student entity by ID.
func (auo *AttendanceUpdateOne) SetStudentID(id string) *AttendanceUpdateOne {
	auo.mutation.SetStudentID(id)
	return auo
}

// SetNillableStudentID sets the "student" edge to the Student entity by ID if the given value is not nil.
func (auo *AttendanceUpdateOne) SetNillableStudentID(id *string) *AttendanceUpdateOne {
	if id != nil {
		auo = auo.SetStudentID(*id)
	}
	return auo
}

// SetStudent sets the "student" edge to the Student entity.
func (auo *AttendanceUpdateOne) SetStudent(s *Student) *AttendanceUpdateOne {
	return auo.SetStudentID(s.ID)
}

// Mutation returns the AttendanceMutation object of the builder.
func (auo *AttendanceUpdateOne) Mutation() *AttendanceMutation {
	return auo.mutation
}

// ClearAttendanceDay clears the "attendanceDay" edge to the AttendanceDay entity.
func (auo *AttendanceUpdateOne) ClearAttendanceDay() *AttendanceUpdateOne {
	auo.mutation.ClearAttendanceDay()
	return auo
}

// ClearStudent clears the "student" edge to the Student entity.
func (auo *AttendanceUpdateOne) ClearStudent() *AttendanceUpdateOne {
	auo.mutation.ClearStudent()
	return auo
}

// Where appends a list predicates to the AttendanceUpdate builder.
func (auo *AttendanceUpdateOne) Where(ps ...predicate.Attendance) *AttendanceUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttendanceUpdateOne) Select(field string, fields ...string) *AttendanceUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attendance entity.
func (auo *AttendanceUpdateOne) Save(ctx context.Context) (*Attendance, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttendanceUpdateOne) SaveX(ctx context.Context) *Attendance {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttendanceUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttendanceUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AttendanceUpdateOne) check() error {
	if v, ok := auo.mutation.Value(); ok {
		if err := attendance.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Attendance.value": %w`, err)}
		}
	}
	return nil
}

func (auo *AttendanceUpdateOne) sqlSave(ctx context.Context) (_node *Attendance, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(attendance.Table, attendance.Columns, sqlgraph.NewFieldSpec(attendance.FieldID, field.TypeString))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Attendance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attendance.FieldID)
		for _, f := range fields {
			if !attendance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attendance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Value(); ok {
		_spec.SetField(attendance.FieldValue, field.TypeEnum, value)
	}
	if auo.mutation.AttendanceDayCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendance.AttendanceDayTable,
			Columns: []string{attendance.AttendanceDayColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attendanceday.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AttendanceDayIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendance.AttendanceDayTable,
			Columns: []string{attendance.AttendanceDayColumn},
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
	if auo.mutation.StudentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendance.StudentTable,
			Columns: []string{attendance.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.StudentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendance.StudentTable,
			Columns: []string{attendance.StudentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(student.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Attendance{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attendance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
