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

func (l *List) AddData(data string){
	n := &Node{
		data:data,
		next:nil,
	}
	if l.node == nil {
		l.node = n
	}else {
		l.node.next = n
	}
}