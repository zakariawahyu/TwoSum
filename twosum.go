package main

import "fmt"

func twoSum(arr[]int, target int){
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] + arr[j] == target {
				fmt.Println(i, j)
			}
		}
	}
}

func main()  {
	var array = []int {
		0,
		2,
		1,
	}
	twoSum(array, 3)
}
