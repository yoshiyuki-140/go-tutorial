package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	rows    = 10
	columns = 30
)

type Grid [][]bool

func newGrid() Grid {
	grid := make(Grid, rows)
	for i := range grid {
		grid[i] = make([]bool, columns)
	}
	return grid
}

func (g Grid) print() {
	for _, row := range g {
		for _, cell := range row {
			if cell {
				fmt.Print("■")
			} else {
				fmt.Print("□")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) countNeighbors(row, col int) int {
	count := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < rows && j >= 0 && j < columns && !(i == row && j == col) {
				if g[i][j] {
					count++
				}
			}
		}
	}
	return count
}

func (g Grid) update() Grid {
	newGrid := newGrid()
	for i := range g {
		for j := range g[i] {
			neighbors := g.countNeighbors(i, j)
			if g[i][j] {
				newGrid[i][j] = neighbors == 2 || neighbors == 3
			} else {
				newGrid[i][j] = neighbors == 3
			}
		}
	}
	return newGrid
}

func main() {
	// rand.Seed(time.Now().UnixNano())

	grid := newGrid()

	// ランダムに初期状態を生成
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2) == 1
		}
	}

	for i := 0; i < 20; i++ {
		fmt.Printf("Generation %d\n", i+1)
		grid.print()
		time.Sleep(500 * time.Millisecond)
		grid = grid.update()
	}
}
