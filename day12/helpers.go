package main

import (
	"bufio"
	"fmt"
	"os"
)

type line struct {
	cl []int // combination line
	nv []int // numeric values
}

func itFits(pattern []int, pos int, cl []int) bool {
	if len(pattern)+pos > len(cl) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if pattern[i] != cl[pos+i] && cl[pos+i] != 0 {
			return false
		}
	}
	return true
}

var cache map[string]int

func convertToKey(i int, a []int, j int) (key string) {
	out := make([]rune, len(a))
	s := fmt.Sprintf("%d,%d,", i, j)
	for idx, v := range a {
		out[idx] = rune(v)
	}
	return s + string(out)
}

func getFitting(combinationList []int, startPos int, numericValueList []int, index int) (result int) {
	// get current pattern
	pattern := make([]int, numericValueList[index]+1)
	pattern[0] = 1
	for k := 0; k < numericValueList[index]; k++ {
		pattern[k+1] = 2
	}
	// find next valid match
	for j := startPos; j < len(combinationList); j++ {
		if itFits(pattern, j, combinationList) {
			if index == len(numericValueList)-1 {
				valid := true
				for h := j + len(pattern); h < len(combinationList); h++ {
					if combinationList[h] == 2 {
						valid = false
						break
					}
				}
				if valid {
					result += 1
				}
			} else {
				if value, ok := cache[convertToKey(j, pattern, index+1)]; ok {
					result += value
				} else {
					val := getFitting(combinationList, j+len(pattern), numericValueList, index+1)
					cache[convertToKey(j, pattern, index+1)] = val
					result += val
				}
			}
		}
		for h := startPos; h < j+1; h++ {
			if combinationList[h] == 2 {
				return result
			}
		}
	}
	return result
}

func getCombinations(s string, parseString func(s string) (l line, err error)) (cs int, err error) {
	var l line
	if l, err = parseString(s); err != nil {
		return cs, err
	}
	cache = make(map[string]int)
	return getFitting(l.cl, 0, l.nv, 0), err
}

func getSumCombinations(s []string, parseString func(s string) (l line, err error)) (sum int, err error) {
	var v int
	for _, c := range s {
		if v, err = getCombinations(c, parseString); err != nil {
			return sum, err
		}
		sum += v
	}
	return sum, err
}

func ReadInput(inputFile string) (content []string, err error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return content, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	err = scanner.Err()

	return content, err
}
