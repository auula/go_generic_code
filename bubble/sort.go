package main

import (
	"fmt"
)

func bubbleSort(sequence []int64) {
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
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[int8]float64{
		-128: 35.98,
		127:  26.99,
	}

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

	fmt.Println(whoisMin[int64](100, 1000))
}
