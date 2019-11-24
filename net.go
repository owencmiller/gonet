package gonet

import (
	"fmt"
	"math"
	//"time"
	mu "github.com/owencmiller/gonet/linlib"
)

// Neural Network layers of weights
type Network struct {
	layers []mu.Matrix
	lr     float64
}

// Activation function
func sigmoid(num float64) float64 {
	return 1 / (1 + math.Exp(-num))
}
func divSigmoid(num float64) float64 {
	return sigmoid(num) * (1 - sigmoid(num))
}
func relu(num float64) float64 {
	if num < 0 {
		return 0
	}
	return num
}
func divRelu(num float64) float64 {
	if num < 0 {
		return 0
	}
	return 1
}

// MSE Loss function
func MeanSquaredError(guess mu.Matrix, goal mu.Matrix) mu.Matrix {
	diff := mu.ApplyFunc(guess, goal, func(num1 float64, num2 float64) float64 {
		return (math.Pow(num1-num2, 2)) / 2
	})
	return diff
}

// Create a NN from layerAmounts - including input and output amounts
func CreateNetwork(lr float64, layerAmounts ...int) Network {
	layers := make([]mu.Matrix, 0)
	for i, val := range layerAmounts {
		if i+1 == len(layerAmounts) {
			break
		}
		layers = append(layers, mu.GenerateMatrixRand(layerAmounts[i+1], val))
	}

	return Network{layers: layers, lr: lr}
}

// Forward propagate through the Network
func (net Network) ForwardProp(inputMain mu.Matrix) (mu.Matrix, []mu.Matrix, []mu.Matrix) {
	weightedInput := make([]mu.Matrix, 0)
	activation := make([]mu.Matrix, 0)
	input := mu.CopyMatrix(inputMain.Transpose())
	for _, layer := range net.layers {
		input = layer.Dot(input)
		weightedInput = append(weightedInput, mu.CopyMatrix(input))
		input = mu.ApplyConst(input, sigmoid)
		activation = append(activation, mu.CopyMatrix(input))
	}
	return input.Transpose(), weightedInput, activation
}

// Backpropagation
func (net Network) backProp(weightedInputs []mu.Matrix, activ []mu.Matrix, inputs mu.Matrix, goalMain mu.Matrix) {
	goal := goalMain.Transpose()
	numOfLayers := len(net.layers)
	pos := numOfLayers - 1
	errors := make([]mu.Matrix, numOfLayers)

	diff := mu.ApplyFunc(activ[pos], goal, mu.Subtract)
	deriv := mu.ApplyConst(weightedInputs[pos], divSigmoid)
	delta := mu.ApplyFunc(diff, deriv, mu.Multiply)
	errors[pos] = delta.Dot(activ[pos-1].Transpose())

	for i := pos; i >= 1; i-- {
		diff = net.layers[i].Transpose().Dot(delta)
		deriv = mu.ApplyConst(weightedInputs[i-1], divSigmoid)
		delta = mu.ApplyFunc(diff, deriv, mu.Multiply)
		if i == 1 {
			errors[i-1] = delta.Dot(inputs)
		} else {
			errors[i-1] = delta.Dot(activ[i-2].Transpose())
		}
	}

	for i := 0; i <= pos; i++ {
		errors[i] = mu.MultiplyConst(errors[i], net.lr)
		net.layers[i] = mu.ApplyFunc(net.layers[i], errors[i], mu.Subtract)
	}
}

// Train a Network
func (net Network) Train(inputs mu.Matrix, goal mu.Matrix, iterations int) {
	counter := 0
	for {
		output, weightedInputs, activations := net.ForwardProp(inputs)
		errors := MeanSquaredError(output, goal)
		net.backProp(weightedInputs, activations, inputs, goal)
		if counter%1000 == 0 {
			fmt.Println("Error - ", mu.AverageEntries(errors))
		}
		if mu.AverageEntries(errors) < .0000005 || counter == iterations{
			break
		}
		counter++
	}
}
