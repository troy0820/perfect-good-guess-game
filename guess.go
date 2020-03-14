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
		} else if checkIncludes(n, g[i]) {
			s = append(s, "good")
		} else {
			s = append(s, "bad")
		}
	}
	return s
}

type number int

type Num struct {
	number
}

func (n *Num) Set(s string) error {
	num, _ := strconv.Atoi(s)
	n.number = number(num)
	return nil
}

func (n number) String() string {
	return strconv.Itoa(int(n))
}

func NumFlag(s string, n number, usage string) *number {
	f := Num{n}
	flag.CommandLine.Var(&f, s, usage)
	return &f.number
}

func main() {
	//Flag for amount of numbers you want to guess
	number := NumFlag("number", 3, "a number")
	flag.Parse()
	if int(*number) < 3 {
		fmt.Println("At least try 3 numbers")
		os.Exit(1)
	}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	var nums numbers
	//Random numbers in a slice
	for i := 0; i < int(*number); i++ {
		nums = append(nums, r.Intn(9))
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		//Get input from stdin
		fmt.Println("Enter your guess: ")
		text, _ := reader.ReadString('\n')
		var guess numbers
		text = strings.Replace(text, " ", "", -1) //strip white space from text
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
			fmt.Println("You Got It!")
			break
		}
		fmt.Println("Your guess", guess)
		fmt.Println("answer", answer)
		//		fmt.Println("this is answer", answer)
	}
}
