package main

import "fmt"
import (
	"learn/src/model"
)

var cacheCapibality int = 10

func updateCache(node *model.Node, list *model.LinkedList){
	index := list.Index(node)
	if index > -1 {
		list.Remove(index)
	}else{
		if list.Length >= cacheCapibality{
			list.Remove(list.Length -1 )
		}
	}
	list.Insert(0, node)
}

func main(){
	e1 := new(model.Node)
	e1.Value = 3
	e2 := new(model.Node)
	e2.Value = 4
	e1.Next = e2
	e3 := &model.Node{Value: 5}
	e3.Value = 5
	e2.Next = e3
	e3.Next = nil


	linkedList := new (model.LinkedList)
	linkedList.Header = e1
	linkedList.Length = 3

	fmt.Println("----------------------------- After Initialization")
	for n := linkedList.Header; n != nil ; n = n.Next {
		fmt.Printf("Current node values is %v and the next node is %v\n", n.Value, n.Next)
	}
	updateCache(&model.Node{ Value: 7}, linkedList)

	fmt.Println("----------------------------- After update 7")
	for n := linkedList.Header; n != nil ; n = n.Next {
		fmt.Printf("Current node values is %v and the next node is %v\n", n.Value, n.Next)
	}

	updateCache(&model.Node{ Value: 5}, linkedList)
	fmt.Println("----------------------------- After update 5")
	for n := linkedList.Header; n != nil ; n = n.Next {
		fmt.Printf("Current node values is %v and the next node is %v\n", n.Value, n.Next)
	}
}
