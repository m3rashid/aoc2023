package main

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Pos struct {
	x, y int
}

func getExpandedUniverse(lines []string, expand int) []Pos {
	expY := make([]int, len(lines))
	expX := make([]int, len(lines[0]))
	for x := range expX {
		expX[x] = expand - 1
	}
	for y := range expY {
		expY[y] = expand - 1
	}

	var galaxies []Pos
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				galaxies = append(galaxies, Pos{x, y})
				expX[x] = 0
				expY[y] = 0
			}
		}
	}
	for x := range expX[1:] {
		expX[x+1] += expX[x]
	}
	for y := range expY[1:] {
		expY[y+1] += expY[y]
	}
	for i, g := range galaxies {
		galaxies[i] = Pos{g.x + expX[g.x], g.y + expY[g.y]}
	}
	return galaxies
}
