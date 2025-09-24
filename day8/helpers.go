package main

import (
	"regexp"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

var reRoute = regexp.MustCompile(`^(\w+) = \((\w+), (\w+)\)$`)

func parseRoutes(lines []string) map[string][2]string {

	routes := map[string][2]string{}
	for _, line := range lines {
		m := reRoute.FindStringSubmatch(line)
		if m == nil {
			panic("invalid route")
		}
		routes[m[1]] = [2]string{m[2], m[3]}
	}
	return routes
}

type Path struct {
	EndPos int
	EndVal string

	LoopStart    int
	LoopEnd      int
	LoopStartVal string

	// Cur
}

func gcd(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	a, b := nums[0], gcd(nums[1:]...)
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	a, b := nums[0], lcm(nums[1:]...)
	return a * b / gcd(a, b)
}
