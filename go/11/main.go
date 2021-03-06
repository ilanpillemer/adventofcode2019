package main

import "io/ioutil"
import (
	//	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type P struct {
	x int
	y int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

//part 1 is the answer for noun = 12 and verb = 2
const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	orig := strings.Split(string(b), ",")

	prog := make([]string, math.MaxInt32)

	copy(prog, orig)
	pc := 0
	base := 0
	input := ""
	fmt.Println("starting")
	hull := map[P]int{}
	current := P{0, 0}
	facing := up
	paintmode := true
	totalPainted := 0
	hull[current] = 1

	input = itoa(hull[current])
Loop:
	for {
		precode := prog[pc]
		code, mode1, mode2, mode3 := modes(precode)
		switch code {
		case "1", "01": //addition
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog, base)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog, base)
			p3 := atoi(prog[pc+3])
			if mode3 == 2 {
				p3 = base + p3
			}
			prog[p3] = strconv.Itoa(op1 + op2)
			pc = pc + 4
		case "2", "02": //multiplication
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog, base)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog, base)
			p3 := atoi(prog[pc+3])
			if mode3 == 2 {
				p3 = base + p3
			}
			prog[p3] = strconv.Itoa(op1 * op2)
			pc = pc + 4
		// instructions with one parameter
		case "3", "03": //save
			p1 := atoi(prog[pc+1])
			//p1 is zero
			//base is 1000
			if mode1 == 2 {
				p1 = base + p1
			}
			prog[p1] = input
			pc = pc + 2
		case "4", "04": //output
			p1 := atoi(prog[pc+1])
			output := getvalue(mode1, p1, prog, base)
			if paintmode {
				//fmt.Println("(paint) output => ", output)
				pc = pc + 2
				paintmode = !paintmode
				hull[current] = output
				totalPainted++
				//fmt.Println("painted ", current)
				//display(hull, current)
				continue
			}
			//fmt.Println("(turn) output => ", output)
			pc = pc + 2
			switch output {
			case 0: // turn left
				switch facing {
				case up:
					current.x = current.x - 1
					facing = left
				case down:
					current.x = current.x + 1
					facing = right
				case left:
					current.y = current.y + 1
					facing = down
				case right:
					current.y = current.y - 1
					facing = up
				}
			case 1: // turn right
				switch facing {
				case up:
					current.x = current.x + 1
					facing = right
				case down:
					current.x = current.x - 1
					facing = left
				case left:
					current.y = current.y - 1
					facing = up
				case right:
					current.y = current.y + 1
					facing = down
				}

			}
			input = itoa(hull[current])
			paintmode = !paintmode
		// jumps
		case "5", "05": //jump if true
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog, base)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog, base)

			if op1 != 0 {
				pc = op2
				continue
			}
			pc = pc + 3
		case "6", "06": //jump if false
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog, base)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog, base)

			if op1 == 0 {
				pc = op2
				continue
			}
			pc = pc + 3
		//comparisons
		case "7", "07": //less than
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog, base)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog, base)
			p3 := atoi(prog[pc+3])
			if mode3 == 2 {
				p3 = base + p3
			}

			if op1 < op2 {
				prog[p3] = strconv.Itoa(1)
			} else {
				prog[p3] = strconv.Itoa(0)
			}
			pc = pc + 4
		case "8", "08": //lequals
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog, base)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog, base)
			p3 := atoi(prog[pc+3])
			if mode3 == 2 {
				p3 = base + p3
			}

			if op1 == op2 {
				prog[p3] = strconv.Itoa(1)
			} else {
				prog[p3] = strconv.Itoa(0)
			}
			pc = pc + 4
		case "9", "09": //change relative base
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog, base)
			base = base + op1

			pc = pc + 2
		case "99":
			fmt.Println()
			fmt.Println("HALT!")
			fmt.Println("POS 0", prog[0])
			fmt.Println("panels painted: ", len(hull))
			fmt.Println("panels painted: ", totalPainted)
			display(hull, current)
			break Loop
		default:
			fmt.Println("code", code, mode1)

			panic("oh dear")
		}
	}

}

func display(grid map[P]int, current P) {
	grid[current] = 8
	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			switch grid[P{x, y}] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Printf("#")
			}

		}
		fmt.Println()
	}
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
func atoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("illegal", s, err.Error())
		panic("oh dear")
	}
	return i
}

func getvalue(mode int, i int, prog []string, base int) int {
	if mode == 0 {
		return atoi(prog[i])
	}

	if mode == 2 {
		return atoi(prog[base+i])
	}
	return i
}

// tidy this up before it gets even more complex!!
func modes(s string) (opcode string, mode1 int, mode2 int, mode3 int) {
	s = strings.TrimSpace(s)
	//fmt.Println(s)
	if len(s) == 1 {
		opcode = s
		mode1 = 0
		mode2 = 0
		mode3 = 0
		return
	}

	if len(s) == 2 {
		opcode = s[0:]
		mode1 = 0
		mode2 = 0
		mode3 = 0
		return
	}

	if len(s) == 3 {
		opcode = s[1:]
		mode1 = atoi(s[0:1])
		mode2 = 0
		mode3 = 0
		return
	}

	if len(s) == 4 {
		opcode = s[2:]
		mode1 = atoi(s[1:2])
		mode2 = atoi(s[0:1])
		mode3 = 0
		return
	}

	if len(s) == 5 {
		opcode = s[3:]
		mode1 = atoi(s[2:3])
		mode2 = atoi(s[1:2])
		mode3 = atoi(s[0:1])
		return
	}
	panic("ohdear")
}
