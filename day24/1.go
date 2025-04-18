package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Hail struct {
	px, py, pz float64
	vx, vy, vz float64
}

func stringToFloat(s string) float64 {
	res, _ := strconv.ParseFloat(s, 64)
	return res
}

func parseFile(filename string) []Hail {
	// format: px py pz @ vx vy vz
	file, _ := os.Open(filename)
	defer file.Close()

	allData := make([]Hail, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}

		fields := strings.Fields(txt)
		vals := make([]float64, 0)
		for _, f := range fields {
			item := strings.Trim(f, ",")
			if item == "" || item == "@" {
				continue
			}
			vals = append(vals, stringToFloat(item))
		}

		allData = append(allData, Hail{
			vals[0], vals[1], vals[2], // px, py, pz
			vals[3], vals[4], vals[5], // vx, vy, vz
		})
	}

	return allData
}

var lowerLimit float64 = 200000000000000
var upperLimit float64 = 400000000000000

func collides(a, b Hail) bool {
	xa := a.vx
	xb := -1 * b.vx
	xc := a.vy
	xd := -1 * b.vy
	xe := b.px - a.px
	xf := b.py - a.py

	den := xa*xd - xb*xc
	if den == 0 {
		return false
	}

	t1 := (xe*xd - xf*xb) / den
	t2 := (xa*xf - xe*xc) / den

	if t1 < 0 || t2 < 0 {
		return false
	}

	x := a.px + t1*a.vx
	y := a.py + t1*a.vy

	if x < lowerLimit || x > upperLimit || y < lowerLimit || y > upperLimit {
		return false
	}

	return true
}

func Solution1() {
	allPoints := parseFile("input1.txt")
	length := len(allPoints)

	numCollisions := 0

	for i := range length {
		for j := i + 1; j < length; j++ {
			if collides(allPoints[i], allPoints[j]) {
				numCollisions++
			}
		}
	}
	println("Number of collisions: ", numCollisions)
}
