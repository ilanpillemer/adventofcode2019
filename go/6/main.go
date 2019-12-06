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
	//	leafs := map[string]bool{}
	nodes := map[string]bool{}
	for in.Scan() {
		l := in.Text()
		e := strings.Split(l, ")")
		tree[e[0]] = e[1]
		rtree[e[1]] = e[0]
		nodes[e[0]] = true
		nodes[e[1]] = true
	}
	for k, v := range rtree {
		//		leafs[v] = true
		fmt.Println(k, "=>", v)
	}

	//	fmt.Println("D", height("D", rtree, 0))
	total := 0
	for k := range nodes {
		total += height(k, rtree, 0)
//		fmt.Println(k, height(k, rtree, 0))
	}
	fmt.Println("total", total)

}

func height(child string, tree map[string]string, count int) int {
	parent := tree[child]
	//	fmt.Println("p", parent)
	if parent == "" {
		return count
	}
	count++
	return height(parent, tree, count)
}
