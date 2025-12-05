package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func singleNumber(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isValid(s string) bool {
	var stack []rune
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[v] {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}
	return len(stack) == 0
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 && prefix != "" {
			prefix = prefix[:len(prefix)-1]
		}
	}

	return prefix
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)
}

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	cur := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		if start <= cur[1] {
			if end > cur[1] {
				cur[1] = end
			}
		} else {
			res = append(res, []int{cur[0], cur[1]})
			cur[0] = start
			cur[1] = end
		}
	}
	res = append(res, []int{cur[0], cur[1]})

	return res
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		want := target - v
		if idx, ok := m[want]; ok {
			return []int{i, idx}
		}
		m[v] = i
	}
	return nil
}

func main() {
	fmt.Println("tsk1")
}
