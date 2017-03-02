package main

type list struct {
	value interface{}
	rest  *list
}

func newList(v interface{}, r *list) *list {
	return &list{v, r}
}

func (l *list) Rest() *list {
	l.rest = newList(42, nil)
	return l.rest
}

func main() {
	l := newList(42, nil)

	for {
		l = l.Rest()
	}
}
