package main

import (
	"fmt"
	"time"
)

func part1(lines []string) {
	start := time.Now()
	galaxies := getExpandedUniverse(lines, 2)

	var sum int
	for i, g1 := range galaxies[1:] {
		for _, g2 := range galaxies[:i+1] {
			dist := abs(g1.x-g2.x) + abs(g1.y-g2.y)
			sum += dist
		}
	}
	fmt.Println("Part 1:", sum, "\tin", time.Since(start))
}
