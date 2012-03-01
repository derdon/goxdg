package basedir

import "testing"

func TestMapEquality(t *testing.T) {
	m1 := make(environ)
	m1["foo"] = "bar"
	m2 := make(environ)
	m2["foo"] = "bar"
	if !m1.Eq(m2) {
		t.Errorf("expected %v and %v to be equal", m1, m2)
	}
}

func TestDifferentLengthMapEquality(t *testing.T) {
	m1 := make(environ)
	m1["foo"] = "bar"
	m1["spam"] = "eggs"
	m2 := make(environ)
	m2["foo"] = "bar"
	if m1.Eq(m2) {
		t.Errorf("%v ≠ %v", m1, m2)
	}
}

func TestFirstMapEmptyMapEquality(t *testing.T) {
	m1 := make(environ)
	m2 := make(environ)
	m2["foo"] = "bar"
	if m1.Eq(m2) {
		t.Errorf("%v ≠ %v", m1, m2)
	}
}
