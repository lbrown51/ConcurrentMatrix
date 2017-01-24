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
	in      chan Matrix
	out     chan Matrix
}

// Router takes in the output of some Block computation
// The last column is probably how I'm gonna hold where each
// incoming matrix goes
type Router struct {
	in   chan Matrix
	outs [][]chan Block
}

// Supervisor tells Router what to do and
// Creates new Blocks
type Supervisor struct {
	rtr Router
}

func (sup Supervisor) newBlock() {
	mat := Matrix{make([][]float64, 10)}
	backRtr := make(chan Matrix)
	blk := Block{1, mat, sup.rtr.in}
	fmt.Println(blk)

}

func main() {

	rtr := Router{make(chan Matrix), make([][]chan Block, 10)}
	sup := Supervisor{rtr}
	sup.newBlock()
	fmt.Println(sup)
}
