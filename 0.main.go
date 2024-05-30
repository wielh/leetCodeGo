package main

import (
	"sort"
	"strconv"
	"time"
)

type Expresion interface {
	result() int
}

type StringMultiplication struct {
	expresion string
}

type StringAddition struct {
	expresion string
}

func (s StringMultiplication) result() int {
	a, _ := strconv.Atoi(s.expresion[0:1])
	b, _ := strconv.Atoi(s.expresion[2:3])
	return a * b
}

func (s StringAddition) result() int {
	a, _ := strconv.Atoi(s.expresion[0:1])
	b, _ := strconv.Atoi(s.expresion[2:3])
	return a + b
}

func getResult(e Expresion) int {
	return e.result()
}

func generateVal(channel chan int, query string) {
	time.Sleep(500 * time.Millisecond)
	val := 0
	if query[1] == '+' {
		val = getResult(StringAddition{query})
	} else {
		val = getResult(StringMultiplication{query})
	}
	channel <- val
}

func solution(queries []string) []int {
	channel := make(chan int)
	for i := 0; i < 5; i++ {
		go generateVal(channel, queries[i])
		queries[i] = ""
	}

	count := 0
	answer := []int{}
	for value := range channel {
		count += 1
		answer = append(answer, value)
		if count >= 5 {
			break
		}
	}
	sort.Ints(answer)
	return answer
}
