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

func Test_AddDataToAnExistentList(t *testing.T) {
	is := is.New(t)
	l := New()
	l.AddData("First Node")
	l.AddData("Second Node")

	is.True(l.node.data == "First Node")
	is.True(l.node.next.data == "Second Node")
}

func Test_InsertAtPositionX(t *testing.T) {
	is := is.New(t)
	l := New()
	l.AddData("First Node")
	l.AddData("Second Node")
	l.Insert(2, "New Second Node")

	is.True(l.node.data == "First Node")
	is.True(l.node.next.data == "New Second Node")
	is.True(l.node.next.next.data == "Second Node")
}