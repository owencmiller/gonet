package linlib

import (
	"fmt"
	"math/rand"
)

// Matrix Structure
type Matrix struct{
	Rows int
	Cols int
	Mat[][] float64
}

// Print multiple Matricies
func printMatrices(matrices... Matrix){
	for i, matrix := range matrices{
		fmt.Println("Matrix -", i)
		fmt.Println(matrix)
	}
}

// Print a Matrix in a visually appealing way
func PrintMatrix(matrix Matrix){
	for i := 0; i < matrix.Rows; i++{
		fmt.Print("[")
		for j := 0; j < matrix.Cols; j++{
			fmt.Printf("%1.3f ", matrix.Mat[i][j])
		}
		fmt.Println("]")
	}
}

// Transpose a matrix
func (mat Matrix) Transpose() Matrix{
	temp := make([][]float64, mat.Cols)
	for i := 0; i < mat.Rows; i++{
		for j := 0; j < mat.Cols; j++{
			temp[j] = append(temp[j], mat.Mat[i][j])
		}
	}
	return Matrix{Rows:mat.Cols, Cols:mat.Rows, Mat: temp}
}

// Average all entries in a matrix
func AverageEntries(matrix Matrix) float64{
	sum := 0.0
	for i := 0; i < matrix.Rows; i++{
		for j := 0; j < matrix.Cols; j++{
			sum += matrix.Mat[i][j]
		}
	}
	return (sum / ((float64) (matrix.Rows * matrix.Cols)))
}

// Method for dotting two matrices
func (Mat1 Matrix) Dot(Mat2 Matrix) Matrix{
	if Mat1.Cols != Mat2.Rows{
		printMatrices(Mat1, Mat2)
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
		printMatrices(Mat1, Mat2)
		panic("Can't apply function to different sized rices")
	}
	for i, _ := range Mat1.Mat{
		for j, _ := range Mat1.Mat[i]{
			Mat1.Mat[i][j] = operation(Mat1.Mat[i][j], Mat2.Mat[i][j])
		}
	}
	return Mat1
}

func Multiply(num1, num2 float64) float64{
	return num1 * num2
}
func Subtract(num1, num2 float64) float64{
	return num1 - num2
}

// Method for applying an operation to all elements in a rix
func ApplyConst(matrix Matrix, operation func(num float64) float64) Matrix{
	temp := GenerateMatrixZeros(matrix.Rows, matrix.Cols)
	for i := range matrix.Mat{
		for j := range matrix.Mat[0]{
			temp.Mat[i][j] = operation(matrix.Mat[i][j])
		}
	}
	return temp
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
			w[i][j] = (rand.Float64()*2)-1
		}
	}
	return Matrix{Rows: Rows, Cols: Cols, Mat: w}
}
