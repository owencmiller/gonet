# gonet
Deep Modular Neural Network in Go

## Setup
Go into your /go directory and run the following commands
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

## Design
This project uses the equations found on https://sudeepraja.github.io/Neural/

### Contributors
Owen Miller
