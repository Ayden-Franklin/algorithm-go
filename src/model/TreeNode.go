package model

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Tree interface{
	Find(root *TreeNode, data int) *TreeNode
 	Insert(root *TreeNode, data int)
}

func (node TreeNode) PreOrder(root *TreeNode){
	if root == nil{
		return
	}
	fmt.Printf("This node has value :%v\n", root.Val)
	node.PreOrder(root.Left)
	node.PreOrder(root.Right)
}
func (node TreeNode) InOrder(root *TreeNode){
	if root == nil{
		return
	}
	node.InOrder(root.Left)
	fmt.Printf("This node has value :%v\n", root.Val)
	node.InOrder(root.Right)
}
func (node TreeNode) PostOrder(root *TreeNode){
	if root == nil{
		return
	}
	node.PostOrder(root.Left)
	node.PostOrder(root.Right)
	fmt.Printf("This node has value :%v\n", root.Val)
}