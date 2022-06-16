package fizzbuzz

import (
	"strings"
	"testing"
)

var testMap = map[int]string{1: "1", 2: "2", 3: "fizz", 4: "4", 5: "buzz", 15: "fizzbuzz", 30: "fizzbuzz", 33: "fizz", 55: "buzz"}

func TestFizzBuzz(t *testing.T) {
	for number, expected := range testMap {
		fb := strings.ToLower(fizzBuzz(number))
		if fb != expected {
			t.Fatalf("expected %s to be returned on input %d, got %s", expected, number, fb)
		}
	}
}
