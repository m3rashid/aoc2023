package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isNumber(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

func toNumber(char string) int {
	if num, err := strconv.Atoi(char); err == nil {
		return num
	}
	return 0
}

func makeMatrix(fileName string) ([][]string, int, int) {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows int
	var cols int
	var rawMatrix = [][]string{}
	for scanner.Scan() {
		str := scanner.Text()
		rawMatrixCol := strings.Split(str, "")
		rawMatrix = append(rawMatrix, rawMatrixCol)
		rows++
		cols = len(rawMatrixCol)
	}

	return rawMatrix, rows, cols
}

func Solution1() {
	rawMatrix, rows, cols := makeMatrix("input1.txt")

	isSymbol := func(i, j int) bool {
		if i < 0 || i >= rows || j < 0 || j >= cols {
			return false
		}

		return rawMatrix[i][j] != "." && !isNumber(rawMatrix[i][j])
	}

	ans := 0
	for i := range rows {
		start := 0
		j := 0

		for j < cols {
			start = j
			num := ""
			for j < cols && isNumber(rawMatrix[i][j]) {
				num += rawMatrix[i][j]
				j++
			}

			if num == "" {
				j++
				continue
			}

			/*
					 XXXX
				-> X20X <-
					 XXXX
			*/
			if isSymbol(i, start-1) || isSymbol(i, j) {
				ans += toNumber(num)
				j++
				continue
			}

			/*
				-> XXXX <-
					 X20X
				-> XXXX <-
			*/
			for k := start - 1; k < j+1; k++ {
				if isSymbol(i-1, k) || isSymbol(i+1, k) {
					ans += toNumber(num)
					break
				}
			}
		}
	}

	fmt.Println(ans)
}
