package main

import (
	"fmt"
	"math"
	mu "github.com/owencmiller/gonet/matrixutil"
)

// Neural Network layers of weights
type Network struct{
	layers[] mu.Matrix
	bias[] mu.Matrix
}

// Apply bias to a weighted sum
func addBias(weights mu.Matrix, bias mu.Matrix) mu.Matrix{
	if weights.Cols != bias.Cols{
		panic("Bias cannot be added")
	}
	for i, row := range weights.Mat{
		temp := make([][]float64, 1)
		for _, val := range row{
			temp[0] = append(temp[0], val)
		}
		weights.Mat[i] = mu.ApplyFunc(mu.CreateMatrix(temp), bias, func(num1 float64, num2 float64)float64{
			return num1+num2
		}).Mat[0]
	}
	return weights
}

// Forward propogate through the Network 
func (net Network) forwardProp(input mu.Matrix) (mu.Matrix, mu.Matrix, mu.Matrix){
	weightedInput, activation := make([]mu.Matrix, len(net.layers))
	for i, layer := range net.layers{
		input = input.Dot(layer)
		weightedInput = append(weightedInput, mu.CopyMatrix(input))
		input = addBias(input, net.bias[i])
		activation = append(activation, mu.CopyMatrix(input))
		input.ApplyConst(relu)
	}
	return input, weightedInput, activation
}

// MSE Loss function
func meanSquaredError(guess mu.Matrix, goal mu.Matrix) mu.Matrix{
	diff := mu.ApplyFunc(guess, goal, func(num1 float64, num2 float64)float64{
		return (math.Pow(num1-num2,2))/2
	})
	return diff
}

// Activation function
func sigmoid(num float64) float64{
	return 1/(1+math.Exp(-num))
}
func divSigmoid(num float64) float64{
	return sigmoid(num)*(1-sigmoid(num))
}
func relu(num float64) float64{
	if num < 0{
		return 0
	}else{
		return num
	}
}
func divRelu(num float64) float64{
	if num < 0{
		return 0
	}else{
		return 1
	}
}


// Create a NN from layerAmounts - including input and output amounts
func createNetwork(layerAmounts ...int) Network{
	layers := make([]mu.Matrix, 0)
	for i, val := range layerAmounts{
		if i+1 == len(layerAmounts){
			break
		}
		layers = append(layers, mu.GenerateMatrixRand(val, layerAmounts[i+1]))
	}

	biases := make([]mu.Matrix, 0)
	for i, _ := range layerAmounts{
		if i+1 == len(layerAmounts){ 
			break 
		}
		biases = append(biases, mu.GenerateMatrixRand(1,layerAmounts[i+1]))
	}
	return Network{layers: layers, bias: biases}
}

// Backpropagation
func (net Network) backProp(output mu.Matrix, goal mu.Matrix){
	for i := len(net.layers)-1; i > 0; i--{
		
	}
}


// Testing and Running
func main() {

	input := [][]float64{
		{1,1},
		{0,0},
	}

	goal := [][]float64{
		{1},
		{0},
	}
	goalMat := mu.CreateMatrix(goal)


	network := createNetwork(2,3,1)
	output := network.forwardProp(mu.CreateMatrix(input))
	loss := meanSquaredError(mu.CopyMatrix(output), goalMat)

	fmt.Println("Network:", network)
	fmt.Println("Output:", output)
	fmt.Println("Loss:", loss)

}
