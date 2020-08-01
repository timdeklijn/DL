package dlm

import "fmt"

// Matrix
type Matrix struct {
	m [][]float64
}

func NewMatrix(m [][]float64) Matrix {
	return Matrix{m}
}

func (m *Matrix) Transpose() {
	var tmpMat [][]float64
	for i := 0; i < len(m.m[0]); i++ {
		var tmpRow []float64
		for j := 0; j < len(m.m); j++ {
			tmpRow = append(tmpRow, m.m[j][i])
		}
		tmpMat = append(tmpMat, tmpRow)
	}
	m.m = tmpMat
}

func (m *Matrix) Print() {
	for _, row := range m.m {
		fmt.Println(row)
	}
}

func (m *Matrix) AddMatrix(n Matrix) {
	var tmpMat [][]float64
	for i := 0; i < len(m.m); i++ {
		var row []float64
		for j := 0; j < len(m.m[0]); j++ {
			row = append(row, m.m[i][j]+n.m[i][j])
		}
		tmpMat = append(tmpMat, row)
	}
	m.m = tmpMat
}

func (m *Matrix) AddScalar(s float64) {
	for i := 0; i < len(m.m); i++ {
		for j := 0; j < len(m.m[0]); j++ {
			m.m[i][j] = m.m[i][j] + s
		}
	}
}

func (m *Matrix) MultiplyScalar(s float64) {
	for i := 0; i < len(m.m); i++ {
		for j := 0; j < len(m.m[0]); j++ {
			m.m[i][j] = m.m[i][j] * s
		}
	}
}

func (m *Matrix) Product(n Matrix) {
	var tmpMat [][]float64
	for i := 0; i < len(m.m); i++ {
		var row []float64
		for j := 0; j < len(n.m[0]); j++ {
			s := 0.0
			for ii := 0; ii < len(m.m[0]); ii++ {
				s += m.m[i][ii] * n.m[ii][j]
			}
			row = append(row, s)
		}
		tmpMat = append(tmpMat, row)
	}
	m.m = tmpMat
}
