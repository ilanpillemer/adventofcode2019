package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type P struct {
	x int
	y int
}

var maxx int
var maxy int

func main() {
	grid := map[P]bool{}
	in := bufio.NewScanner(os.Stdin)
	y := 0
	for in.Scan() {
		line := in.Text()
		maxx = (len(line))
		for x, c := range line {
			if c == '#' {
				grid[P{x, y}] = true
			}
		}
		y++
	}
	maxy = y
	display(grid)
	displayCounts(grid)
	maxVisible(grid)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func visible(p P, grid map[P]bool) int {
	//the number of slopes that contain asteroids are the number of visible asteroids
	slopes := map[float64]bool{}
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			// in the equation
			// slope = dx/dy
			if grid[P{x, y}] && (P{x, y} != p) { // the point contains an asteroid
				dx := p.x - x
				dy := p.y - y
				slope := math.Atan2(float64(dy), float64(dx))
				slopes[slope] = true
			}
		}
	}
	return len(slopes)
}

func display(grid map[P]bool) {
	fmt.Println(maxx, maxy)
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			if grid[P{x, y}] {
				fmt.Print("#")
				continue
			}
			print(".")
		}
		fmt.Println()
	}
}

func displayCounts(grid map[P]bool) {
	fmt.Println(maxx, maxy)
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			if grid[P{x, y}] {
				fmt.Print(visible(P{x, y}, grid))
				continue
			}
			print(".")
		}
		fmt.Println()
	}
}

func maxVisible(grid map[P]bool) {
	fmt.Println(maxx, maxy)
	max := 0
	pmax := P{}
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			if grid[P{x, y}] {
				if max < visible(P{x, y}, grid) {
					max = visible(P{x, y}, grid)
					pmax.x = x
					pmax.y = y
				}
				continue
			}
		}
	}
	fmt.Printf("%d Max Visible at %v \n", max, pmax)
}
