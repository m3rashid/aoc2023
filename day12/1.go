package main

import (
	"errors"
	"strconv"
	"strings"
)

func parseString(s string) (l line, err error) {
	r := strings.Split(s, " ")
	if len(r) != 2 {
		return l, errors.New("invalid input")
	}
	l.cl = make([]int, len(r[0])+1)
	l.cl[0] = 1 // adding padding
	for i, v := range r[0] {
		switch v {
		case '.':
			l.cl[i+1] = 1
		case '#':
			l.cl[i+1] = 2
		case '?':
			l.cl[i+1] = 0
		default:
			return l, errors.New("invalid input")
		}
	}
	r1 := strings.Split(r[1], ",")
	l.nv = make([]int, len(r1))
	var n int
	for i, v := range r1 {
		if n, err = strconv.Atoi(v); err != nil {
			return l, errors.New("invalid input")
		}
		l.nv[i] = n
	}
	return l, err
}
