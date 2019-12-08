package main

import (
	"bufio"
	"fmt"
	"os"
)

var w = 25
var h = 6

func main() {
	page := ""
	lines := []string{}
	freqs := []map[rune]int{}
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		l := in.Text()
		page = l
	}

	for i := 1; i < 101; i++ {
		l := page[:w*h]
		lines = append(lines, l)
		page = page[w*h:]
	}
	for _, chars := range lines {
		m := map[rune]int{}
		for _, c := range chars {
			count := m[c]
			count = count + 1
			m[c] = count
		}
		freqs = append(freqs, m)
	}

	fmt.Println("part 1: frequencies")
	for _, f := range freqs {
		fmt.Printf("%d 1:%d 2:%d  (%d) \n", f['0'], f['1'], f['2'], f['1']*f['2'])
	}
	fmt.Println()

	fmt.Println()
	fmt.Println("part 2: decode")
	fmt.Println("pic")
	for i := 0; i < 150; i++ {
		if ((i) % 25) == 0 {
			fmt.Println()
		}
		fmt.Print(top(i, lines))
	}
	fmt.Println()
}

func top(i int, lines []string) string {
	for _, line := range lines {
		if line[i] != '2' {
			c := line[i]
			if c == '0' {
				return " "
			}
			return "#"
		}
	}
	panic("no top colour")
}
