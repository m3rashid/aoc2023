package main

import (
	"fmt"
	"time"
)

func part1(lines []string) {
	timeStart := time.Now()
	path := lines[0]
	routes := parseRoutes(lines[2:])

	var i, steps int
	next := "AAA"

	visited := map[string]bool{}
	for next != "ZZZ" {
		key := fmt.Sprintf("%s:%d", next, i)
		if visited[key] {
			fmt.Println("Part 1: loop", next, i, steps)
			return
		}

		visited[key] = true
		if _, ok := routes[next]; !ok {
			fmt.Println("Part 1: invalid", next, i, steps)
			return
		}

		switch path[i] {
		case 'L':
			next = routes[next][0]
		case 'R':
			next = routes[next][1]
		default:
			panic("unknown direction")
		}

		steps++
		i++

		if i >= len(path) {
			i = 0
		}
	}

	fmt.Println("Part 1:", steps, "\tin", time.Since(timeStart))
}
