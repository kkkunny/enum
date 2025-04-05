package enum

import "testing"

func TestNewString(t *testing.T) {
	v := New[struct {
		A string
		B string
	}]()
	if v.A != "A" {
		t.FailNow()
	}
	if v.B != "B" {
		t.FailNow()
	}
}

func TestNewStringTag(t *testing.T) {
	v := New[struct {
		A string `enum:"a"`
		B string
	}]()
	if v.A != "a" {
		t.FailNow()
	}
	if v.B != "B" {
		t.FailNow()
	}
}

func TestNewInteger(t *testing.T) {
	v := New[struct {
		Zero int64
		One  int64
	}]()
	if v.Zero != 0 {
		t.FailNow()
	}
	if v.One != 1 {
		t.FailNow()
	}
}

func TestNewIntegerTag(t *testing.T) {
	v := New[struct {
		Zero int64 `enum:"1"`
		One  int64
	}]()
	if v.Zero != 1 {
		t.FailNow()
	}
	if v.One != 2 {
		t.FailNow()
	}
}

func TestNewInteger2(t *testing.T) {
	v := New[struct {
		Zero int64
		one  int64
		Two  int64
	}]()
	if v.Zero != 0 {
		t.FailNow()
	}
	if v.one != 0 {
		t.FailNow()
	}
	if v.Two != 2 {
		t.FailNow()
	}
}
