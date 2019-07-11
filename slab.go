package slab

// Entry entry
type Entry struct {
	Vacant int
	Val    interface{}
}

// Slab pre-allocated storage for a uniform data type
type Slab struct {
	entries []*Entry
	next    int
	len     int
}

// NewSlab creates new Slab list
func NewSlab(size int) *Slab {
	slab := &Slab{
		entries: make([]*Entry, 0, size),
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
		s.entries = append(s.entries, &Entry{Val: val})
		s.next = key + 1
	} else {
		entry := s.get(key)
		if entry != nil && entry.Val != nil {
			panic("unreachable " + string(key))
		}
		s.next = entry.Vacant
		s.entries[key] = &Entry{Val: val}
	}
	s.len++
}

func (s *Slab) get(key int) *Entry {
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
	entry := s.get(key)
	if entry != nil {
		return entry.Val
	}
	return nil
}

// Remove a val
func (s *Slab) Remove(key int) {
	entry := s.get(key)
	if entry == nil || entry.Val == nil {
		return
	}
	s.entries[key] = &Entry{
		Vacant: s.next,
	}
	s.len--
	s.next = key
}
