package types2

import "cmd/compile/internal/syntax"

// A Tuple represents an ordered list values.
type TupleType struct {
	elems []Type
}

// NewTupleType returns a new tuple type with the given element types.
func NewTupleType(elems []Type) *TupleType {
	return &TupleType{elems}
}

// Len returns the number of tuple elements.
func (t *TupleType) Len() int {
	return len(t.elems)
}

// At returns the i'th element type of tuple t.
func (t *TupleType) At(i int) Type { return t.elems[i] }

func (t *TupleType) Underlying() Type { return t }
func (t *TupleType) String() string   { return TypeString(t, nil) }

// ----------------------------------------------------------------------------
// Implementation

func (check *Checker) tupleType(typ *TupleType, e *syntax.TupleExpr) {
	typ.elems = make([]Type, len(e.ElemList))
	for i, e := range e.ElemList {
		typ.elems[i] = check.varType(e)
	}
}
