package generic_set

type gSet[V comparable] struct {
	data map[V]struct{}
}

func New[V comparable](init ...V) gSet[V] {
	ret := gSet[V]{data: make(map[V]struct{}, len(init))}
	for _, v := range init {
		ret.Add(v)
	}
	return ret
}

func (gs gSet[V]) Add(value V) {
	gs.data[value] = struct{}{}
}

func (gs gSet[V]) Remove(value V) {
	delete(gs.data, value)
}

func (gs gSet[V]) Contains(value V) bool {
	_, ok := gs.data[value]
	return ok
}

func (gs gSet[V]) Slice() []V {
	ret := make([]V, 0, len(gs.data))
	for s := range gs.data {
		ret = append(ret, s)
	}
	return ret
}

func Intersect[V comparable](s1, s2 gSet[V]) gSet[V] {
	less := s1
	more := s2
	if len(s2.data) < len(s1.data) {
		less = s2
		more = s1
	}

	ret := New[V]()
	for v := range less.data {
		if more.Contains(v) {
			ret.Add(v)
		}
	}

	return ret
}
