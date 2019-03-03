package main

import (
	"fmt"
	"math"
	"time"
	mu "github.com/owencmiller/gonet/matrixutil"
)

// Neural Network layers of weights
type Network struct{
	layers[] mu.Matrix
	bias[] mu.Matrix
	lr float64
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
func (net Network) forwardProp(input mu.Matrix) (mu.Matrix, []mu.Matrix, []mu.Matrix){
	weightedInput := make([]mu.Matrix, 0)
	activation := make([]mu.Matrix, 0)
	for i, layer := range net.layers{
		input = input.Dot(layer)
		input = addBias(input, net.bias[i])
		weightedInput = append(weightedInput, mu.CopyMatrix(input))
		input.ApplyConst(relu)
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
		layers = append(layers, mu.GenerateMatrixRand(val, layerAmounts[i+1]))
	}

	biases := make([]mu.Matrix, 0)
	for i, _ := range layerAmounts{
		if i+1 == len(layerAmounts){ 
			break 
		}
		biases = append(biases, mu.GenerateMatrixRand(1,layerAmounts[i+1]))
	}
	return Network{layers: layers, bias: biases, lr: lr}
}


// Backpropagation
func (net Network) backProp(weightedInputs []mu.Matrix, activations []mu.Matrix, inputs mu.Matrix, goal mu.Matrix){
	numOfLayers := len(net.layers)
	
	// Eo = (O-y)dot(sigmoid'(weightedInput))
	weightedInputs[numOfLayers-1].ApplyConst(divRelu)
	err := mu.ApplyFunc(mu.ApplyFunc(activations[numOfLayers-1], goal, func(num1,num2 float64)float64{return num1-num2}), weightedInputs[numOfLayers-1], func(num1,num2 float64)float64{return num1*num2})
	var delta mu.Matrix
	if numOfLayers - 2 < 0{
		delta = inputs.Dot(err)
	}else{
		delta = activations[numOfLayers-2].Dot(err)
	}
	// Apply learning rate
	delta.ApplyConst(func(num float64)float64{return num*net.lr})
	net.layers[numOfLayers-1] = mu.ApplyFunc(net.layers[numOfLayers-1], delta, func(num1,num2 float64)float64{return num1-num2})

	// Calculate errors and deltas for deep layers
	for i := numOfLayers-2; i >= 0; i--{
		weightedInputs[i].ApplyConst(divRelu)
		err = mu.ApplyFunc(mu.ApplyFunc(err, net.layers[i+1], func(num1,num2 float64)float64{return num1*num2}), weightedInputs[i], func(num1,num2 float64)float64{return num1*num2})
		if i == 0{
			delta = inputs.Dot(err)
		}else{
			delta = activations[i-1].Dot(err)
		}
		// Apply learning rate
		delta.ApplyConst(func(num float64)float64{return num*net.lr})
		net.layers[i] = mu.ApplyFunc(net.layers[i], delta, func(num1,num2 float64)float64{return num1-num2})
	}	
}


// Train a Network
func train(network Network, inputs mu.Matrix, goal mu.Matrix){

	for {
		output, weightedInputs, activations := network.forwardProp(inputs)
		err := meanSquaredError(output, goal)
		fmt.Println("Error -", err)
		network.backProp(weightedInputs, activations, inputs, goal)
		time.Sleep(2 * time.Second)
		if err.Mat[0][0] < .0005{
			break
		}
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
	inputMat := mu.CreateMatrix(input)
	goalMat := mu.CreateMatrix(goal)

	learningRate := 0.05

	network := createNetwork(learningRate,2,3,1)
	train(network, inputMat, goalMat)


	test := [][]float64{
		{0,0},
	}
	testGoal := [][]float64{
		{0},
	}
	testMat := mu.CreateMatrix(test)
	testGoalMat := mu.CreateMatrix(testGoal)

	output, _, _ := network.forwardProp(testMat)

	fmt.Println()
	fmt.Println("Guess -", output)
	fmt.Println("Error -", meanSquaredError(output, testGoalMat))
}
