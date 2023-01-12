package duplicate

// https://leetcode.com/problems/contains-duplicate/
// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

// Example 1:
// Input: nums = [1,2,3,1]
// Output: true

// Example 2:
// Input: nums = [1,2,3,4]
// Output: false

// Example 3:
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true

func check_map(n []int, map_int map[int]int, index int) {
	_, ok := map_int[n[index]]
	if ok {
		map_int[n[index]]++

	} else {
		map_int[n[index]] = 1
	}
}

func containsDuplicate(n []int) bool {
	map_int := make(map[int]int)
	left := 0
	right := len(n) - 1

	for left < right {
		check_map(n, map_int, left)
		if map_int[n[left]] > 1 {
			return true
		}
		check_map(n, map_int, right)
		if map_int[n[right]] > 1 {
			return true
		}

		left++
		right--
	}
	return false
}
