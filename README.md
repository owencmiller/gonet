# gonet
Deep Modular Neural Network in Go

## Setup
cd into your /go directory and run the following commands
```
$ go get github.com/owencmiller/gonet.git
$ cd gonet/net
$ go build
$ ./net
```
Creation of the network is as easy as 
```
network = createNetwork(lr float64, layerAmounts ...int)
``` 
## Implementation
This project is a fully modular deep neural network in Golang. 
It contains a linear algrebra abstraction which allows for easy modification along with easy NN creation.

## Why?
This project came out of a want to learn Golang and to better understand backpropagation. I have had experience with neural networks in the past, however I have never attempted to create one with knowledge of lienar algebra. Having heard of Golang and being interested in the language, I decided to create this project in my spare time. Having implemented the matrixutil aspect of the project, I had learned the basics of Go syntax and had setup the steps for neural network creation. The actual neural network implementation was straightforward except for the backpropagation. This by far took the longest time to understand and to implement. Nonetheless, once completed it works like a charm.

## Next Steps
The future implementations of this project plan to have many more types of neural networks as well as higher modularity. Currently, the goal is to create a similar library to tensorflow, where one can create a neural network by specifying the layers and activation types.
Another update may seperate the matrixutil into a different linear algebra library.

## Design
This project uses the equations found on https://sudeepraja.github.io/Neural/

### Contributors
Owen Miller
