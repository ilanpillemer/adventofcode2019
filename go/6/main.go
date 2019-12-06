package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	tree := map[string]string{}
	rtree := map[string]string{}
	nodes := map[string]bool{}
	for in.Scan() {
		l := in.Text()
		e := strings.Split(l, ")")
		tree[e[0]] = e[1]
		rtree[e[1]] = e[0]
		nodes[e[0]] = true
		nodes[e[1]] = true
	}
	total := 0
	for k := range nodes {
		total += height(k, rtree, 0)
	}
	fmt.Println("orbit count checksum", total)
	you := path("SAN", rtree, []string{"SAN"})
	santa := path("YOU", rtree, []string{"YOU"})
	outerspace := connect(you, santa)
	leg1 := len(santa) - len(outerspace)
	leg2 := len(you) - len(outerspace)

	fmt.Println("journey length: ", (leg1+leg2)-2)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func connect(src []string, dst []string) []string {
	c := []string{}
	lsrc := len(src) - 1
	ldst := len(dst) - 1
	for {
		if src[lsrc] != dst[ldst] {
			return c
		}
		c = append(c, src[lsrc])
		lsrc--
		ldst--
		continue
	}
}

func path(child string, tree map[string]string, p []string) []string {
	parent := tree[child]
	if parent == "COM" {
		p = append(p, "COM")
		return p
	}
	p = append(p, parent)
	return path(parent, tree, p)
}

func height(child string, tree map[string]string, count int) int {
	parent := tree[child]
	if parent == "" {
		return count
	}
	count++
	return height(parent, tree, count)
}
