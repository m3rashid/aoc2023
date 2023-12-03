package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution2() {
	grandTotal := 0
	count := 1

	file, err := os.Open("day2/input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		if count < 10 {
			str = str[8:]
		} else if count < 100 {
			str = str[9:]
		} else {
			str = str[10:]
		}

		str = strings.ReplaceAll(str, "red", "r")
		str = strings.ReplaceAll(str, "blue", "b")
		str = strings.ReplaceAll(str, "green", "g")

		maxRed := 0
		maxBlue := 0
		maxGreen := 0

		for _, game := range strings.Split(str, "; ") {
			for _, ball := range strings.Split(game, ", ") {
				scoreGame := strings.Split(ball, " ")
				if scoreGame[0] == "" {
					continue
				}

				num, err := strconv.Atoi(scoreGame[0])
				if err != nil {
					fmt.Println("Error converting string to int", count, scoreGame[0])
					panic(err)
				}

				if scoreGame[1] == "r" {
					maxRed = max(maxRed, num)
				} else if scoreGame[1] == "g" {
					maxGreen = max(maxGreen, num)
				} else if scoreGame[1] == "b" {
					maxBlue = max(maxBlue, num)
				}
			}
		}

		grandTotal += (maxRed * maxGreen * maxBlue)
		count += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Grand total:", grandTotal)
}
