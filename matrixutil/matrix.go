package matrixutil

import (
	"math/rand"
)

// Matrix Structure
type Matrix struct{
	rows int
	cols int
	mat[][] float64
}

// Method for multiplying two matrices
func (mat1 Matrix) Multiply(mat2 Matrix) Matrix{
	if mat1.cols != mat2.rows{
		panic("matrix multiplication not possible")
	}
	ans := make([][]float64, mat1.rows)
	for i, _ := range ans{
		ans[i] = make([]float64, mat2.cols)
		for j, _ := range ans[i]{
			for k, _ := range mat2.mat{
				ans[i][j] += mat1.mat[i][k]*mat2.mat[k][j]
			}
		}
	}
	return Matrix{rows:len(mat1.mat), cols:len(mat2.mat[0]), mat: ans}
}

// Apply a function to two matrices. (ex. subtraction, addition)
func ApplyFunc(mat1 Matrix, mat2 Matrix, operation func(num1 float64,num2 float64) float64) Matrix{
	if mat1.rows != mat2.rows || mat1.cols != mat2.cols{
		panic("Can't apply function to different sized matrices")
	}
	for i, _ := range mat1.mat{
		for j, _ := range mat1.mat[i]{
			mat1.mat[i][j] = operation(mat1.mat[i][j], mat2.mat[i][j])
		}
	}
	return mat1
}

// Method for applying an operation to all elements in a matrix
func (matrix Matrix) ApplyConst(operation func(num float64) float64){
	for i := range matrix.mat{
		for j := range matrix.mat[0]{
			matrix.mat[i][j] = operation(matrix.mat[i][j])
		}
	}
}

// Create a Matrix from 2d array/slice
func CreateMatrix(data [][]float64) Matrix{
	rows := len(data)
	cols := len(data[0])

	return Matrix{rows: rows, cols: cols, mat: data}
}

// Generate a Matrix of zeros
func GenerateMatrixZeros(rows, cols int) Matrix{
	w := make([][]float64, rows)
	for i, _ := range w {
		w[i] = make([]float64, cols)
		for j, _ := range w[i]{
			w[i][j] = 0
		}
	}
	mat := Matrix{rows: rows, cols: cols, mat: w}
	return mat
}

// Generate a matrix of random float64's
func GenerateMatrixRand(rows, cols int) Matrix{
	w := make([][]float64, rows)
	for i, _ := range w {
		w[i] = make([]float64, cols)
		for j, _ := range w[i]{
			w[i][j] = rand.Float64()
		}
	}
	mat := Matrix{rows: rows, cols: cols, mat: w}
	return mat
}
