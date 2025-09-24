package main

import (
	"fmt"
	"time"
)

func part1(lines []string) {
	timeStart := time.Now()
	posStart := findStart(lines)

	// find connections
	startDirs := findStartDirections(lines, posStart)
	if len(startDirs) != 2 {
		panic("Start has to have 2 directions")
	}

	var steps int
	pos := posStart
	dir := startDirs[0]
	for {
		pos = pos.Add(dir)
		steps++
		if pos == posStart {
			break
		}
		dir = pos.GetNext(dir, lines)
	}
	if steps%2 != 0 {
		panic("Steps has to be even")
	}
	fmt.Println("Part 1:", steps/2, "\tin", time.Since(timeStart))
}
