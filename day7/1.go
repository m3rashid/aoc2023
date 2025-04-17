package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	bid  int
	hand string
}

func stringToInt(str string) int {
	res, _ := strconv.Atoi(str)
	return res
}

func parseFile(filename string) []Data {
	f, _ := os.Open(filename)
	defer f.Close()

	data := make([]Data, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := strings.Fields(scanner.Text())
		data = append(data, Data{bid: stringToInt(str[1]), hand: str[0]})
	}

	return data
}

type Pair struct {
	rank    int
	handMap map[string]int
}

var cache = make(map[string]Pair, 0)

func countPairs(hand string) Pair {
	if val, ok := cache[hand]; ok {
		return val
	}

	count := make(map[string]int)
	for i := range hand {
		count[string(hand[i])]++
	}

	pair := Pair{rank: 0, handMap: count}

	// length := len(count)
	// if length == 1 { // (5)
	// 	pair.rank = 7 // 5 of a kind
	// } else if length == 5 { // (1, 1, 1, 1, 1)
	// 	pair.rank = 1 // high card
	// } else if length == 4 { // (1, 1, 1, 2)
	// 	pair.rank = 2
	// } else {
	// 	vals := make([]int, 0)
	// 	for _, v := range count {
	// 		vals = append(vals, v)
	// 	}

	// 	sort.Slice(vals, func(i, j int) bool {
	// 		return vals[i] > vals[j]
	// 	})

	// 	if length == 3 { // (1, 2, 2) or (1, 1, 3)
	// 		if vals[0] == 3 {
	// 			pair.rank = 3
	// 		} else {
	// 			pair.rank = 2
	// 		}
	// 	} else if length == 2 { // (2, 3) or (1, 4)
	// 		if vals[0] == 4 {
	// 			pair.rank = 4
	// 		} else {
	// 			pair.rank = 6
	// 		}
	// 	}
	// }

	cache[hand] = pair
	return pair
}

func compareData(a, b Data) bool {
	aPair, bPair := countPairs(a.hand), countPairs(b.hand)
	fmt.Println(a.hand, aPair.rank, b.hand, bPair.rank)
	if aPair.rank != bPair.rank {
		return aPair.rank < bPair.rank
	}

	return true
}

func Solution1() {
	data := parseFile("sample.txt")
	sort.Slice(data, func(i, j int) bool {
		return compareData(data[i], data[j])
	})

	fmt.Println(data)
}
