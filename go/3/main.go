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
	//	prog[1] = "12"
	//	prog[2] = "2"

	//	code := "0"
	pc := 0
	input := "1"
Loop:
	for {
		//		fmt.Println("code", code)
		precode := prog[pc]
		code, mode1, mode2, _ := modes(precode)
		//		fmt.Println("len", len(prog))
		//		fmt.Println("prog", prog)
		//fmt.Println("code", code, mode1)
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
			output := atoi(prog[p1])
			fmt.Println("output => ", output)
			pc = pc + 2
		case "99":
			//	if prog[0] == "19690720" {
			fmt.Println("HALT!")
			//				fmt.Printf("noun %v verb %v \n", noun, verb)
			fmt.Println("POS 0", prog[0])
			//	}
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
	i, err := strconv.Atoi(s)
	if err != nil {
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

func modes(s string) (opcode string, mode1 int, mode2 int, mode3 int) {
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
