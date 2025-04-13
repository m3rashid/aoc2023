package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	time     int
	distance int
}

func stringToInt(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}

func parseFile(filename string) []Data {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	timeStrs := strings.Fields(strings.Split(scanner.Text(), ":")[1])

	scanner.Scan()
	distanceStrs := strings.Fields(strings.Split(scanner.Text(), ":")[1])

	data := make([]Data, len(timeStrs))
	for i := range len(timeStrs) {
		data[i] = Data{
			time:     int(stringToInt(timeStrs[i])),
			distance: int(stringToInt(distanceStrs[i])),
		}
	}

	return data
}

func process(data Data) int {
	ret := 0
	for i := range data.time {
		speed := i
		time := data.time - i
		distance := speed * time

		if distance > data.distance {
			ret++
		}
	}

	return ret
}

func Solution1() {
	data := parseFile("input1.txt")

	val := 1
	for _, d := range data {
		res := process(d)
		fmt.Println(d, res)
		val *= res
	}

	fmt.Println("Result: ", val)
}
