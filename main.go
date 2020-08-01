package main

import "github.com/timdeklijn/DL/dlm"

func main() {
	matrix := [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
	m := dlm.NewMatrix(matrix)
	m.Print()
}
