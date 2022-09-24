// Code generated by ent, DO NOT EDIT.

package attendancedaysyncs

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/vmkevv/rigelapi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// LastSyncID applies equality check predicate on the "last_sync_id" field. It's identical to LastSyncIDEQ.
func LastSyncID(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDEQ applies the EQ predicate on the "last_sync_id" field.
func LastSyncIDEQ(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDNEQ applies the NEQ predicate on the "last_sync_id" field.
func LastSyncIDNEQ(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDIn applies the In predicate on the "last_sync_id" field.
func LastSyncIDIn(vs ...string) predicate.AttendanceDaySyncs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastSyncID), v...))
	})
}

// LastSyncIDNotIn applies the NotIn predicate on the "last_sync_id" field.
func LastSyncIDNotIn(vs ...string) predicate.AttendanceDaySyncs {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastSyncID), v...))
	})
}

// LastSyncIDGT applies the GT predicate on the "last_sync_id" field.
func LastSyncIDGT(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDGTE applies the GTE predicate on the "last_sync_id" field.
func LastSyncIDGTE(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDLT applies the LT predicate on the "last_sync_id" field.
func LastSyncIDLT(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDLTE applies the LTE predicate on the "last_sync_id" field.
func LastSyncIDLTE(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDContains applies the Contains predicate on the "last_sync_id" field.
func LastSyncIDContains(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDHasPrefix applies the HasPrefix predicate on the "last_sync_id" field.
func LastSyncIDHasPrefix(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDHasSuffix applies the HasSuffix predicate on the "last_sync_id" field.
func LastSyncIDHasSuffix(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDEqualFold applies the EqualFold predicate on the "last_sync_id" field.
func LastSyncIDEqualFold(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLastSyncID), v))
	})
}

// LastSyncIDContainsFold applies the ContainsFold predicate on the "last_sync_id" field.
func LastSyncIDContainsFold(v string) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLastSyncID), v))
	})
}

// HasTeacher applies the HasEdge predicate on the "teacher" edge.
func HasTeacher() predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeacherTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeacherTable, TeacherColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeacherWith applies the HasEdge predicate on the "teacher" edge with a given conditions (other predicates).
func HasTeacherWith(preds ...predicate.Teacher) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeacherInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TeacherTable, TeacherColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AttendanceDaySyncs) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AttendanceDaySyncs) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AttendanceDaySyncs) predicate.AttendanceDaySyncs {
	return predicate.AttendanceDaySyncs(func(s *sql.Selector) {
		p(s.Not())
	})
}
