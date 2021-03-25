package main
import "fmt"

func main(){
	array := []int{8,1,6,3,24,26,4,7,97,36,29,83,46}
	for i := 0;i< len(array); i ++ {
		min := i
		for k := i + 1;k < len(array); k++ {
			if array[k] < array[min] {
				min = k
			}
		}
		if min !=  i {
			array[i], array[min] = array[min], array[i]
		}
	}
	fmt.Println(array)
}
