package dlm

import "testing"

// compareMatrices compares two matrices by first checking if the shapes
// match and afterwards check if all elements are equal.
func compareMatrices(m1, m2 [][]float64) bool {
	if len(m1) != len(m2) || len(m1[0]) != len(m2[0]) {
		return false
	}
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m1[0]); j++ {
			if m1[i][j] != m2[i][j] {
				return false
			}
		}
	}
	return true
}

func TestNewMatrix(t *testing.T) {
	matrix := [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
	m := NewMatrix(matrix)
	if !compareMatrices(m.m, matrix) {
		t.Errorf("Got: %v, expected: %v", m.m, matrix)
	}
}

func TestMatrix_Transpose(t *testing.T) {
	matrix := [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
	expected := [][]float64{{1.0, 4.0}, {2.0, 5.0}, {3.0, 6.0}}
	m := NewMatrix(matrix)
	m.Transpose()
	if !compareMatrices(m.m, expected) {
		t.Errorf("Got: %v, expected: %v", m.m, expected)
	}
}

func TestMatrix_AddMatrix(t *testing.T) {
	m1 := [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
	m := NewMatrix(m1)
	m2 := [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
	n := NewMatrix(m2)
	m.AddMatrix(n)
	expected := [][]float64{{2.0, 4.0, 6.0}, {8.0, 10.0, 12.0}}
	if !compareMatrices(m.m, expected) {
		t.Errorf("Got: %v, expected: %v", m.m, expected)
	}
}

func TestMatrix_AddScalar(t *testing.T) {
	matrix := [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
	expected := [][]float64{{11.0, 12.0, 13.0}, {14.0, 15.0, 16.0}}
	m := NewMatrix(matrix)
	m.AddScalar(10.0)
	if !compareMatrices(m.m, expected) {
		t.Errorf("Got: %v, expected: %v", m.m, expected)
	}
}

func TestMatrix_MultiplyScalar(t *testing.T) {
	matrix := [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
	m := NewMatrix(matrix)
	expected := [][]float64{{3.0, 6.0, 9.0}, {12.0, 15.0, 18.0}}
	m.MultiplyScalar(3.0)
	if !compareMatrices(m.m, expected) {
		t.Errorf("Got: %v, expected: %v", m.m, expected)
	}

}

func TestMatrix_Product(t *testing.T) {
	m1 := [][]float64{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}}
	m := NewMatrix(m1)
	m2 := [][]float64{{7.0, 8.0}, {9.0, 10.0}, {11.0, 12.0}}
	n := NewMatrix(m2)
	expected := [][]float64{{58.0, 64.0}, {139.0, 154.0}}
	m.Product(n)
	if !compareMatrices(m.m, expected) {
		t.Errorf("Got: %v, expected: %v", m.m, expected)
	}
}
