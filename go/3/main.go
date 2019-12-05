package main

import "io/ioutil"
import (
	//	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//part 1 is the answer for noun = 12 and verb = 2

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	orig := strings.Split(string(b), ",")

	//	for i := 0; i < 100; i++ {
	//		for j := 0; j < 100; j++ {
	//			prog := make([]string, len(orig))
	//			copy(prog, orig)
	//			noun := strconv.Itoa(i)
	//			verb := strconv.Itoa(j)
	//			prog[1] = noun
	//			prog[2] = verb
	//
	//			code := "0"
	//			pc := 0
	//		Loop:

	prog := make([]string, len(orig))
	copy(prog, orig)
	pc := 0
	//part1
	//	input := "1"
	//part2
	input := "5"
Loop:
	for {
		precode := prog[pc]
		code, mode1, mode2, _ := modes(precode)
		switch code {
		case "1", "01": //addition
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog)
			p3 := atoi(prog[pc+3])

			prog[p3] = strconv.Itoa(op1 + op2)
			pc = pc + 4
		case "2", "02": //multiplication
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog)
			p3 := atoi(prog[pc+3])

			prog[p3] = strconv.Itoa(op1 * op2)
			pc = pc + 4
		// instructions with one parameter
		case "3": //save
			saveTo := atoi(prog[pc+1])

			prog[saveTo] = input
			pc = pc + 2
		case "4", "04": //output
			p1 := atoi(prog[pc+1])
			output := getvalue(mode1, p1, prog)
			fmt.Println("output => ", output)
			pc = pc + 2
		// jumps
		case "5", "05": //jump if true
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog)

			if op1 != 0 {
				pc = op2
				continue
			}
			pc = pc + 3
		case "6", "06": //jump if false
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog)

			if op1 == 0 {
				pc = op2
				continue
			}
			pc = pc + 3
		//comparisons
		case "7", "07": //less than
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog)
			p3 := atoi(prog[pc+3])
			if op1 < op2 {
				prog[p3] = strconv.Itoa(1)
			} else {
				prog[p3] = strconv.Itoa(0)
			}
			pc = pc + 4
		case "8", "08": //lequals
			p1 := atoi(prog[pc+1])
			op1 := getvalue(mode1, p1, prog)
			p2 := atoi(prog[pc+2])
			op2 := getvalue(mode2, p2, prog)
			p3 := atoi(prog[pc+3])
			if op1 == op2 {
				prog[p3] = strconv.Itoa(1)
			} else {
				prog[p3] = strconv.Itoa(0)
			}
			pc = pc + 4
		case "99":
			fmt.Println("HALT!")
			fmt.Println("POS 0", prog[0])
			break Loop
		default:
			fmt.Println("code", code, mode1)

			panic("oh dear")
		}
	}
	//		} //j
	//	}//i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}
func atoi(s string) int {
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("illegal", s, err.Error())
		panic("oh dear")
	}
	return i
}

func getvalue(mode int, i int, prog []string) int {
	if mode == 0 {
		return atoi(prog[i])
	}
	return i
}

// tidy this up before it gets even more complex!!
func modes(s string) (opcode string, mode1 int, mode2 int, mode3 int) {
	s = strings.TrimSpace(s)
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
		fmt.Println(s)
		opcode = s[3:]
		mode1 = atoi(s[2:3])
		mode2 = atoi(s[1:2])
		mode3 = atoi(s[0:1])
		return
	}
	panic("ohdear")
}
