package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	t := targets(P{11, 13}, grid)
	fmt.Println("#keys ", len(t))
	keys := make([]int, 0, len(t))
	for k := range t {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	//sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	i := 1
	start := false
	//145 for test 2
	up := 15707963267 // pi / 2
	for _, k := range keys {
		//	fmt.Println(i, k, t[k])
		if k == up {
			start = true
		}
		if i == 200 {
			fmt.Println("200th asteroid destroyed:", t[k])
		}
		if start {
			i++
		}

	}
	for _, k := range keys {
		//	fmt.Println(i, k, t[k])
	if i == 200 {
			fmt.Println("200th asteroid destroyed:", t[k])
		}
		if start {
			i++
		}

	}

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
				//fmt.Println(int(slope * 1E10))
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
		fmt.Printf("%2d \n", y)
	}
}

func displayCounts(grid map[P]bool) {
	fmt.Println(maxx, maxy)
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			if grid[P{x, y}] {
				fmt.Printf("%4d", visible(P{x, y}, grid))
				continue
			}
			print(" .. ")
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

func targets(p P, grid map[P]bool) map[int][]int {
	//the number of slopes that contain asteroids are the number of visible asteroids
	targets := map[int][]int{}
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			// in the equation
			// slope = dx/dy
			if grid[P{x, y}] && (P{x, y} != p) { // the point contains an asteroid
				dx := p.x - x
				dy := p.y - y
				slope := math.Atan2(float64(dy), float64(dx))
				key := (int(slope * 1e10))
				e, ok := targets[key]
				if !ok {
					e = []int{}
				}
				e = append(e, (x*100)+y)
				targets[key] = e
			}
		}
	}
	return targets
}
