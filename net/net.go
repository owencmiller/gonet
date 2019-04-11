package main

import (
	"fmt"
	"math"
	mu "github.com/owencmiller/gonet/matrixutil"
)

// Neural Network layers of weights
type Network struct{
	layers[] mu.Matrix
	lr float64
}

// Forward propogate through the Network 
func (net Network) forwardProp(inputMain mu.Matrix) (mu.Matrix, []mu.Matrix, []mu.Matrix){
	weightedInput := make([]mu.Matrix, 0)
	activation := make([]mu.Matrix, 0)
	input := mu.CopyMatrix(inputMain)
	for _, layer := range net.layers{
		input = layer.Dot(input)
		weightedInput = append(weightedInput, mu.CopyMatrix(input))
		input = mu.ApplyConst(input, sigmoid)
		activation = append(activation, mu.CopyMatrix(input))
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
func createNetwork(lr float64, layerAmounts ...int) Network{
	layers := make([]mu.Matrix, 0)
	for i, val := range layerAmounts{
		if i+1 == len(layerAmounts){
			break
		}
		layers = append(layers, mu.GenerateMatrixRand(layerAmounts[i+1], val))
	}

	return Network{layers: layers, lr: lr}
}


// Backpropagation
func (net Network) backProp(weightedInputs []mu.Matrix, activ []mu.Matrix, inputs mu.Matrix, goal mu.Matrix){
	numOfLayers := len(net.layers)
	pos := numOfLayers-1
	errors := make([]mu.Matrix, numOfLayers)
	
	diff := mu.ApplyFunc(activ[pos], goal, mu.Subtract)
	deriv := mu.ApplyConst(weightedInputs[pos], divSigmoid)
	delta := mu.ApplyFunc(diff, deriv, mu.Multiply)
	errors[pos] = delta.Dot(activ[pos-1].Transpose())

	for i := pos; i >= 1; i--{
		diff = net.layers[i].Transpose().Dot(delta)
		deriv = mu.ApplyConst(weightedInputs[i-1], divSigmoid)
		delta = mu.ApplyFunc(diff, deriv, mu.Multiply)
		if i == 1{
			errors[i-1] = delta.Dot(inputs.Transpose())
		}else{
			errors[i-1] = delta.Dot(activ[i-2].Transpose())
		}
	}

	for i := 0; i <= pos; i++{
		net.layers[i] = mu.ApplyFunc(net.layers[i], errors[i], mu.Subtract)
	}
}


// Train a Network
func train(network Network, inputs mu.Matrix, goal mu.Matrix){

	for {
		output, weightedInputs, activations := network.forwardProp(inputs)
		err := meanSquaredError(output, goal)
		network.backProp(weightedInputs, activations, inputs, goal)
		if err.Mat[0][0] < .000005{
			break
		}
	}
}

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
	network := createNetwork(learningRate,4,5,3,2)
	
	train(network, inputMat, goalMat)

	output, _, _ := network.forwardProp(inputMat)

	fmt.Println("Goal - ", goalMat)
	fmt.Println("Output - ", output)
}

func test(){
	input := [][]float64{
		{1,0},
		{1,0},
	}
	inputMat := mu.CreateMatrix(input)
	
	temp := inputMat.Transpose()
	fmt.Println(temp)
}


// Testing and Running
func main() {
	run()
}
