// Code generated by entc, DO NOT EDIT.

package userdetail

import (
	"orse/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// Rank applies equality check predicate on the "rank" field. It's identical to RankEQ.
func Rank(v float64) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// Pic applies equality check predicate on the "pic" field. It's identical to PicEQ.
func Pic(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPic), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.UserDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.UserDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.UserDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.UserDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAge), v))
	})
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.UserDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAge), v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.UserDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAge), v...))
	})
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAge), v))
	})
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAge), v))
	})
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAge), v))
	})
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAge), v))
	})
}

// AgeIsNil applies the IsNil predicate on the "age" field.
func AgeIsNil() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAge)))
	})
}

// AgeNotNil applies the NotNil predicate on the "age" field.
func AgeNotNil() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAge)))
	})
}

// RankEQ applies the EQ predicate on the "rank" field.
func RankEQ(v float64) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// RankNEQ applies the NEQ predicate on the "rank" field.
func RankNEQ(v float64) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRank), v))
	})
}

// RankIn applies the In predicate on the "rank" field.
func RankIn(vs ...float64) predicate.UserDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRank), v...))
	})
}

// RankNotIn applies the NotIn predicate on the "rank" field.
func RankNotIn(vs ...float64) predicate.UserDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRank), v...))
	})
}

// RankGT applies the GT predicate on the "rank" field.
func RankGT(v float64) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRank), v))
	})
}

// RankGTE applies the GTE predicate on the "rank" field.
func RankGTE(v float64) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRank), v))
	})
}

// RankLT applies the LT predicate on the "rank" field.
func RankLT(v float64) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRank), v))
	})
}

// RankLTE applies the LTE predicate on the "rank" field.
func RankLTE(v float64) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRank), v))
	})
}

// RankIsNil applies the IsNil predicate on the "rank" field.
func RankIsNil() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRank)))
	})
}

// RankNotNil applies the NotNil predicate on the "rank" field.
func RankNotNil() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRank)))
	})
}

// PicEQ applies the EQ predicate on the "pic" field.
func PicEQ(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPic), v))
	})
}

// PicNEQ applies the NEQ predicate on the "pic" field.
func PicNEQ(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPic), v))
	})
}

// PicIn applies the In predicate on the "pic" field.
func PicIn(vs ...string) predicate.UserDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPic), v...))
	})
}

// PicNotIn applies the NotIn predicate on the "pic" field.
func PicNotIn(vs ...string) predicate.UserDetail {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserDetail(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPic), v...))
	})
}

// PicGT applies the GT predicate on the "pic" field.
func PicGT(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPic), v))
	})
}

// PicGTE applies the GTE predicate on the "pic" field.
func PicGTE(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPic), v))
	})
}

// PicLT applies the LT predicate on the "pic" field.
func PicLT(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPic), v))
	})
}

// PicLTE applies the LTE predicate on the "pic" field.
func PicLTE(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPic), v))
	})
}

// PicContains applies the Contains predicate on the "pic" field.
func PicContains(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPic), v))
	})
}

// PicHasPrefix applies the HasPrefix predicate on the "pic" field.
func PicHasPrefix(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPic), v))
	})
}

// PicHasSuffix applies the HasSuffix predicate on the "pic" field.
func PicHasSuffix(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPic), v))
	})
}

// PicIsNil applies the IsNil predicate on the "pic" field.
func PicIsNil() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPic)))
	})
}

// PicNotNil applies the NotNil predicate on the "pic" field.
func PicNotNil() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPic)))
	})
}

// PicEqualFold applies the EqualFold predicate on the "pic" field.
func PicEqualFold(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPic), v))
	})
}

// PicContainsFold applies the ContainsFold predicate on the "pic" field.
func PicContainsFold(v string) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPic), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserDetail) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserDetail) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
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
func Not(p predicate.UserDetail) predicate.UserDetail {
	return predicate.UserDetail(func(s *sql.Selector) {
		p(s.Not())
	})
}
