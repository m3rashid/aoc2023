package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"unicode"
)

func generateSubsets(str string) []string {
	var subsets []string

	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			subsets = append(subsets, str[i:j+1])
		}
	}
	return subsets
}

func Solution2() {
	count := 0
	grandTotal := 0

	validDigits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var subsets []string
	for _, digit := range validDigits {
		subsets = append(subsets, generateSubsets(digit)...)
	}

	// remove duplicates
	var allValidDaysSubsets []string
	for _, subset := range subsets {
		if !slices.Contains(allValidDaysSubsets, subset) {
			allValidDaysSubsets = append(allValidDaysSubsets, subset)
		}
	}

	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		var numbers []int
		var currentDigitString string

		for _, ch := range str {
			if unicode.IsDigit(ch) {
				numbers = append(numbers, int(ch)-48)
				currentDigitString = ""
			} else {
				currentDigitString += string(ch)

				var maybeValid bool
				for _, subset := range allValidDaysSubsets {
					if currentDigitString == subset {
						maybeValid = true
						break
					}
				}

				if maybeValid {
					for digitNo, digit := range validDigits {
						if currentDigitString == digit {
							numbers = append(numbers, digitNo+1)
							currentDigitString = ""
						}
					}
				} else {
					currentDigitString = currentDigitString[1:]
				}
			}
		}

		count++

		num := numbers[0]*10 + numbers[len(numbers)-1]
		fmt.Println(count, numbers, num)
		grandTotal += num
	}

	fmt.Println(grandTotal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// func Solution2() {
// 	grandTotal := 0

// 	file, err := os.Open("day1/input1.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		str := scanner.Text()
// 		var numbers []int

// 		// for num, digitStr := range numberMap {
// 		// 	if strings.Contains(str, digitStr) {
// 		// 		str = strings.ReplaceAll(str, digitStr, strconv.Itoa(num+1))
// 		// 	}
// 		// }

// 		for word, digit := range numberMap {
// 			if strings.Contains(str, word) {
// 				str = strings.ReplaceAll(str, word, digit)
// 			}
// 		}

// 		for _, ch := range str {
// 			if unicode.IsDigit(ch) {
// 				numbers = append(numbers, int(ch)-'0')
// 			}
// 		}

// 		grandTotal += numbers[0]*10 + numbers[len(numbers)-1]
// 	}

// 	fmt.Println(grandTotal)

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
