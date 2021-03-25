package main

import "fmt"

func main(){
	array := []int{ 2,3, 5, 9}
	k := 49
	n := 0
	result := 0

	for i:=len(array) -1; i >= 0;i--{
		d := (k - n) / array[i]
		result += d
		n += array[i]*d

		fmt.Printf("Put int [%v] %v time(s)\n", array[i], d)
	}
	if n < k {
		result = -1
	}

	fmt.Println(result)
}
