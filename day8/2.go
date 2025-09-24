package main

import (
	"fmt"
	"time"
)

func part2(lines []string) {
	timeStart := time.Now()
	dirs := lines[0]
	routes := parseRoutes(lines[2:])

	var next []string
	for k := range routes {
		if k[2] == 'A' {
			next = append(next, k)
		}
	}

	paths := make([]Path, len(next))
	for i, n := range next {
		var j, steps int
		visited := map[string]int{}
		for {
			if n[2] == 'Z' && paths[i].EndPos == 0 {
				paths[i].EndPos = steps
				paths[i].EndVal = n
			}
			key := fmt.Sprintf("%s:%d", n, j)
			if _, ok := visited[key]; ok {
				paths[i].LoopStart = visited[key]
				paths[i].LoopStartVal = key
				paths[i].LoopEnd = steps
				break
			}
			visited[key] = steps
			if _, ok := routes[n]; !ok {
				fmt.Println("Part 2: invalid", n, i, j, steps)
				return
			}

			switch dirs[j] {
			case 'L':
				n = routes[n][0]
			case 'R':
				n = routes[n][1]
			default:
				panic("unknown direction")
			}

			steps++
			j++
			if j >= len(dirs) {
				j = 0
			}
		}
	}

	// AoC task is tuned to have simple solution, so we can just check two things:
	// 1. end position is in the loop
	for i, p := range paths {
		if p.EndPos < p.LoopStart || p.LoopEnd < p.LoopEnd {
			fmt.Println("Part 2: end position is not in loop", i, p.EndPos, p.LoopStart, p.LoopEnd)
			return
		}
	}
	// 2. end position is equal to loop length, so we can just use lcm
	for i, p := range paths[1:] {
		if p.EndPos != p.LoopEnd-p.LoopStart {
			fmt.Println("Part 2: end position is not equal to loop length", i, p.EndPos, p.LoopEnd-p.LoopStart)
			return
		}
	}

	nums := make([]int, len(paths))
	for i, p := range paths {
		nums[i] = p.EndPos
	}
	fmt.Println("Part 2:", lcm(nums...), "\tin", time.Since(timeStart))
}
