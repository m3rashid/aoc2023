package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func handleCardlinePart1(str string) float64 {
	numsStr := strings.Trim(strings.Split(str, ":")[1], " ")
	reqNums := strings.Split(numsStr, "|")

	mineNums := strings.Split(strings.Trim(reqNums[1], " "), " ")
	winingNums := strings.Split(strings.Trim(reqNums[0], " "), " ")

	newMineNums := []string{}
	for i := range mineNums {
		mineNums[i] = strings.Trim(mineNums[i], " ")
		if mineNums[i] != "" {
			newMineNums = append(newMineNums, mineNums[i])
		}
	}
	mineNums = newMineNums

	newWinningNums := []string{}
	for i := range winingNums {
		winingNums[i] = strings.Trim(winingNums[i], " ")
		if winingNums[i] != "" {
			newWinningNums = append(newWinningNums, winingNums[i])
		}
	}
	winingNums = newWinningNums

	numMatch := 0
	for _, num := range mineNums {
		if slices.Contains(winingNums, num) {
			numMatch++
		}
	}

	if numMatch > 0 {
		return math.Pow(2, float64(numMatch)-1)
	}

	return 0
}

func Solution1() {
	file, _ := os.Open("input2.txt")
	defer file.Close()

	var ans float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ans += handleCardlinePart1(scanner.Text())
	}

	fmt.Println("Answer:", ans)
}
