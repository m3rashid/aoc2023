package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func Solution1() {
	file, err := os.Open("day1/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grandTotal int
	for scanner.Scan() {
		str := scanner.Text()
		var numbers []int

		for _, ch := range str {
			if unicode.IsDigit(ch) {
				numbers = append(numbers, int(ch)-48)
			}
		}

		if len(numbers) == 0 {
			continue
		}
		grandTotal += numbers[0]*10 + numbers[len(numbers)-1]
	}

	fmt.Println(grandTotal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
