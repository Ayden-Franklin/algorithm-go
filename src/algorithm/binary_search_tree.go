package main

import "fmt"
import (
	"learn/src/model"
)
type BinarySearchTree struct {

}
func (tree BinarySearchTree) Find(root *model.TreeNode, data int) *model.TreeNode{
	if root.Val == data {
		return root
	}
	p := root
	for ;p != nil; {
		if p.Val > data {
			p = p.Left
		}else if p.Val < data{
			p = p.Right
		}else {
			break
		}
	}
	return p
}
func (tree BinarySearchTree) Insert(root *model.TreeNode, data int) {
	if root == nil {
		return
	}
	if data < root.Val{
		if root.Left != nil {
			tree.Insert(root.Left, data)
		}else{
			root.Left = &model.TreeNode{Val: data}
		}
	}else {
		if root.Right != nil {
			tree.Insert(root.Right, data)
		}else{
			root.Right = &model.TreeNode{Val: data}
		}
	}
}
func main() {

	array := []int{8,1,6,3,24,26,4,7,97,36,29,83,46}
	var tree model.Tree = BinarySearchTree{}
	root := &model.TreeNode{Val:array[0]}
	for i := 1; i < len(array); i++ {
		tree.Insert(root, array[i])
	}
	//fmt.Printf("Find 6 at %v\n", find(root, 6).Val)
	root.PreOrder(root)
	fmt.Println("--------------------")
	root.InOrder(root)
	fmt.Println("--------------------")
	root.PostOrder(root)
}