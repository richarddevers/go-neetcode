package duplicate

import "testing"

func TestDuplicate(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		want := true
		n := []int{1, 2, 3, 1}

		got := containsDuplicate(n)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("two", func(t *testing.T) {
		want := false
		n := []int{1, 2, 3, 4}

		got := containsDuplicate(n)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("three", func(t *testing.T) {
		want := true
		n := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

		got := containsDuplicate(n)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
