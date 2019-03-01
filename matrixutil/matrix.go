package matrixutil

import (
	"math/rand"
)

// Matrix Structure
type Matrix struct{
	Rows int
	Cols int
	Mat[][] float64
}

// Method for multiplying two matrices
func (Mat1 Matrix) Multiply(Mat2 Matrix) Matrix{
	if Mat1.Cols != Mat2.Rows{
		panic("rix multiplication not possible")
	}
	ans := make([][]float64, Mat1.Rows)
	for i, _ := range ans{
		ans[i] = make([]float64, Mat2.Cols)
		for j, _ := range ans[i]{
			for k, _ := range Mat2.Mat{
				ans[i][j] += Mat1.Mat[i][k]*Mat2.Mat[k][j]
			}
		}
	}
	return Matrix{Rows:len(Mat1.Mat), Cols:len(Mat2.Mat[0]), Mat: ans}
}

// Apply a function to two matrices. (ex. subtraction, addition)
func ApplyFunc(Mat1 Matrix, Mat2 Matrix, operation func(num1 float64,num2 float64) float64) Matrix{
	if Mat1.Rows != Mat2.Rows || Mat1.Cols != Mat2.Cols{
		panic("Can't apply function to different sized rices")
	}
	for i, _ := range Mat1.Mat{
		for j, _ := range Mat1.Mat[i]{
			Mat1.Mat[i][j] = operation(Mat1.Mat[i][j], Mat2.Mat[i][j])
		}
	}
	return Mat1
}

// Method for applying an operation to all elements in a rix
func (matrix Matrix) ApplyConst(operation func(num float64) float64){
	for i := range matrix.Mat{
		for j := range matrix.Mat[0]{
			matrix.Mat[i][j] = operation(matrix.Mat[i][j])
		}
	}
}

// Create a matrix from 2d array/slice
func CreateMatrix(data [][]float64) Matrix{
	Rows := len(data)
	Cols := len(data[0])

	return Matrix{Rows: Rows, Cols: Cols, Mat: data}
}

// Copy Matrix
func CopyMatrix(mat Matrix) Matrix{
	temp := make([][]float64, mat.Rows)
	for i, _ := range temp{
		temp[i] = make([]float64, mat.Cols)
		for j, _ := range temp[i]{
			temp[i][j] = mat.Mat[i][j]
		}
	}
	return CreateMatrix(temp)
}

// Generate a rix of zeros
func GenerateMatrixZeros(Rows, Cols int) Matrix{
	w := make([][]float64, Rows)
	for i, _ := range w {
		w[i] = make([]float64, Cols)
		for j, _ := range w[i]{
			w[i][j] = 0
		}
	}
	return Matrix{Rows: Rows, Cols: Cols, Mat: w}
}

// Generate a rix of random float64's
func GenerateMatrixRand(Rows, Cols int) Matrix{
	w := make([][]float64, Rows)
	for i, _ := range w{ 
		w[i] = make([]float64, Cols)
		for j, _ := range w[i]{
			w[i][j] = rand.Float64()
		}
	}
	return Matrix{Rows: Rows, Cols: Cols, Mat: w}
}
