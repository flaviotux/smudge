package todo

import (
	"github.com/scylladb/gocqlx/v2/qb"
)

// ID filters vertices based on their ID field.
func ID(id string) qb.Cmp {
	return qb.EqLit(FieldID, id)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) qb.Cmp {
	return qb.EqLit(FieldID, id)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) qb.Cmp {
	return qb.NeLit(FieldID, id)
}

// IDIn applies the In predicate on the ID field.
func IDIn(id string) qb.Cmp {
	return qb.InLit(FieldID, id)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) qb.Cmp {
	return qb.GtLit(FieldID, id)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) qb.Cmp {
	return qb.GtOrEqLit(FieldID, id)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) qb.Cmp {
	return qb.LtLit(FieldID, id)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) qb.Cmp {
	return qb.LtOrEqLit(FieldID, id)
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) qb.Cmp {
	return qb.EqLit(FieldText, v)
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) qb.Cmp {
	return qb.EqLit(FieldText, v)
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) qb.Cmp {
	return qb.NeLit(FieldText, v)
}

// TextIn applies the In predicate on the "text" field.
func TextIn(v string) qb.Cmp {
	return qb.InLit(FieldText, v)
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) qb.Cmp {
	return qb.GtLit(FieldText, v)
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) qb.Cmp {
	return qb.GtOrEqLit(FieldText, v)
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) qb.Cmp {
	return qb.LtLit(FieldText, v)
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) qb.Cmp {
	return qb.LtOrEqLit(FieldText, v)
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) qb.Cmp {
	return qb.ContainsLit(FieldText, v)
}

// Done applies equality check predicate on the "done" field. It's identical to DoneEQ.
func Done(v string) qb.Cmp {
	return qb.EqLit(FieldDone, v)
}

// DoneEQ applies the EQ predicate on the "done" field.
func DoneEQ(v string) qb.Cmp {
	return qb.EqLit(FieldDone, v)
}

// DoneNEQ applies the NEQ predicate on the "done" field.
func DoneNEQ(v string) qb.Cmp {
	return qb.NeLit(FieldDone, v)
}

// UserID filters vertices based on their UserID field.
func UserID(id string) qb.Cmp {
	return qb.EqLit(FieldUserID, id)
}

// UserIDEQ applies the EQ predicate on the UserID field.
func UserIDEQ(id string) qb.Cmp {
	return qb.EqLit(FieldUserID, id)
}

// UserIDNEQ applies the NEQ predicate on the UserID field.
func UserIDNEQ(id string) qb.Cmp {
	return qb.NeLit(FieldUserID, id)
}

// UserIDIn applies the In predicate on the UserID field.
func UserIDIn(id string) qb.Cmp {
	return qb.InLit(FieldUserID, id)
}

// UserIDGT applies the GT predicate on the UserID field.
func UserIDGT(id string) qb.Cmp {
	return qb.GtLit(FieldUserID, id)
}

// UserIDGTE applies the GTE predicate on the UserID field.
func UserIDGTE(id string) qb.Cmp {
	return qb.GtOrEqLit(FieldUserID, id)
}

// UserIDLT applies the LT predicate on the UserID field.
func UserIDLT(id string) qb.Cmp {
	return qb.LtLit(FieldUserID, id)
}

// UserIDLTE applies the LTE predicate on the UserID field.
func UserIDLTE(id string) qb.Cmp {
	return qb.LtOrEqLit(FieldUserID, id)
}
