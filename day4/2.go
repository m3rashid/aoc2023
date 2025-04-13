package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func handleCardlinePart2(str string) int {
	strs := strings.Split(str, "|")
	winningNumbers := strings.Fields(strings.Split(strs[0], ":")[1])
	numbersIHave := strings.Fields(strs[1])

	count := 0
	for _, num := range winningNumbers {
		if slices.Contains(numbersIHave, num) {
			count++
		}
	}

	return count
}

func Solution2() {
	file, _ := os.Open("input1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	results := make([]int, 0)
	for scanner.Scan() {
		results = append(results, handleCardlinePart2(scanner.Text()))
	}

	numberOfLines := len(results)

	decks := make([]int, numberOfLines)
	for i := range numberOfLines {
		decks[i] = 1
	}

	res := 0
	for i := range numberOfLines {
		currentRes := results[i]
		numDecks := decks[i]
		res += numDecks

		for j := range currentRes {
			if i+j+1 < numberOfLines {
				decks[i+j+1] += numDecks
			}
		}
	}

	fmt.Println("Result:", res)
}
