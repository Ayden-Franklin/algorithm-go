package main

import "fmt"

func twoSum(nums []int, target int) []int {
	cached := make(map[int]int, len(nums))
	for i:=0; i< len(nums);i++{
		v := target - nums[i]
		if index, ok := cached[v]; ok{
			return []int{index, i}
		}else{
			cached[nums[i]] = i
		}
	}
	return []int{0,0}
}
func main()  {
	array := []int{2,6,7,8,23,46,17,9,35}
	result := twoSum(array, 43)
	fmt.Printf("result is [%v, %v]\n", result[0], result[1])
}
