package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int{
	cachedMap := make(map[string]int)
	max := 0
	left := 0
	for i:= 0 ; i < len(s); i++ {
		v := string(s[i])
		fmt.Printf("Try to find this value: %v\n", v)
		if vv, ok := cachedMap[v]; ok {
			fmt.Printf("Find this value %v at %v\n", v, vv)
			if vv +1 > left {
				left = vv + 1
			}

		}
		cachedMap[v] = i
		fmt.Printf("------Current [left:%v, right:%v]\n", left, i)

		if i-left +1  > max {
			max = i - left +1
			fmt.Printf("!!!update max value %v]\n", max)
		}
	}
	return max
}
func main()  {
	//fmt.Println(lengthOfLongestSubstring("abcdhwifwfgb"))
	//fmt.Println(lengthOfLongestSubstring("abdeacbe"))
	fmt.Println(lengthOfLongestSubstring("acbcbfghkmbcd"))
	//fmt.Println(lengthOfLongestSubstring("123"))
	//fmt.Println(lengthOfLongestSubstring(""))
	//fmt.Println(lengthOfLongestSubstring( " "))
	//fmt.Println(lengthOfLongestSubstring("bbbb"))
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
