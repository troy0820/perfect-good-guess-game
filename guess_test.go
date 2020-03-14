package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIncludes(t *testing.T) {
	t.Parallel()
	answer := 3
	guess := []int{1, 1, 3}

	t.Run("Test if checkIncludes returns true", func(t *testing.T) {
		if want, got := true, checkIncludes(guess, answer); got != want {
			t.Errorf("Error: This is what we want %+v, and this is what you got %+v", want, got)
		}
	})
	wrong := 5
	t.Run("Test if checkincludes returns false", func(t *testing.T) {
		if want, got := false, checkIncludes(guess, wrong); got != want {
			t.Errorf("Error: This is what we want %+v, and this is what we got %+v", want, got)
		}
	})
}

func TestCheckAnswer(t *testing.T) {
	t.Parallel()
	guess := []int{3, 4, 5}
	testGuess := []int{3, 4, 5}
	answer := []string{"perfect", "perfect", "perfect"}

	t.Run("Test if answer is true", func(t *testing.T) {
		if want, got := true, reflect.DeepEqual(answer, checkAnswer(guess, testGuess)); got != want {
			t.Errorf("Error: This is what we want %+v, and this is what you got %+v", answer, checkAnswer(guess, testGuess))
		}
	})
	wrong := []int{6, 7, 8}
	t.Run("Test if answer is false", func(t *testing.T) {
		if want, got := false, reflect.DeepEqual(answer, checkAnswer(guess, wrong)); got != want {
			t.Errorf("Error: This is what we want %+v, and this is what you got %+v", answer, checkAnswer(guess, wrong))
		}
	})

	includes := []int{4, 4, 4}
	t.Run("Test if check includes", func(t *testing.T) {
		assert.Equal(t, checkAnswer(guess, includes), []string{"good", "perfect", "good"}, "These are not equal")
	})
}
