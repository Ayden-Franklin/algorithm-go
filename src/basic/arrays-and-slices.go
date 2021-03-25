package main

func SumArray(a [5]int) int{
	result := 0
	/*for i:= 0; i < len(a); i ++ {
		result += a[i]
	}*/
	for _, number := range a{
		result += number
	}
	return result
}

func SumSlice(a []int) (int, int){
	result := 0
	for _, number := range a{
		result += number
	}
	return len(a), result
}