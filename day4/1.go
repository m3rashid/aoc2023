package day4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func handle_cardline(str string) float64 {
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

	num_match := 0
	for _, num := range mine_nums {
		if slices.Contains(wining_nums, num) {
			num_match++
		}
	}

	if num_match > 0 {
		return math.Pow(2, float64(num_match)-1)
	}

	return 0
}

func Solution1() {
	file, _ := os.Open("day4/input1.txt")
	defer file.Close()

	var ans float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ans += handle_cardline(scanner.Text())
	}

	fmt.Println("Answer:", ans)
}
