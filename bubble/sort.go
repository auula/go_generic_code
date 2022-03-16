package main

import (
	"fmt"
)

func bubbleSort(sequence []int64) {

}

func bubbleSortByGeneric[T int64 | float64](sequence []T) {
	for i := 0; i < len(sequence)-1; i++ {
		for j := 0; j < len(sequence)-1-i; j++ {
			if sequence[j] > sequence[j+1] {
				sequence[j], sequence[j+1] = sequence[j+1], sequence[j]
			}
		}
	}
}

func bubbleSortByInterface(sequence []interface{}) {
	switch sequence[0].(type) {

	case int64:

		var arr []int64
		for _, val := range sequence {
			arr = append(arr, val.(int64))
		}

		for i := 0; i < len(arr)-1; i++ {
			for j := 0; j < len(arr)-1-i; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
		fmt.Println(arr)
	case float64:

		var arr []float64
		for _, val := range sequence {
			arr = append(arr, val.(float64))
		}

		for i := 0; i < len(arr)-1; i++ {
			for j := 0; j < len(arr)-1-i; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
		fmt.Println(arr)
	default:
		panic("Type not support")
	}
}

// GenericFunc 一个标准的泛型函数模板
func GenericFunc[T any](args T) {

}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func add[T int64 | float64 | ~int8](a, b T) T {
	return a + b
}

type MyInt int8

type Number interface {
	int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func whoisMin[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	var sequence1 = []int64{100, 23, 14, 66, 78, 12, 8}
	bubbleSortByGeneric(sequence1)
	fmt.Println(sequence1)
	var sequence2 = []float64{120.13, 2.3, 112.3, 66.5, 78.12, 1.2, 8}
	bubbleSortByGeneric(sequence2)
	fmt.Println(sequence2)
}
