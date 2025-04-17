package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	px, py, pz int
	vx, vy, vz int
}

func stringToInt(s string) int {
	// res, _ := strconv.ParseInt(s, 10, 64)
	// return res
	res, _ := strconv.Atoi(s)
	return res
}

func parseFile(filename string) []Point {
	// format: px py pz @ vx vy vz
	// px, py, pz = position
	// vx, vy, vz = velocity (distance moved per nanosecond)
	file, _ := os.Open(filename)
	defer file.Close()

	allData := make([]Point, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}

		fields := strings.Fields(txt)
		vals := make([]int, 0)
		for _, f := range fields {
			item := strings.Trim(f, ",")
			if item == "" || item == "@" {
				continue
			}
			vals = append(vals, stringToInt(item))
		}

		allData = append(allData, Point{
			px: vals[0],
			py: vals[1],
			pz: vals[2],
			vx: vals[3],
			vy: vals[4],
			vz: vals[5],
		})
	}

	return allData
}

func collides(p1, p2 Point) bool {
	//
}

func Solution1() {
	allPoints := parseFile("sample.txt")
	fmt.Println("allData: ", allPoints)
}
