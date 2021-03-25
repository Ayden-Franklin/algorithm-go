package main

import "fmt"

func removeDuplicates(nums []int) int {
	l:=0
	r:=1
	for r< len(nums){
	 	fmt.Printf("------test [left:%v, right:%v]\n", l, r)
		if nums[l] == nums[r]{
			r++
		}else {
			fmt.Printf("------test complete [left:%v, right:%v]\n", l, r)

			nums[l+1] = nums[r]

			l++
			r++
		}
	}
	return l+1
};
func main()  {
	array := []int{0,0,1,1,1,2,2,3,3,4,5,6,6,6,7,7,8,9}
	//array := []int{0,1,1}
	result := removeDuplicates(array)
	fmt.Printf("result is [%v]\n", result)
	for i:=0;i< result;i++ {
		fmt.Println(array[i])
	}
}
