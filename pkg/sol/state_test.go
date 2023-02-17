package sol

import "testing"

type VisitRec map[StateKey]struct{}

func NewVisitRec() VisitRec {
	return make(VisitRec)
}

func (v VisitRec) RecordIfNotVisited(s *StateM) bool {
	var k = s.Key()
	if _, ok := v[k]; ok {
		return false
	}
	v[k] = struct{}{}
	return true
}

func TestVisitRec_RecordIfNotVisited(t *testing.T) {
	var v = NewVisitRec()
	var s = NewState()
	var ok = v.RecordIfNotVisited(s)
	if !ok {
		t.Fatal("should be ok")
	}
	ok = v.RecordIfNotVisited(s)
	if ok {
		t.Fatal("should not be ok")
	}

	var ns = NewStateFrom(1, 2, 3, 4, 5, Pile{Cards: []Card{
		CardC(Spades, 1),
	}})
	ok = v.RecordIfNotVisited(ns)
	if !ok {
		t.Fatal("should be ok")
	}
	ok = v.RecordIfNotVisited(ns)
	if ok {
		t.Fatal("should not be ok")
	}
}
