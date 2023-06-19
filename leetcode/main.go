package main

import (
	"bytes"
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) {
	//b := []int{}
	var b [][]byte
	for i := 0; i < len(strs); i++ {
		var a []byte
		for j := 0; j < len(strs[i]); j++ {
			a = append(a, strs[i][j])
		}
		b = append(b, a)
	}

	dp := make(map[string][]string)
	for i1 := 0; i1 < len(b); i1++ {
		sort.Slice(b[i1], func(i, j int) bool {
			return b[i1][i] < b[i1][j]
		})
		dp[string(b[i1])] = append(dp[string(b[i1])], string(b[i1]))
		fmt.Println(b[i1])
	}
	fmt.Println(dp)
	fmt.Println(bytes.EqualFold(b[0], b[2]))
	fmt.Println(b)
	// sort.Sort(sort.StringSlice)

	// sort.
	fmt.Println(string(b[0]))
	// 	fmt.Println(reflect.DeepEqual(b[0], b[2]))

	// fmt.Println(b)
}

func groupAnagramsa(a []string) {
	dp := make(map[string][]string)
	for i1 := 0; i1 < len(a); i1++ {
		b := []byte(a[i1])
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		// fmt.Println([]byte(a[i1]))
		dp[string(b)] = append(dp[string(b)], a[i1])
		// fmt.Println(b[i1])
	}

	var s [][]string
	for _, v := range dp {
		s = append(s, v)
	}
}

func combinationSum(candidates []int, target int) {
	var a []int
	for i := 0; i < len(candidates); i++ {
		temp := candidates[i]
		for temp <= target {
			a = append(a, candidates[i])
			temp += candidates[i]
		}
	}
	// fmt.Println(a)
	var a1 [][]int
	L := 0
	R := len(candidates) - 1

	for L < R {
		temp := candidates[L] + candidates[R]
		if temp == target {
			a1 = append(a1, []int{candidates[L], candidates[R]})
			// fmt.Println(candidates[L], candidates[R])
		}
		L++
		R--
	}
	for i := 0; i < len(a); i++ {
		temp := 0
		for j := i; j < len(a); j++ {
			temp += a[j]
			if temp > target {
				break
			}
			if temp == target {
				a1 = append(a1, a[i:j+1])
				// fmt.Println(a[i : j+1])
			}
		}
	}
	fmt.Println(a1)
}

func Sum(a []int) int {
	temp := 0
	for _, v := range a {
		temp += v
	}
	return temp
}

func subsets(nums []int) {
	var a [][]int
	a = append(a, []int{})
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			a = append(a, nums[i:j+1])
			// nums[i : j+1])
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return Sum(a[i]) < Sum(a[j])
	})

}

func generateMatrix(n int) {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	var a []int
	for i := 1; i <= n*n; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Printf("%v%v %v", i, j, " ")

			fmt.Println(a)
			matrix[i][j] = a[i+j]
		}
		fmt.Println(" ")
	}
	fmt.Println(matrix)
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return [][]int{}
	} else {
		L := 0
		R := 0
		var b [][]int
		for L < len(firstList) && R < len(secondList) {
			if L == 0 && R == 0 {
				b = append(b, []int{secondList[R][0], firstList[L][1]})
			}
			if R < len(secondList) && L+1 < len(firstList) {
				fmt.Println(firstList[L+1][0])
				fmt.Println(secondList[R][1])
				if firstList[L+1][0] == secondList[R][1] {
					b = append(b, []int{firstList[R][1], firstList[L+1][0]})
					// fmt.Println(firstList[L+1][0])
					// fmt.Println(firstList[R][1])
				}
			}
			L++
			R++

		}
		return b
	}

}

func singleNumber(nums []int) int {

	hashmap := make(map[int]int)
	for _, v := range nums {
		hashmap[v]++
	}
	var a int
	for _, v := range hashmap {
		if v == 1 {
			a = v
		}
	}
	return a
}

func findTheLongestBalancedSubstring(s string) {
	// hashmap := make(map[int][]string)
	// var a []int
	// var s string
	// for i := 0; i < len(s); i++ {
	// 	for j := 0; j < count; j++ {

	// 	}
	// }

	// fmt.Println(hashmap)
	// fmt.Println(count)
}
func Fun(a int) int {
	fmt.Println(a)
	return a

}

func frequencySort(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

func isWinner(player1 []int, player2 []int) int {
	var count1 int
	var count2 int
	for i := 0; i < len(player1); i++ {
		j := i - 1
		j2 := i - 2
		if j >= 0 && player1[j] == 10 {
			count1 += 2 * player1[i]
		} else if j2 >= 0 && player1[j2] == 10 {
			count1 += 2 * player1[i]
		} else {
			count1 += player1[i]
		}
		if j >= 0 && player2[j] == 10 {
			count2 += 2 * player2[i]

		} else if j2 >= 0 && player2[j2] == 10 {
			count2 += 2 * player2[i]
		} else {
			count2 += player2[i]
		}

	}
	if count1 == count2 {
		return 0
	} else if count1 > count2 {
		return 1
	} else {
		return 2
	}
}

func main() {
	isWinner([]int{}, []int{})

	// fmt.Println(frequencySort("cccaaa"))

	// defer Fun(10)
	// defer Fun(20)
	// fmt.Println("OK")

	// fmt.Println("bjjvjb")
	// fmt.Println("rfjnfk")
	// var B = Fun(10)
	// fmt.Println(B(10))

	// findTheLongestBalancedSubstring("01000111")
	// singleNumber([]int{2, 2, 3, 2})
	// a := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	// b := [][]int{{1, 5}, {8, 12}, {15, 24}}
	// a := [][]int{{1, 3}}
	// b := [][]int{{2, 4}}

	// fmt.Println(intervalIntersection(a, b))
	// a := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// candidates, target := []int{2, 3}, 5
	// combinationSum(candidates, target)
	// a := []int{1, 2, 3, 4, 5, 6}
	// subsets(a)

	// generateMatrix(3)

	// groupAnagrams(a)

	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	ch <- 10
	// 	close(ch)
	// }()
	// for v := range ch {
	// 	fmt.Println(v)
	// }
	// wg.Wait()
	// _ = ch

}
