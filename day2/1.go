package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution1() {
	grandTotal := 0
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	count := 1

	file, err := os.Open("input1.txt")
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

		skipIter := false
		for game := range strings.SplitSeq(str, "; ") {
			for ball := range strings.SplitSeq(game, ", ") {
				scoreGame := strings.Split(ball, " ")
				if scoreGame[0] == "" {
					continue
				}

				num, err := strconv.Atoi(scoreGame[0])
				if err != nil {
					fmt.Println("Error converting string to int", count, scoreGame[0])
					panic(err)
				}

				if scoreGame[1] == "r" && num > maxRed {
					skipIter = true
					grandTotal += count
					break
				} else if scoreGame[1] == "g" && num > maxGreen {
					skipIter = true
					grandTotal += count
					break
				} else if scoreGame[1] == "b" && num > maxBlue {
					skipIter = true
					grandTotal += count
					break
				}
			}

			if skipIter {
				skipIter = false
				break
			}
		}

		count += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count*(count-1)/2 - grandTotal)
}
