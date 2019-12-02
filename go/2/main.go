package main

import "io/ioutil"
import (
	//	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//replace position 1 with the value 12 and replace position 2 with the
//value 2

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	prog := strings.Split(string(b), ",")
	for i, v := range prog {
		fmt.Println(i, v)
	}

	prog[1] = "12"
	prog[2] = "2"


	code := "0"
	pc := 0
	for {
		fmt.Println(pc)
		code = prog[pc]
		switch code {
		case "1":
			p1 := atoi(prog[pc+1])
			op1 := atoi(prog[p1])
			p2 := atoi(prog[pc+2])
			op2 := atoi(prog[p2])
			p3 := atoi(prog[pc+3])

			fmt.Println("op1", op1)
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
			fmt.Println("HALT!")
			for i, v := range prog {
				fmt.Println(i, v)
			}
			fmt.Println("POS 0", prog[0])
			os.Exit(0)
		default:
			panic("oh dear")
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
