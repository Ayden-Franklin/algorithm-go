package main

import "fmt"

func partition(array []int, left int, right int) int {
	index := left
	for i := left;i<right;i++ {
		if array[i] < array[right] {
			if i != index {
				array[i], array[index] = array[index], array[i]
			}
			index++
		}
	}
	array[right], array[index] = array[index], array[right]
	return index
}
func quickSort(array []int, start int, end int){
	if start >= end{
		return
	}
	p := partition(array, start, end)
	quickSort(array, start, p-1)
	quickSort(array, p+1, end)
}
func main() {
	array := []int{8,1,6,3,24,26,4,7,97,36,29,83,46}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
