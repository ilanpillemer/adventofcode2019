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

	//	fmt.Println(len(page))
//	n := len(page) / w * h
	//	fmt.Println(w * h)
	//	fmt.Println(n)

	for i := 1; i < 101; i++ {
		l := page[:w*h]
		//		fmt.Println(len(l), len(page))
		//		fmt.Println(i, l)

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

	for _, f := range freqs {
		fmt.Printf("%d 1:%d 2:%d  (%d) \n", f['0'], f['1'], f['2'], f['1']*f['2'])
	}

	//	for i, layer := range page {
	//		fmt.Println("line: ", i, layer)
	//	}

}
