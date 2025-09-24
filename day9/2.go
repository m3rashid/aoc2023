package main

import (
	"fmt"
	"strings"
	"time"
)

func part2(lines []string) {
	timeStart := time.Now()
	var sum int
	for _, line := range lines {
		row := toInts(strings.Fields(line))
		sum += prev(row)
	}

	fmt.Println("Part 2:", sum, "\tin", time.Since(timeStart))
}
