package slab

import (
	"sync"
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

type st struct {
	a string
}

func TestPointer(t *testing.T) {
	s := NewSlab(2)
	vi := &st{a: "ok"}
	a := s.Insert(vi)
	v := s.Get(a).(*st)
	if v.a != "ok" {
		t.Fatal("invaild pointer val")
	}
}

func TestRace(t *testing.T) {
	s := NewSlab(2)
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)
	go func() {
		mu.Lock()
		s.Insert("a")
		mu.Unlock()
		wg.Done()
	}()
	go func() {
		mu.Lock()
		s.Insert("b")
		mu.Unlock()
		wg.Done()
	}()
	wg.Wait()
}
