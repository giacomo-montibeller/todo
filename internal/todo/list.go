package todo

type List struct {
	id    int
	tasks []string
}

func New() List {
	return List{
		id:    1,
		tasks: []string{},
	}
}
