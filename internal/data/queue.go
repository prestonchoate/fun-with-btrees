package data

type Queue[T any] struct {
	items []*T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]*T, 0),
	}
}

func (s *Queue[T]) Enqueue(val T) {
	s.items = append(s.items, &val)
}

func (s *Queue[T]) Dequeue() *T {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

func (s *Queue[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Queue[T]) Reset() {
	s.items = []*T{}
}

func (s *Queue[T]) Dump() []*T {
	vals := make([]*T, len(s.items))
	copy(vals, s.items)
	s.items = []*T{}
	return vals
}

func (s *Queue[T]) Peek() *T {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[0]
}
