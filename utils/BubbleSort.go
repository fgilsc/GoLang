package main

import "fmt"

func bubbleSort(array []int) []int {
	res := array

	for i := 1; i <= len(res); i++ {
		for j := 0; j < len(res)-i; j++ {
			if res[j] > res[j+1] {
				aux := res[j]
				res[j] = res[j+1]
				res[j+1] = aux
			}
		}
	}
	return res
}

// Testing

func main() {
	b := []int{2, 7, 5, 1, 4, 9, 3, 6, 8, 10}
	a := []int{21, 17, 45, 301, 74, 98, 36, 69, 8, 110}
	s := bubbleSort(b)
	t := bubbleSort(a)
	fmt.Println(s)
	fmt.Println(t)
}
