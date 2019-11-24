package main

import(
	"encoding/binary"
	"fmt"
	mat "github.com/owencmiller/gonet/linlib"
	"path/filepath"

	//net "github.com/owencmiller/gonet"
	"os"
)

/**
*  Spec: http://yann.lecun.com/exdb/mnist/
**/

type Mnist_data struct{
	num_images uint32
	num_rows, num_cols uint32
	features []mat.Matrix
	labels [] uint32
}


func check(e error){
	if e != nil{
		panic(e)
	}
}

func read_headers(mnist *Mnist_data, dat *os.File){
	b4 := make([]byte, 4)
	n1, err := dat.Read(b4)
	magic_number := binary.BigEndian.Uint32(b4)
	check(err)
	fmt.Printf("%d bytes: %d \n", n1, magic_number)

	n2, err := dat.Read(b4)
	num_images := binary.BigEndian.Uint32(b4)
	check(err)
	fmt.Printf("%d bytes: %d \n", n2, num_images)

	if mnist.num_images == 0{
		mnist.num_images = num_images
	}else{
		if mnist.num_images != num_images{
			panic("Unequal number of features and labels")
		}
	}
}

func read_features(mnist *Mnist_data, dat *os.File){
	read_headers(mnist, dat)

	b4 := make([]byte, 4)
	n3, err := dat.Read(b4)
	num_rows := binary.BigEndian.Uint32(b4)
	check(err)
	fmt.Printf("%d bytes: %d \n", n3, num_rows)

	n4, err := dat.Read(b4)
	num_cols := binary.BigEndian.Uint32(b4)
	check(err)
	fmt.Printf("%d bytes: %d \n", n4, num_cols)

	mnist.num_rows = num_rows
	mnist.num_cols = num_cols
	read_images(mnist, dat)
}

func read_images(mnist *Mnist_data, dat *os.File){
	for i := 0; i < int(mnist.num_images); i++{
		read_image(mnist, dat)
		fmt.Printf("Image: %d\n", i)
	}
}

func read_image(mnist *Mnist_data, dat *os.File){
	image := make([][]float64, mnist.num_rows)
	for i := 0; i < int(mnist.num_rows); i++{
		image[i] = make([]float64, mnist.num_cols)
		for j := 0; j < int(mnist.num_cols); j++{
			b1 := make([]byte, 1)
			_, err := dat.Read(b1)
			check(err)
			image[i][j] = float64(b1[0])
		}
	}
	mnist.features = append(mnist.features, mat.CreateMatrix(image))
}


func read_mnist(mnist *Mnist_data){
	absPath, _ := filepath.Abs("./mnist_data/train-images-idx3-ubyte")
	dat, err := os.Open(absPath)
	check(err)
	read_features(mnist, dat)
}


func mnist(){
	mnist := &Mnist_data{features: []mat.Matrix{}, labels: []uint32{}}
	read_mnist(mnist)
}

func main(){
	mnist()
}

