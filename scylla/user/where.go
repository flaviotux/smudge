package user

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

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) qb.Cmp {
	return qb.EqLit(FieldName, v)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) qb.Cmp {
	return qb.NeLit(FieldName, v)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(v string) qb.Cmp {
	return qb.InLit(FieldName, v)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) qb.Cmp {
	return qb.GtLit(FieldName, v)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) qb.Cmp {
	return qb.GtOrEqLit(FieldName, v)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) qb.Cmp {
	return qb.LtLit(FieldName, v)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) qb.Cmp {
	return qb.LtOrEqLit(FieldName, v)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) qb.Cmp {
	return qb.ContainsLit(FieldName, v)
}
