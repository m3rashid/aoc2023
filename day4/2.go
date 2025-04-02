package day4

import (
	"os"
	"strings"
)

func handle_cardline_part2(str string) float64 {
	nums_str := strings.Trim(strings.Split(str, ":")[1], " ")
	req_nums := strings.Split(nums_str, "|")

	mine_nums := strings.Split(strings.Trim(req_nums[1], " "), " ")
	wining_nums := strings.Split(strings.Trim(req_nums[0], " "), " ")

	new_mine_nums := []string{}
	for i := range mine_nums {
		mine_nums[i] = strings.Trim(mine_nums[i], " ")
		if mine_nums[i] != "" {
			new_mine_nums = append(new_mine_nums, mine_nums[i])
		}
	}
	mine_nums = new_mine_nums

	new_winning_nums := []string{}
	for i := range wining_nums {
		wining_nums[i] = strings.Trim(wining_nums[i], " ")
		if wining_nums[i] != "" {
			new_winning_nums = append(new_winning_nums, wining_nums[i])
		}
	}
	wining_nums = new_winning_nums

	return 0
}

func Solution2() {
	file, _ := os.Open("day4/input1.txt")
	defer file.Close()
}
