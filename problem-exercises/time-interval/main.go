package main

import (
	"fmt"
)

func main() {

	//timeTuple := [][]int{{1, 2}, {3, 5}, {4, 7}, {7, 9}, {10, 11}}
	//timeTuple := [][]int{{1,3}, {2,4}, {5,7}, {6,8}}
	timeTuple := [][]int{{1,9}, {2,4}, {5,7}, {6,8}}
	
	
	fmt.Println(timeTuple)
	//mergedList := mergeTimes(timeTuple)
	//fmt.Println(mergedList)
	fmt.Println(mergeIntervals(timeTuple))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mergeIntervals(timeTuple [][]int) [][]int {
	mergedList := [][]int{}
	first := timeTuple[0]

	for i := 1; i < len(timeTuple); i++ {
		tt := timeTuple[i]
		fmt.Printf("i=%v, first=%v, tt=%v\n", i, first, tt)

		if first[1] >= tt[0] {
			first[0] = min(first[0], tt[0])
			first[1] = max(first[1], tt[1])
		} else {
			mergedList = append(mergedList, first)
			first = tt
		}
		if i == len(timeTuple)-1 {
			mergedList = append(mergedList, first)
		}
		
	}
	return mergedList
}
