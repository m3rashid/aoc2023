package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_number(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

func to_number(char string) int {
	if num, err := strconv.Atoi(char); err == nil {
		return num
	}
	return 0
}

func make_matrix(fileName string) ([][]string, int, int) {
	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows int
	var cols int
	var raw_matrix = [][]string{}
	for scanner.Scan() {
		str := scanner.Text()
		raw_matrix_col := strings.Split(str, "")
		raw_matrix = append(raw_matrix, raw_matrix_col)
		rows++
		cols = len(raw_matrix_col)
	}

	return raw_matrix, rows, cols
}

func Solution1() {
	raw_matrix, rows, cols := make_matrix("day3/input1.txt")

	is_symbol := func(i, j int) bool {
		if i < 0 || i >= rows || j < 0 || j >= cols {
			return false
		}

		return raw_matrix[i][j] != "." && !is_number(raw_matrix[i][j])
	}

	ans := 0
	for i := range rows {
		start := 0
		j := 0

		for j < cols {
			start = j
			num := ""
			for j < cols && is_number(raw_matrix[i][j]) {
				num += raw_matrix[i][j]
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
			if is_symbol(i, start-1) || is_symbol(i, j) {
				ans += to_number(num)
				j++
				continue
			}

			/*
				-> XXXX <-
					 X20X
				-> XXXX <-
			*/
			for k := start - 1; k < j+1; k++ {
				if is_symbol(i-1, k) || is_symbol(i+1, k) {
					ans += to_number(num)
					break
				}
			}
		}
	}

	fmt.Println(ans)
}
