/**
This is the main package for Gonet; a modular neural network framework written in pure Golang. 
**/
package main

import (
	"fmt"
	"time"
	mu "github.com/owencmiller/gonet/linlib"
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

	learningRate := 1.0
	network := net.CreateNetwork(learningRate,4,5,3,2)

    start := time.Now()
	network.Train(inputMat, goalMat)
    end := time.Since(start)


	output, _, _ := network.ForwardProp(inputMat)

	fmt.Println("Goal - ")
	mu.PrintMatrix(goalMat)
	fmt.Println("Output - ")
	mu.PrintMatrix(output)
	fmt.Println("Error - ")
	mu.PrintMatrix(net.MeanSquaredError(output, goalMat))
    fmt.Printf("Time Elapsed - %s\n", end)
}

func main() {
	run()
}
