package unidirectional_linked_list

type List struct{
	node* Node
	len int
}

type Node struct {
	data string
	next *Node
}

func New() *List {
	return &List{}
}

func (l *List) size() int {
	return l.len
}

/**
 Add new data.
 */
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

	l.len += 1
}

/**
 Insert new data at given position.
 If given position is bigger than current size, just add to front of list.
*/
func (l *List) Insert(pos int, data string){
	if pos > l.size() {
		l.AddData(data)
		return
	}

	n := &Node{
		data:data,
		next:nil,
	}

	i := 1
	for l.node.next != nil {
		if i == (pos-1) {
			n.next = l.node.next
			l.node.next = n
			l.len += 1
			return
		}
	}
}