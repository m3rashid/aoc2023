package main

import (
	"errors"
	"strconv"
	"strings"
)

func parseString2(s string) (l line, err error) {
	r := strings.Split(s, " ")
	if len(r) != 2 {
		return l, errors.New("invalid input")
	}
	l.cl = make([]int, (len(r[0])+1)*5)
	l.cl[0] = 1 // adding padding
	for j := 0; j < 5; j++ {
		for i, v := range r[0] {
			switch v {
			case '.':
				l.cl[i+1+j*(len(r[0])+1)] = 1
			case '#':
				l.cl[i+1+j*(len(r[0])+1)] = 2
			case '?':
				l.cl[i+1+j*(len(r[0])+1)] = 0
			default:
				return l, errors.New("invalid input")
			}
		}
		if j < 4 {
			l.cl[len(r[0])+1+j*(len(r[0])+1)] = 0
		}
	}
	r1 := strings.Split(r[1], ",")
	l.nv = make([]int, len(r1)*5)
	var n int
	for j := 0; j < 5; j++ {
		for i, v := range r1 {
			if n, err = strconv.Atoi(v); err != nil {
				return l, errors.New("invalid input")
			}
			l.nv[i+j*len(r1)] = n
		}
	}
	return l, err
}
