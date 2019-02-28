package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type numbers []int

func checkIncludes(n numbers, g int) bool {
	for i := 0; i < len(n); i++ {
		if n[i] == g {
			return true
		}
	}
	return false
}

func checkAnswer(n numbers, g numbers) []string {
	s := []string{}
	for i := 0; i < len(n); i++ {
		if n[i] == g[i] {
			s = append(s, "perfect")
		} else if checkIncludes(n, g[i]) == true {
			s = append(s, "good")
		} else {
			s = append(s, "bad")
		}
	}
	return s
}

func main() {
	number := flag.Int("number", 3, "a number")
	flag.Parse()
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	var nums numbers
	for i := 0; i < *number; i++ {
		nums = append(nums, r.Intn(9))
	}
	reader := bufio.NewReader(os.Stdin)
	loop := true
	for loop == true {
		fmt.Println("Enter your guess: ")
		text, _ := reader.ReadString('\n')
		var guess numbers
		text = strings.Replace(text, " ", "", -1)
		for _, v := range text {
			num, _ := strconv.Atoi(string(v))
			guess = append(guess, num)
		}
		guess = guess[:len(guess)-1]
		answer := checkAnswer(nums, guess)
		perfect := 0
		for _, v := range answer {
			if "perfect" == string(v) {
				perfect++
			}
		}
		if perfect == len(answer) {
			fmt.Println("got it")
			break
		}
		fmt.Println("Your guess", guess)
		fmt.Println("answer", answer)
	}
}
