package main

import (
	"reflect"
	"testing"
)

func TestCheckIncludes(t *testing.T) {
	answer := 3
	guess := []int{1, 1, 3}

	t.Run("Test if checkIncludes returns true", func(t *testing.T) {
		if want, got := true, checkIncludes(guess, answer); got != want {
			t.Errorf("Error: This is what you we want %+v, and this is what you got %+v", want, got)
		}
	})
}

func TestCheckAnswer(t *testing.T) {
	guess := []int{3, 4, 5}
	testGuess := []int{3, 4, 5}
	answer := []string{"perfect", "perfect", "perfect"}

	t.Run("Test if answer is true", func(t *testing.T) {
		if want, got := true, reflect.DeepEqual(answer, checkAnswer(guess, testGuess)); got != want {
			t.Errorf("Error: This is what we want %+v, and this is what you got %+v", answer, checkAnswer(guess, testGuess))
		}
	})
}
