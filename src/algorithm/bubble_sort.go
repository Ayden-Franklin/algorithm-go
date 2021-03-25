package main

import "fmt"

func main(){
	array := []int{8,1,6,3,24,26,4,7,97,36,29,83,46}
	for i := 0;i< len(array); i++ {
		flag := false
		for k := 0;k< len(array)-i-1; k++ {
			if array[k] > array[k+1] {
				array[k],array[k+1] = array[k+1],array[k]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println(array)
}
