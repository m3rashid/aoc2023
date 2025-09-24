package main

import "strconv"

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func next(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	subs, nonZero := getSubs(nums)
	if nonZero {
		return nums[len(nums)-1] + next(subs)
	}
	return nums[len(nums)-1]
}

func getSubs(nums []int) (subs []int, nonZero bool) {
	subs = make([]int, len(nums)-1)
	for i, a := range nums[1:] {
		sub := a - nums[i]
		subs[i] = sub
		if sub != 0 {
			nonZero = true
		}
	}
	return subs, nonZero
}

func prev(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	subs, nonZero := getSubs(nums)
	if nonZero {
		return nums[0] - prev(subs)
	}
	return nums[0]
}

func toInts(ss []string) []int {
	is := make([]int, len(ss))
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		catch(err)
		is[i] = n
	}
	return is
}
