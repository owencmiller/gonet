package main

import (
	"fmt"
	"math"
	mu "github.com/owencmiller/gonet/matrixutil"
)

// Neural Network layers of weights
type Network struct{
	layers[] mu.Matrix
}

// Activation function
func sigmoid(num float64) float64{
	return 1/(1+math.Exp(-num))
}

// Forward propogate through the Network TODO make this a method for the Network
func forwardProp(input mu.Matrix, net Network) mu.Matrix{
	for _, layer := range net.layers{
		input = input.Multiply(layer)
		input.ApplyConst(sigmoid)
	}
	return input
}

func meanSquaredError(guess mu.Matrix, goal mu.Matrix) mu.Matrix{
	diff := mu.ApplyFunc(guess, goal, func(num1 float64, num2 float64)float64{
		return (math.Pow(num1-num2,2))/2
	})
	return diff
}

// Create a NN from layerAmounts - including input and output amounts
func createNetwork(layerAmounts ...int) Network{
	temp := make([]mu.Matrix, 0)
	for i, num := range layerAmounts{
		if i+1 == len(layerAmounts){
			break
		}
		temp = append(temp, mu.GenerateMatrixRand(num, layerAmounts[i+1]))
	}
	return Network{layers: temp}
}

// Testing and Running
func main() {
	input := [][]float64{
		{1,2},
		{3,4},
	}
	
	input2 := [][]float64{
		{1,2},
		{3,4},
	}
	
	fmt.Println(mu.ApplyFunc(mu.CreateMatrix(input), mu.CreateMatrix(input2), func(num1 float64, num2 float64)float64{
		return num1+num2
	}))

	network := createNetwork(2,3,1)
	output := forwardProp(mu.CreateMatrix(input), network)
	fmt.Println("Network:", network)
	fmt.Println("Output:", output)

}
