package set

type Set[T comparable] struct {
	elementsMap map[T]struct{}
}

func New[T comparable]() Set[T] {
	elementsMap := make(map[T]struct{})
	return Set[T]{elementsMap}
}

func FromList[T comparable](elements []T) Set[T] {
	s := New[T]()
	for _, element := range elements {
		s.Add(element)
	}
	return s
}

func (s Set[T]) Size() int {
	return len(s.elementsMap)
}

func (s Set[T]) Has(element T) bool {
	_, ok := s.elementsMap[element]
	return ok
}

func (s *Set[T]) Add(element T) bool {
	if s.Has(element) {
		return false
	}

	s.elementsMap[element] = struct{}{}
	return true
}

func (s *Set[T]) Remove(element T) bool {
	if !s.Has(element) {
		return false
	}

	delete(s.elementsMap, element)
	return true
}

func (s Set[T]) List() []T {
	ret := make([]T, 0, len(s.elementsMap))
	for e := range s.elementsMap {
		ret = append(ret, e)
	}
	return ret
}
