package main

import "fmt"

func main(){
	array := []int{8,1,6,3,24,26,4,7,97,36,29,83,46}
	for i := 1;i< len(array); i ++ {
		for k := i;k> 0 && array[k] < array[k-1]; k-- {
			array[k],array[k-1] = array[k-1], array[k]
		}
	}
	fmt.Println(array)
}
