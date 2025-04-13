package main

import (
	"fmt"
	"sync"
)

func Solution2() {
	maps, seeds := parseFile("input1.txt")
	values := make([]int64, len(seeds)/2)

	wg := sync.WaitGroup{}
	for i := range len(seeds) / 2 {
		seedStart := seeds[i*2]
		seedRange := seeds[i*2+1]

		wg.Add(1)
		go func(i int, seedStart, seedRange int64) {
			defer wg.Done()
			var res int64 = 1<<63 - 1
			for j := range seedRange {
				closest := findClosest(seedStart+j, maps)
				res = min(res, closest)
			}
			values[i] = res
		}(i, seedStart, seedRange)
	}
	wg.Wait()

	var lowest int64 = 1<<63 - 1
	for i, v := range values {
		fmt.Println(i, v)
		lowest = min(lowest, v)
	}

	fmt.Println("Lowest:", lowest)
}
