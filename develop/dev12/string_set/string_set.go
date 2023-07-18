package string_set

type stringSet struct {
	data map[string]struct{}
}

func New(init ...string) stringSet {
	ret := stringSet{data: make(map[string]struct{}, len(init))}
	for _, s := range init {
		ret.Add(s)
	}
	return ret
}

func (ss stringSet) Add(s string) {
	ss.data[s] = struct{}{}
}

func (ss stringSet) Remove(s string) {
	delete(ss.data, s)
}

func (ss stringSet) Contains(s string) bool {
	_, ok := ss.data[s]
	return ok
}

func (ss stringSet) Slice() []string {
	ret := make([]string, 0, len(ss.data))
	for s := range ss.data {
		ret = append(ret, s)
	}
	return ret
}
