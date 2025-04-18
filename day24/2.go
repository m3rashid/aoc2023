package main

import "fmt"

func determinant2x2(mat [2][2]float64) float64 {
	return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
}

func determinant3x3(mat [3][3]float64) float64 {
	return mat[0][0]*(mat[1][1]*mat[2][2]-mat[1][2]*mat[2][1]) -
		mat[0][1]*(mat[1][0]*mat[2][2]-mat[1][2]*mat[2][0]) +
		mat[0][2]*(mat[1][0]*mat[2][1]-mat[1][1]*mat[2][0])
}

func determinant4x4(mat [4][4]float64) float64 {
	return mat[0][0]*determinant3x3([3][3]float64{
		{mat[1][1], mat[1][2], mat[1][3]},
		{mat[2][1], mat[2][2], mat[2][3]},
		{mat[3][1], mat[3][2], mat[3][3]},
	}) -
		mat[0][1]*determinant3x3([3][3]float64{
			{mat[1][0], mat[1][2], mat[1][3]},
			{mat[2][0], mat[2][2], mat[2][3]},
			{mat[3][0], mat[3][2], mat[3][3]},
		}) +
		mat[0][2]*determinant3x3([3][3]float64{
			{mat[1][0], mat[1][1], mat[1][3]},
			{mat[2][0], mat[2][1], mat[2][3]},
			{mat[3][0], mat[3][1], mat[3][3]},
		}) -
		mat[0][3]*determinant3x3([3][3]float64{
			{mat[1][0], mat[1][1], mat[1][2]},
			{mat[2][0], mat[2][1], mat[2][2]},
			{mat[3][0], mat[3][1], mat[3][2]},
		})
}

func result(a, b, c, d, e Hail) (float64, float64, float64) {
	D := determinant4x4([4][4]float64{
		{a.vx - b.vx, b.vy - a.vy, a.py - b.py, b.px - a.px},
		{b.vx - c.vx, c.vy - b.vy, b.py - c.py, c.px - b.px},
		{c.vx - d.vx, d.vy - c.vy, c.py - d.py, d.px - c.px},
		{d.vx - e.vx, e.vy - d.vy, d.py - e.py, e.px - d.px},
	})

	resMat := [4]float64{
		b.px*b.vy - b.py*b.vx + a.py*a.vx - a.px*a.vy,
		c.px*c.vy - c.py*c.vx + b.py*b.vx - b.px*b.vy,
		d.px*d.vy - d.py*d.vx + c.py*c.vx - c.px*c.vy,
		e.px*e.vy - e.py*e.vx + d.py*d.vx - d.px*d.vy,
	}

	DY := determinant4x4([4][4]float64{
		{resMat[0], b.vy - a.vy, a.py - b.py, b.px - a.px},
		{resMat[1], c.vy - b.vy, b.py - c.py, c.px - b.px},
		{resMat[2], d.vy - c.vy, c.py - d.py, d.px - c.px},
		{resMat[3], e.vy - d.vy, d.py - e.py, e.px - d.px},
	})

	DX := determinant4x4([4][4]float64{
		{a.vx - b.vx, resMat[0], a.py - b.py, b.px - a.px},
		{b.vx - c.vx, resMat[1], b.py - c.py, c.px - b.px},
		{c.vx - d.vx, resMat[2], c.py - d.py, d.px - c.px},
		{d.vx - e.vx, resMat[3], d.py - e.py, e.px - d.px},
	})

	DVX := determinant4x4([4][4]float64{
		{a.vx - b.vx, b.vy - a.vy, resMat[0], b.px - a.px},
		{b.vx - c.vx, c.vy - b.vy, resMat[1], c.px - b.px},
		{c.vx - d.vx, d.vy - c.vy, resMat[2], d.px - c.px},
		{d.vx - e.vx, e.vy - d.vy, resMat[3], e.px - d.px},
	})

	// DVY := determinant4x4([4][4]float64{
	// 	{a.vx - b.vx, b.vy - a.vy, a.py - b.py, resMat[0]},
	// 	{b.vx - c.vx, c.vy - b.vy, b.py - c.py, resMat[1]},
	// 	{c.vx - d.vx, d.vy - c.vy, c.py - d.py, resMat[2]},
	// 	{d.vx - e.vx, e.vy - d.vy, d.py - e.py, resMat[3]},
	// })

	X := DX / D
	Y := DY / D

	VX := DVX / D
	// YV := DVY / D

	tempA := (a.px - X) / (VX - a.vx)
	tempB := (b.px - X) / (VX - b.vx)

	detD := determinant2x2([2][2]float64{
		{1, tempA},
		{1, tempB},
	})

	detDZ := determinant2x2([2][2]float64{
		{a.pz + tempA*a.vz, tempA},
		{b.pz + tempB*b.vz, tempB},
	})

	Z := detDZ / detD

	return X, Y, Z
}

func Solution2() {
	allPoints := parseFile("sample.txt")

	X, Y, Z := result(allPoints[0], allPoints[1], allPoints[2], allPoints[3], allPoints[4])
	fmt.Println("Result: ", X+Y+Z)
}
