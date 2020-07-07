package unidirectional_linked_list

import (
	"testing"

	"github.com/matryer/is"
)

func Test(t *testing.T) {
	is := is.New(t)
	l := New()
	is.True(l != nil)
}
