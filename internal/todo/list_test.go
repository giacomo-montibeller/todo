package todo

import "testing"
import "github.com/stretchr/testify/assert"

func CreateList(t *testing.T) {
	list := New()

	assert.NotNil(t, list)
	assert.Len(t, list.tasks, 0)
}
