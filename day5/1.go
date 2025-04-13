package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var headers = []string{
	"seed-to-soil map:",
	"soil-to-fertilizer map:",
	"fertilizer-to-water map:",
	"water-to-light map:",
	"light-to-temperature map:",
	"temperature-to-humidity map:",
	"humidity-to-location map:",
}

func stringToInt(s string) int64 {
	res, _ := strconv.ParseInt(s, 10, 64)
	return res
}

func parseFile(filename string) ([7][][]int64, []int64) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seeds := make([]int64, 0)

	scanner.Scan() // for the first line
	seedsStr := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	for _, seed := range seedsStr {
		seeds = append(seeds, stringToInt(seed))
	}

	maps := [7][][]int64{}
	for i := range headers {
		maps[i] = make([][]int64, 0)
	}

	currentHeader := -1
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}

		if strings.Contains(txt, ":") {
			currentHeader = slices.Index(headers, strings.TrimSpace(txt))
			continue
		}

		if currentHeader == -1 {
			continue
		}

		nums := strings.Fields(txt)
		numsInt := make([]int64, 0)
		for _, num := range nums {
			if num == "" {
				continue
			}
			numsInt = append(numsInt, stringToInt(num))
		}
		maps[currentHeader] = append(maps[currentHeader], numsInt)
	}

	return maps, seeds
}

func findClosest(seed int64, maps [7][][]int64) int64 {
	toMatch := seed

	var val int64 = 0
	for i := range 7 {
		nums := maps[i]
		var currentVal int64 = -1

		for i := range nums {
			destinationRangeStart := nums[i][0]
			sourceRangeStart := nums[i][1]
			rangeLength := nums[i][2]

			if toMatch < sourceRangeStart || toMatch >= sourceRangeStart+rangeLength {
				continue
			}

			diff := toMatch - sourceRangeStart
			if diff < 0 {
				continue
			}

			currentVal = destinationRangeStart + diff
			break
		}

		if currentVal == -1 {
			val = toMatch
		} else {
			val = currentVal
		}

		toMatch = val
	}

	return val
}

func Solution1() {
	maps, seeds := parseFile("input1.txt")

	var lowest int64 = 1<<63 - 1
	for _, seed := range seeds {
		closest := findClosest(seed, maps)
		lowest = min(lowest, closest)
	}

	fmt.Println("Lowest:", lowest)
}
