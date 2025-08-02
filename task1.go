package main

import (
	"fmt"
	"sort"
	"strconv"
)

// leetcode 136 只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

func mainTask1() {
	nums := []int{4, 1, 2, 1, 2}
	result := singleNumber(nums)
	println(result)
}

// leetcode9 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	str := strconv.Itoa(x)
	l := len(str)
	arr := make([]int, l)
	i := l - 1
	for x > 0 {
		arr[i] = x % 10
		x = x / 10
		i--
	}
	sz := len(arr)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}

func mainTask2() {
	fmt.Println(isPalindrome1(121))
}

// 有效括号 leetcode20
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (c == ')' && top != '(') || (c == ']' && top != '[') || (c == '}' && top != '{') {
				return false
			}
		}
	}
	return len(stack) == 0
}

func mainTask3() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(("))
	fmt.Println(isValid("(]"))
}

// 最小公共前缀 leetcode14
func lognestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}
	return prefix
}

func mainTask4() {
	fmt.Println(lognestCommonPrefix([]string{"flower", "flow", "flight"}))
}

// 加一 leetcode66
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
	}
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	return newDigits
}

func mainTask5() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}

// 删除有序数组中的重复项 leetcode26
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

func mainTask6() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

// 合并区间 leetcode56
func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			a := intervals[i]
			b := intervals[j]
			if a[1] >= b[0] && a[0] <= b[1] {
				// 有重叠
				arr := []int{
					a[0],
					a[1],
					b[0],
					b[1],
				}
				sort.Ints(arr)

				intervals[i] = []int{arr[0], arr[3]}
				// delete j
				intervals = append(intervals[:j], intervals[j+1:]...)
				j = i
			}
		}
	}
	return intervals
}

func mainTask11() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

// 两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
