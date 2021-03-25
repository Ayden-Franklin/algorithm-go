package model

import "fmt"

type Node struct{
	Value interface{}
	Next *Node
}

type LinkedList struct{
	Header *Node
	//Tail *Node
	Length int
}

type INode interface {
	Append(node *Node)
	Insert(index int, node *Node)
	Remove(index int)
	IsEmpty() bool
	Index(node *Node) int
}

func (l *LinkedList) IsEmpty() bool{
	return l.Header == nil
}

func (l *LinkedList) Index(node *Node) int{
	index, found := 0, false
	for n := l.Header; n != nil ; n = n.Next {
		if n.Value == node.Value {
			found = true
			break
		}
		index ++
	}
	if found {
		return index
	}else{
		return -1
	}
}
func (l *LinkedList) Append(node *Node){
	for next := l.Header; ; next = next.Next {
		if next.Next == nil {
			next.Next = node
			l.Length = l.Length + 1
			return
		}
	}
}
func (l *LinkedList) Insert(index int, node *Node){
	i := 0
	for n := l.Header; n != nil ; n = n.Next {
		if n.Next == nil{
			n.Next = node
			l.Length = l.Length + 1
			return
		}else if i == index {
			current := n
			next := n.Next
			node.Next = next
			current.Next = node
			l.Length = l.Length + 1
			return
		}
		i ++
	}
}
func (l *LinkedList) Remove(index int) *Node{
	i := 0
	for n := l.Header; n != nil ; n = n.Next {
		if i == index-1 {
			current := n.Next
			n.Next = n.Next.Next
			return current
		}
		i ++
	}
	return nil
}

func reverseLinkedList(list *LinkedList){
	if list == nil || list.Header == nil{
		return
	}
	var current *Node
	next := list.Header
	for next != nil {
		copy := next
		next = next.Next
		copy.Next = current
		current = copy
		//current.next, previous, current = previous, current, current.next
	}
	list.Header = current
}
func main(){
	e1 := new(Node)
	e1.Value = 3
	e2 := new(Node)
	e2.Value = 4
	e1.Next = e2
	e3 := new(Node)
	e3.Value = 5
	e2.Next = e3
	e3.Next = nil

	linkedList := new (LinkedList)
	linkedList.Header = e1
	linkedList.Length = 3

//	nl := new (Node)
//	fmt.Printf("new node value %v\n", e3.next)

	fmt.Println("-----------------------------")
	for n := linkedList.Header; n != nil ; n = n.Next {
		fmt.Printf("Current node values is %v and the next node is %v\n", n.Value, n.Next)
	}

	e4 := new(Node)
	e4.Value = 10
	linkedList.Append(e4)

	fmt.Println("----------------------------- After Append")
	for n := linkedList.Header; n != nil ; n = n.Next {
		fmt.Printf("Current node values is %v and the next node is %v\n", n.Value, n.Next)
	}

	findE3 := linkedList.Index(e3)
	fmt.Printf("Find the node E3 at  %v and the value is %v\n", findE3, e3.Value)




	e5 := new(Node)
	e5.Value = "Inserted"
	linkedList.Insert(2, e5)

	fmt.Println("----------------------------- After Insert")
	for n := linkedList.Header; n != nil ; n = n.Next {
		fmt.Printf("Current node values is %v and the next node is %v\n", n.Value, n.Next)
	}

	linkedList.Remove(2)
	fmt.Println("----------------------------- After Remove")
	for n := linkedList.Header; n != nil ; n = n.Next {
		fmt.Printf("Current node values is %v and the next node is %v\n", n.Value, n.Next)
	}

	reverseLinkedList(linkedList);
	fmt.Println("----------------------------- After Reverse")
	for n := linkedList.Header; n != nil ; n = n.Next {
		fmt.Printf("Current node values is %v and the next node is %v\n", n.Value, n.Next)
	}
}
