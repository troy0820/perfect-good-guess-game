package main

import (
	"flag"
	"strconv"
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
