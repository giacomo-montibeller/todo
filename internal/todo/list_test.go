package todo

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func CreateList(t *testing.T) {
	list := New()

	assert.NotNil(t, list)
	assert.NotEqual(t, uuid.Nil, list.id)
	assert.Len(t, list.tasks, 0)
}
