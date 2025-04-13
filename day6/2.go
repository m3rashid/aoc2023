package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt64(str string) int64 {
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}

func parseFile2(filename string) (int64, int64) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	timeStr := strings.Join(strings.Fields(strings.Split(scanner.Text(), ":")[1]), "")

	scanner.Scan()
	distanceStr := strings.Join(strings.Fields(strings.Split(scanner.Text(), ":")[1]), "")

	return stringToInt64(timeStr), stringToInt64(distanceStr)
}

func processFile2(time, distance int64) int64 {
	var ret int64 = 0
	for i := range time {
		s := i
		t := time - i
		d := s * t

		if d > distance {
			ret++
		}
	}

	return ret

}

func Solution2() {
	time, distance := parseFile2("input1.txt")
	result := processFile2(time, distance)
	fmt.Println("Result:", result)
}
