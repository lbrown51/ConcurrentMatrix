package main

import "fmt"

// Matrix Is a stand-in for a matrix type
type Matrix struct {
	array [][]float64
}

// Block is a concurrent holder for the
// Matrix type or a piece of a Matrix type
type Block struct {
	id      int
	weights Matrix
	//in      chan Matrix
	//out     chan Matrix
}

// Router takes in the output of some Block computation
// The last column is probably how I'm gonna hold where each
// incoming matrix goes
type Router struct {
	//in   chan Matrix
	//outs map[int][10]chan Block
}

// Supervisor tells Router what to do and
// Creates new Blocks
type Supervisor struct {
}

func main() {
	mat := Matrix{[][]float64{{2.0, 2.0}, {1.0, 1.6}}}
	blk := Block{1, mat}
	fmt.Println(blk)
}
