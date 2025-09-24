package main

import (
	"fmt"
	"time"
)

func part2(lines []string) {
	timeStart := time.Now()
	posStart := findStart(lines)

	// find connections
	startDirs := findStartDirections(lines, posStart)
	if len(startDirs) != 2 {
		panic("Start has to have 2 directions")
	}

	// create the path
	mmap := createPath(lines, posStart, startDirs)

	// mark left/right sides
	pos := posStart
	dir := startDirs[0]
	for {
		fill(mmap, pos.Add(dir.turnLeft()), ' ', '<')
		fill(mmap, pos.Add(dir.turnRight()), ' ', '>')
		pos = pos.Add(dir)
		fill(mmap, pos.Add(dir.turnLeft()), ' ', '<')
		fill(mmap, pos.Add(dir.turnRight()), ' ', '>')
		if pos == posStart {
			break
		}
		mmap[pos.y][pos.x] = getLineChar(rune(lines[pos.y][pos.x]))
		dir = pos.GetNext(dir, lines)
	}

	// count charInside
	charInside := '>'
	charOutside := '<'
	if mmap[0][0] == '>' {
		charInside = '<'
		charOutside = '>'
	}
	var inside, outside, nonMarked int
	for y, line := range mmap {
		for x, char := range line {
			switch char {
			case charInside:
				inside++
				mmap[y][x] = '#'
			case charOutside:
				outside++
				mmap[y][x] = ' '
			case ' ':
				nonMarked++
				mmap[y][x] = 'X'
			}
		}
	}
	plot(mmap)
	if nonMarked != 0 {
		panic(fmt.Sprint("Non marked:", nonMarked))
	}

	fmt.Println("Part 2:", inside, "\tin", time.Since(timeStart))
}
