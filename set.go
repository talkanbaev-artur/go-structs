package gostructs

type Set struct {
	m map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{m: make(map[interface{}]struct{})}
}

func (s *Set) Insert(i ...interface{}) {
	for _, v := range i {
		s.m[v] = struct{}{}
	}
}

func (s *Set) Erase(i interface{}) {
	delete(s.m, i)
}

func (s *Set) Has(i interface{}) bool {
	_, ok := s.m[i]
	return ok
}

func (s *Set) Size() int {
	var cnt int
	for range s.m {
		cnt++
	}
	return cnt
}

func (s *Set) ToList() []interface{} {
	l := make([]interface{}, 0)
	for v := range s.m {
		l = append(l, v)
	}
	return l
}
