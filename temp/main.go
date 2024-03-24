package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// fmt.Println(firstUniqChar("leetcode"))
	// fmt.Println(isSubsequence("axc", "ahbgdc"))
	// fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{0, 1, 1}))

	fmt.Println(addStrings("11", "123"))

	fmt.Println(spiralOrder([][]int{{2, 5, 8}, {4, 0, -1}}))

	rotate([][]int{{1}})

	rotate([][]int{{1, 2}, {3, 4}})

	// rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	// rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})
}

func firstUniqChar(s string) int {
	chCnt := map[byte]int{}
	length := len(s)
	for i := 0; i < length; i++ {
		fmt.Printf("%v", s[i])
		chCnt[s[i]]++
	}
	for i := 0; i < length; i++ {
		if chCnt[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func isSubsequence(s string, t string) bool {
	sLen, tLen := len(s), len(t)
	i, j := 0, 0
	for i < sLen && j < tLen {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == sLen
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	i, j := 0, 1
	max := 1
	chMap := map[byte]bool{}
	chMap[s[0]] = true
	for j < len(s) {
		if _, ok := chMap[s[j]]; ok {
			// 已经存在
			if j-i > max {
				max = j - i
			}
			// 向右移动i
			for i < j {
				if s[i] == s[j] {
					break
				}
				delete(chMap, s[i])
				i++
			}
			i++
		}
		chMap[s[j]] = true
		j++
	}
	if j == len(s) && j-i > max {
		max = j - i
	}
	return max
}

func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	res := [][]int{}
	length := len(nums)
	for k := 0; k < length-2; k++ {
		i, j := k+1, length-1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if sum == 0 {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				i++
				j--
			} else if sum > 0 {
				i++
			} else {
				j--
			}
		}
	}
	return res
}

func addStrings(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	maxLen := len1
	if len1 < len2 {
		maxLen = len2
	}
	result := ""
	carry := 0
	for i := 0; i < maxLen; i++ {
		ch1 := 0
		if i < len1 {
			ch1 = int(num1[len1-i-1] - '0')
		}
		ch2 := 0
		if i < len2 {
			ch2 = int(num2[len2-i-1] - '0')
		}
		sum := carry + ch1 + ch2
		result = strconv.Itoa(sum%10) + result
		carry = sum / 10
	}
	if carry == 1 {
		result = strconv.Itoa(carry) + result
	}
	return result
}

func spiralOrder(matrix [][]int) []int {
	height := len(matrix)
	width := len(matrix[0])
	left, right, top, bottom := 0, width-1, 0, height-1
	res := []int{}
	i, j := 0, 0
	res = append(res, matrix[0][0])
	for left <= right || top <= bottom {
		for top <= bottom && j < right {
			j++
			res = append(res, matrix[i][j])
		}
		top++
		for left <= right && i < bottom {
			i++
			res = append(res, matrix[i][j])
		}
		right--
		for top <= bottom && j > left {
			j--
			res = append(res, matrix[i][j])
		}
		bottom--
		for left <= right && i > top {
			i--
			res = append(res, matrix[i][j])
		}
		left++
	}
	return res
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for k := i; k < n-i-1; k++ {
			temp := matrix[i][k]
			matrix[i][k] = matrix[n-1-k][i]
			matrix[n-1-k][i] = matrix[n-1-i][n-1-k]
			matrix[n-1-i][n-1-k] = matrix[k][n-1-i]
			matrix[k][n-1-i] = temp
		}
	}
	fmt.Println(matrix)
}
