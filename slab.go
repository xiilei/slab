package slab

// Slab pre-allocated storage for a uniform data type
type Slab struct {
	entries []*entry
	next    int
	len     int
}

type entry struct {
	vacant int
	val    interface{}
}

// NewSlab creates new Slab list
func NewSlab(size int) *Slab {
	slab := &Slab{
		entries: make([]*entry, 0, size),
	}
	return slab
}

// Reset resets the slab to be empty
func (s *Slab) Reset() {
	s.entries = s.entries[:0]
	s.next = 0
	s.len = 0
}

// Len returns the number of entries
func (s *Slab) Len() int {
	return s.len
}

// Cap returns the capacity of the slab
func (s *Slab) Cap() int {
	return cap(s.entries)
}

// Insert a new val
func (s *Slab) Insert(val interface{}) int {
	key := s.next
	s.insertAt(key, val)
	return key
}

func (s *Slab) insertAt(key int, val interface{}) {
	if key == len(s.entries) {
		s.entries = append(s.entries, &entry{val: val})
		s.next = key + 1
	} else {
		e := s.get(key)
		if e != nil && e.val != nil {
			panic("unreachable " + string(key))
		}
		s.next = e.vacant
		s.entries[key] = &entry{val: val}
	}
	s.len++
}

func (s *Slab) get(key int) *entry {
	if key >= len(s.entries) || key < 0 {
		return nil
	}
	return s.entries[key]
}

// Contains returns the key exists
func (s *Slab) Contains(key int) bool {
	return s.Get(key) != nil
}

// Get a val
func (s *Slab) Get(key int) interface{} {
	e := s.get(key)
	if e != nil {
		return e.val
	}
	return nil
}

// Remove a val
func (s *Slab) Remove(key int) {
	e := s.get(key)
	if e == nil || e.val == nil {
		return
	}
	s.entries[key] = &entry{
		vacant: s.next,
	}
	s.len--
	s.next = key
}
