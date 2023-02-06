package semver

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	passingTest := []struct {
		a, b string
		want int
	}{
		{"1.0", "1.1.1", -1},
		{"1.0", "1.0.0", 0},
		{"1.4", "1.9.1", -1},
		{"1.1", "1.0.1", 1},
		{"1.1", "1.0", 1},
		{"1.1", "1.1", 0},
		{"0.1", "1.1.0", -1},
		{"1.9.9", "0.1", 1},
	}
	failingTest := []struct {
		a, b string
	}{
		{"aab,ae", "1.1.1"},
		{"1.0", ""},
		{"", "1.0"},
		{"", ""},
		{"abc", "abc"},
	}
	fmt.Printf("passingTest: %v \n", passingTest)
	for _, tt := range passingTest {
		got, err := Compare(tt.a, tt.b)
		if got != tt.want && err == nil {
			t.Errorf("Compare(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
	fmt.Printf("failingTest: %v \n", failingTest)
	for _, tt := range failingTest {
		_, err := Compare(tt.a, tt.b)
		if err == nil {
			t.Errorf("Compare(%q, %q) = %s, want Error", tt.a, tt.b, err)
		}
	}
}
