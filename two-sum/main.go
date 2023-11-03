package main

import "fmt"

func twoSum(arr []int, target int) []int {

	arrIndicesMap := make(map[int]int)

	var result []int = make([]int, 2)

	for i := 0; i < len(arr); i++ {
		//arrIndicesMap[arr[i]] = i
		if _, ok := arrIndicesMap[target-arr[i]]; !ok {
			arrIndicesMap[arr[i]] = i
		} else {
			result[0] = arrIndicesMap[target-arr[i]]
			result[1] = i
		}
	}
	return result

}

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}
