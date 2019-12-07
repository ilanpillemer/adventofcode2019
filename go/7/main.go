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
var orig = []string{}

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	orig = strings.Split(string(b), ",")

	//43210
	max := 0
	for i := 0; i < 100000; i++ {
		//	for i := 43210; i < 43211; i++ {
		str := itoa(i)
		str = prepend(str)
		if hasdups(str) {
			continue
		}

		o1 := exec("0", str[:1])
		o2 := exec(o1, str[1:2])
		o3 := exec(o2, str[2:3])
		o4 := exec(o3, str[3:4])
		o5 := exec(o4, str[4:5])

		//		fmt.Println(str[:1])
		//		fmt.Println(str[1:2])
		//		fmt.Println(str[2:3])
		//		fmt.Println(str[3:4])
		//		fmt.Println(str[4:5])
		if atoi(o5) > max {
			max = atoi(o5)
		}
	//	fmt.Println(o5, "<====", str)
	}
	fmt.Println("max thruster", max)

}

func hasdups(str string) bool {
	dup := map[string]bool{}
	dup[str[:1]] = true
	dup[str[1:2]] = true
	dup[str[2:3]] = true
	dup[str[3:4]] = true
	dup[str[4:5]] = true
	return len(dup) < 5
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func prepend(s string) string {
	l := abs(len(s) - 5)
	if l > 0 {
		s = strings.Repeat("0", l) + s
	}
	return s
}

func exec(input string, signal string) string {
	signalled := false
	prog := make([]string, len(orig))
	copy(prog, orig)
	pc := 0
	//part1
	//	input := "1"
	//part2
	//	input := "5"
	outsignal := ""
	//Loop:
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
			data := input
			if !signalled {
				data = signal
				signalled = true
			}
			saveTo := atoi(prog[pc+1])
			prog[saveTo] = data
			pc = pc + 2
		case "4", "04": //output
			p1 := atoi(prog[pc+1])
			output := getvalue(mode1, p1, prog)
			//fmt.Println("output => ", output)
			outsignal = itoa(output)
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
			//	fmt.Println("HALT!")
			//fmt.Println("POS 0", prog[0])
			//fmt.Println("halt output => ", outsignal)

			return outsignal
			//break Loop
		default:
			fmt.Println("code", code, mode1)

			panic("oh dear")
		}
	}
	//		} //j
	//	}//i
	//	return ""
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
