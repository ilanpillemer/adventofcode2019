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

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			prog := make([]string, len(orig))
			copy(prog, orig)
			noun := strconv.Itoa(i)
			verb := strconv.Itoa(j)
			prog[1] = noun
			prog[2] = verb

			code := "0"
			pc := 0
		Loop:
			for {
				code = prog[pc]
				switch code {
				case "1":
					p1 := atoi(prog[pc+1])
					op1 := atoi(prog[p1])
					p2 := atoi(prog[pc+2])
					op2 := atoi(prog[p2])
					p3 := atoi(prog[pc+3])

					prog[p3] = strconv.Itoa(op1 + op2)
					pc = pc + 4
				case "2":
					p1 := atoi(prog[pc+1])
					op1 := atoi(prog[p1])
					p2 := atoi(prog[pc+2])
					op2 := atoi(prog[p2])
					p3 := atoi(prog[pc+3])

					prog[p3] = strconv.Itoa(op1 * op2)
					pc = pc + 4
				case "99":
					if prog[0] == "19690720" {
						fmt.Println("HALT!")
						fmt.Printf("noun %v verb %v \n", noun, verb)
						fmt.Println("POS 0", prog[0])
					}
					break Loop
				default:
					panic("oh dear")
				}
			}
		}
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("oh dear")
	}
	return i
}
