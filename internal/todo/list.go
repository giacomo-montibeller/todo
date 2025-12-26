package todo

import "github.com/google/uuid"

type List struct {
	id    uuid.UUID
	tasks []string
}

func New() List {
	return List{
		id:    uuid.New(),
		tasks: []string{},
	}
}
