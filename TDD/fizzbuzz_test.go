package tdd

import (
	"fmt"
	"testing"
)

// Test_FizzBuzz is a test function
//
// FizzBuzz function must return the same number if not a multiple of 3 or 5
// FizzBuzz function must return "Fizz" if the number is a multiple of 3
// FizzBuzz function must return "Buzz" if the number is a multiple of 5
// FizzBuzz function must return "FizzBuzz" if the number is not a multiple of 3 and 5
func Test_FizzBuzz_On_Not_Multiple_Of_Three_Or_Five(t *testing.T) {

	for _, number := range []int{1, 2, 4, 7, 8, 11, 13, 14, 16, 17, 19} {
		if response := FizzBuzz(number); response != fmt.Sprint(number) {
			t.Errorf("Should return %s, got %s", fmt.Sprint(number), response)
		}

	}
}

func Test_FizzBuzz_On_Multiple_Of_Three_And_Not_Of_Five(t *testing.T) {
	for _, number := range []int{3, 6, 9, 12, 18} {
		if response := FizzBuzz(number); response != "Fizz" {
			t.Errorf("Should return Fizz, got %s", response)
		}
	}
}

func Test_FizzBuzz_On_Multiple_Of_Five_And_Not_Of_Three(t *testing.T) {
	for _, number := range []int{5, 10, 20, 25} {
		if response := FizzBuzz(number); response != "Buzz" {
			t.Errorf("Should return Buzz, got %s", response)
		}
	}
}

func Test_FizzBuzz_On_Multiple_Of_Five_And_Of_Three(t *testing.T) {
	for _, number := range []int{15, 30} {
		if response := FizzBuzz(number); response != "FizzBuzz" {
			t.Errorf("Should return FizzBuzz, got %s", response)
		}
	}
}
