package main

import (
	"fmt"
	"sort"
	"time"
)

func Solution2(lines []string) {
	timeStart := time.Now()
	hands := []*Hand{}
	for _, line := range lines {
		if line == "" || line == "\n" {
			continue
		}
		res := parseHandJ(line)
		if res != nil {
			hands = append(hands, res)
		}
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].typ != hands[j].typ {
			return hands[i].typ < hands[j].typ
		}
		for i, c := range hands[i].cards {
			if c != hands[j].cards[i] {
				return c < hands[j].cards[i]
			}
		}
		panic("duplicate hands")
	})

	var sum int
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}

	fmt.Println("Part 2:", sum, "\tin", time.Since(timeStart))
}
