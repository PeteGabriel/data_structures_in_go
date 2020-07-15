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

func Test_AddData(t *testing.T) {
	is := is.New(t)
	l := New()
	l.AddData("New Node Data")
	is.True(l.node.data == "New Node Data")
}