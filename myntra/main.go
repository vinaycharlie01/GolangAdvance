package main

import (
	"fmt"
)

func ArrSum(a []int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

func combinationSum(candidates []int, target int) {
	var result [][]int
	var a []int
	for i := 0; i < len(candidates); i++ {
		temp := candidates[i]
		for temp <= target {
			a = append(a, candidates[i])
			temp += candidates[i]
		}
	}
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		temp := a[i]

		j := i
		if a[i]+a[j] == target {
			result = append(result, a[i:j+1])
		}
		for temp <= target && j < len(a) {
			fmt.Println(a[i : j+1])
			t := (ArrSum(a[i : j+1]))
			if t == target {
				result = append(result, a[i:j+1])
			}
			temp += a[j]
			j++
		}
	}
	fmt.Println(result)
}

func main() {
	candidates := []int{3, 5, 8}
	target := 11
	combinationSum(candidates, target)

}
