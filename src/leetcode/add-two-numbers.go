package main

import "fmt"

type ListNode struct {
	    Val int
	    Next *ListNode
	}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	result := &ListNode{0, nil}
	head := result
	carryOver := 0
	for l1 != nil || l2 != nil{
		v := carryOver
		if l1 != nil{
			v += l1.Val
		}
		if l2 != nil{
			v += l2.Val
		}
		carryOver = v / 10
		result.Next = &ListNode { v%10, nil}
		result = result.Next
		if l1 != nil{
			l1 = l1.Next
		}
		if l2 != nil{
			l2 = l2.Next
		}
	}
	if carryOver > 0 {
		result.Next = &ListNode { carryOver, nil}
	}
	return head.Next
}
func main()  {
	list1 := ListNode{2, &ListNode{4, &ListNode{5, nil}}}
	list2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	list := addTwoNumbers(&list1, &list2)
	n := list
	for n != nil {
		fmt.Printf("%v\n", n.Val)
		n = n.Next
	}
}