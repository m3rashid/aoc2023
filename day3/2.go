package day3

import (
	"fmt"
)

// gear_ratio = multiplication of two numbers equidistant from *

// type StarPos struct {
// 	col     int
// 	n1, n2  int
// 	is_good bool
// }

// func make_sample_matrix() ([][]string, int, int) {
// 	const sample = `467..114..
// 	...*......
// 	..35..633.
// 	......#...
// 	617*......
// 	.....+.58.
// 	..592.....
// 	......755.
// 	...$.*....
// 	.664.598..`

// 	rows := 10
// 	cols := 10

// 	raw_matrix := [][]string{}
// 	for str := range strings.SplitSeq(sample, "\n") {
// 		raw_matrix_col := strings.Split(str, "")
// 		raw_matrix = append(raw_matrix, raw_matrix_col)
// 	}

// 	return raw_matrix, rows, cols
// }

func Solution2() {
	raw_matrix, rows, cols := make_matrix("day3/input1.txt")
	// raw_matrix, rows, cols := make_sample_matrix()

	good_stars := make([][][]int, rows)
	for i := range good_stars {
		good_stars[i] = make([][]int, cols)
	}

	is_symbol := func(i, j, num int) bool {
		if i < 0 || i >= rows || j < 0 || j >= cols {
			return false
		}

		if raw_matrix[i][j] == "*" {
			good_stars[i][j] = append(good_stars[i][j], num)
		}

		return raw_matrix[i][j] != "." && !is_number(raw_matrix[i][j])
	}

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

			num_int := to_number(num)
			_ = is_symbol(i, start-1, num_int) || is_symbol(i, j, num_int)
			for k := start - 1; k < j+1; k++ {
				_ = is_symbol(i-1, k, num_int) || is_symbol(i+1, k, num_int)
			}
		}
	}

	ans := 0
	for i := range rows {
		for j := range cols {
			nums := good_stars[i][j]
			if raw_matrix[i][j] == "*" && len(nums) == 2 {
				ans += (nums[0] * nums[1])
			}
		}
	}

	fmt.Println(ans)
}
