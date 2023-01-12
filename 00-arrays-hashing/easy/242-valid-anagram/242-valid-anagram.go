package bbb

import (
	"reflect"
)

// https://leetcode.com/problems/valid-anagram/

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:
// Input: s = "rat", t = "car"
// Output: false

func isValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map_s := make(map[byte]int)
	map_t := make(map[byte]int)

	i := 0
	for i < len(s) {
		value := s[i]
		_, ok := map_s[value]
		if ok {
			map_s[value]++
		} else {
			map_s[value] = 1
		}

		value = t[i]
		_, ok = map_t[value]
		if ok {
			map_t[value]++
		} else {
			map_t[value] = 1
		}
		i++
	}
	return reflect.DeepEqual(map_s, map_t)
}
