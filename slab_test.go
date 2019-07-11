package slab

import (
	"testing"
)

func TestInsertGet(t *testing.T) {
	s := NewSlab(3)
	a := s.Insert("a")
	if s.Len() != 1 {
		t.Fatal("invalid len")
	}
	if s.Get(a) != "a" {
		t.Fatal("invalid val")
	}
}

func TestRemove(t *testing.T) {
	s := NewSlab(2)
	a := s.Insert("a")
	s.Insert("b")
	s.Remove(a)
	s.Insert("c")
	if len(s.entries) != 2 && s.len == 2 {
		t.Fatal("invalid real len")
	}
	// 允许相同位置新的成员
	if s.Get(a) != "c" {
		t.Fatal("invaild new val")
	}
}
