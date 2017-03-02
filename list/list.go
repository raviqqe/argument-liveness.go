package list

type List struct {
	value interface{}
	rest  *List
}

func NewList(v interface{}, r *List) *List {
	return &List{v, r}
}

func (l *List) Rest() *List {
	l.rest = NewList(42, nil)
	return l.rest
}
