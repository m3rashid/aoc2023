package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Label int
type Type int

const (
	Labels  = "23456789TJQKA"
	LabelsJ = "J23456789TQKA"

	TypeHigh Type = iota
	TypePair
	TypeTwoPair
	TypeThree
	TypeFullHouse
	TypeFour
	TypeFive
)

type Hand struct {
	cards []Label
	bid   int
	typ   Type
}

func getTypeJ(cards []Label) Type {
	types := make(map[Label]int)
	var jokers int
	for _, c := range cards {
		if c == 0 {
			jokers++
			continue
		}
		types[c]++
	}
	switch len(types) {
	case 0, 1:
		return TypeFive
	case 2:
		for _, v := range types {
			if v+jokers == 4 {
				return TypeFour
			}
		}
		return TypeFullHouse
	case 3:
		for _, v := range types {
			if v+jokers == 3 {
				return TypeThree
			}
		}
		return TypeTwoPair
	case 4:
		return TypePair
	case 5:
		return TypeHigh
	}
	panic(fmt.Sprintf("invalid hand: %v", cards))
}

func getType(cards []Label) Type {
	types := make(map[Label]int)
	for _, c := range cards {
		types[c]++
	}
	switch len(types) {
	case 1:
		return TypeFive
	case 2:
		for _, v := range types {
			if v == 4 {
				return TypeFour
			}
		}
		return TypeFullHouse
	case 3:
		for _, v := range types {
			if v == 3 {
				return TypeThree
			}
		}
		return TypeTwoPair
	case 4:
		return TypePair
	case 5:
		return TypeHigh
	}
	panic(fmt.Sprintf("invalid hand: %v", cards))
}

func parseHandJ(line string) *Hand {
	fields := strings.Fields(line)
	hand, bidStr := fields[0], fields[1]

	labels := make([]Label, len(hand))
	for i, card := range hand {
		labels[i] = Label(strings.IndexRune(LabelsJ, card))
	}
	bid, err := strconv.Atoi(bidStr)
	if err != nil {
		panic(err)
	}

	return &Hand{
		cards: labels,
		bid:   bid,
		typ:   getTypeJ(labels),
	}
}

func parseHand(line string) *Hand {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		return nil
	}

	hand, bidStr := fields[0], fields[1]

	labels := make([]Label, len(hand))
	for i, card := range hand {
		labels[i] = Label(strings.IndexRune(Labels, card))
	}
	bid, err := strconv.Atoi(bidStr)
	if err != nil {
		panic(err)
	}

	return &Hand{
		cards: labels,
		bid:   bid,
		typ:   getType(labels),
	}
}
