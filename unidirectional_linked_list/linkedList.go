package unidirectional_linked_list

type List struct{
	node* Node
}

type Node struct {
	data string
	next *Node
}

func New() *List {
	return &List{}
}