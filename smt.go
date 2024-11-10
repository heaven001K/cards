package main

import "fmt"

func getAvg(arr []int) {
	n := len(arr)
	result := 0
	for _, value := range arr {
		result += value

	}
	result /= n
	fmt.Println(result)
}
