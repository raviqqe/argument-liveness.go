package main

type list struct {
	value interface{}
	rest  *list
}

func newList(v interface{}, r *list) list {
	return list{v, r}
}

func (l *list) Rest() *list {
	r := newList(42, nil)
	l.rest = &r
	return l.rest
}

func foo(l list) {
	for {
		l = *l.Rest()
	}
}

func main() {
	foo(newList(42, nil))
}
