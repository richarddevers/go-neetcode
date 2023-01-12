package bbb

import "testing"

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

func TestValidAnagram(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		ss := "anagram"
		tt := "nagaram"
		want := true

		got := isValidAnagram(ss, tt)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("two", func(t *testing.T) {
		ss := "rat"
		tt := "car"
		want := false

		got := isValidAnagram(ss, tt)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("three", func(t *testing.T) {
		ss := "aabbccddeeffgg"
		tt := "aabbccddeeffggh"
		want := false

		got := isValidAnagram(ss, tt)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("three", func(t *testing.T) {
		ss := "aabbccddeeffgg"
		tt := "aabbccddeeffggg"
		want := false

		got := isValidAnagram(ss, tt)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("three", func(t *testing.T) {
		ss := "aabbccddeeffgg"
		tt := "abcdefgabcdefg"
		want := true

		got := isValidAnagram(ss, tt)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
