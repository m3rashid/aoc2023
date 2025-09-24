package main

import (
	"fmt"
	"strings"
	"time"
)

func part1(lines []string) {
	timeStart := time.Now()
	var sum int
	for _, line := range lines {
		row := toInts(strings.Fields(line))
		sum += next(row)
	}

	fmt.Println("Part 1:", sum, "\tin", time.Since(timeStart))
}
