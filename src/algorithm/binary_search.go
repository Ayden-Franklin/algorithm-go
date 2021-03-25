package main

import "fmt"

func binary_search(array []int, target int) int{
	length := len(array)
	left, right := 0, length-1
	for left <= right {
		middle := left + (right-length)/2
		if array[middle] > target{
			right = middle -1
		}else if array[middle] < target{
			left = middle + 1
		}else{
			return middle
		}
	}
	return -1
}

func binary_search_first_value(array []int, target int) int{
	length := len(array)
	left, right := 0, length-1
	for left <= right{
		middle := left + (right-length)/2
		if array[middle] > target{
			right = middle -1
		}else if array[middle] < target{
			left = middle + 1
		}else{
			for ; array[middle] == target;middle-- {
			}
			middle++
			return middle
		}
	}
	return -1
}
func binary_search_last_value(array []int, target int) int{
	length := len(array)
	left, right := 0, length-1
	for left <= right{
		middle := left + (right-length)/2
		if array[middle] > target{
			right = middle -1
		}else if array[middle] < target{
			left = middle + 1
		}else{
			for ; array[middle] == target;middle++ {
			}
			middle--
			return middle
		}
	}
	return -1
}
func main() {
	//array := []int{1, 3, 4, 6, 7, 8, 24, 26, 29, 36, 46, 83, 97}
	//fmt.Printf("Find 83 at %v\n", binary_search(array, 83))
	array := []int{1, 3, 3, 4, 6, 6,  7, 8, 24, 26, 26, 29, 36, 46, 83, 97}
	fmt.Printf("Find 6 at %v\n", binary_search_last_value(array, 6))
}
