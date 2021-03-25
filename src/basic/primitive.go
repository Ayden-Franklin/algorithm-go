package main

func Add(first int, second int) int{
	return first + second
}
func Repeat(l string, times int) string{
	var result string
	for i := 0; i < times; i ++ {
		result += l
	}
	return result
}
