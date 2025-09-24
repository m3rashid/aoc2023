package main

import (
	"fmt"
	"time"
)

func part2(lines []string) {
	start := time.Now()
	galaxies := getExpandedUniverse(lines, 1e6)

	var sum int
	for i, g1 := range galaxies[1:] {
		for _, g2 := range galaxies[:i+1] {
			dist := abs(g1.x-g2.x) + abs(g1.y-g2.y)
			sum += dist
		}
	}
	fmt.Println("Part 2:", sum, "\tin", time.Since(start))
}
