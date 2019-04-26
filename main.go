package main

import (
	"fmt"
	//"time"
	mu "github.com/owencmiller/gonet/matrixutil"
    net "github.com/owencmiller/gonet/net"
)

func run(){
	input := [][]float64{
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{1, 0, 1, 0},
		{1, 0, 1, 0},
	}
	goal := [][]float64{
		{0, 1, 1, 0},
		{0, 1, 0, 1},
	}
	inputMat := mu.CreateMatrix(input)
	goalMat := mu.CreateMatrix(goal)

	learningRate := 0.15
	network := net.CreateNetwork(learningRate,4,5,3,2)

	net.Train(network, inputMat, goalMat)

	output, _, _ := network.ForwardProp(inputMat)

	fmt.Println("Goal - ")
	mu.PrintMatrix(goalMat)
	fmt.Println("Output - ")
	mu.PrintMatrix(output)
	fmt.Println("Error - ")
	mu.PrintMatrix(net.MeanSquaredError(output, goalMat))
}

func main() {
	run()
}
