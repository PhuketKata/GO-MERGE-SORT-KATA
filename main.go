package main

import (
	"fmt"
)

func merge(left []int, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	middle := len(slice) / 2
	return merge(mergeSort(slice[:middle]), mergeSort(slice[middle:]))
}

func main() {
	arr := []int{23, 45, 5, 1, 98, 23}
	x := mergeSort(arr)
	fmt.Println(x)
}
