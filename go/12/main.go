package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type P struct {
	x  int
	y  int
	z  int
	vx int
	vy int
	vz int
}

func main() {
	in := bufio.NewScanner(os.Stdin)

	state := []P{}
	for in.Scan() {
		l := in.Text()
		var xyz = regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)`)
		m := xyz.FindStringSubmatch(l)
		state = append(state, P{atoi(m[1]), atoi(m[2]), atoi(m[3]), 0, 0, 0})
	}
	fmt.Println(state)
	for i := 0; i < 1000; i++ {
		state = next(state)
	}
	fmt.Println(state)
	fmt.Println(total(state))
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func pot(p P) int {
	return abs(p.x) + abs(p.y) + abs(p.z)
}

func kin(p P) int {
	return abs(p.vx) + abs(p.vy) + abs(p.vz)
}

func energy(p P) int {
	return pot(p) * kin(p)
}

func total(s []P) int {
	t := 0
	for _, p := range s {
		t = t + energy(p)
	}
	return t
}

func next(s []P) []P {
	n := []P{}
	n = append(n, s...)

	var d1x, d1y, d1z, d2x, d2y, d2z int
	d1x, d1y, d1z, d2x, d2y, d2z = gravity(s[0], s[1])
	applyGravity(d1x, d1y, d1z, n, 0)
	applyGravity(d2x, d2y, d2z, n, 1)

	d1x, d1y, d1z, d2x, d2y, d2z = gravity(s[0], s[2])
	applyGravity(d1x, d1y, d1z, n, 0)
	applyGravity(d2x, d2y, d2z, n, 2)

	d1x, d1y, d1z, d2x, d2y, d2z = gravity(s[0], s[3])
	applyGravity(d1x, d1y, d1z, n, 0)
	applyGravity(d2x, d2y, d2z, n, 3)

	d1x, d1y, d1z, d2x, d2y, d2z = gravity(s[1], s[2])
	applyGravity(d1x, d1y, d1z, n, 1)
	applyGravity(d2x, d2y, d2z, n, 2)

	d1x, d1y, d1z, d2x, d2y, d2z = gravity(s[1], s[3])
	applyGravity(d1x, d1y, d1z, n, 1)
	applyGravity(d2x, d2y, d2z, n, 3)

	d1x, d1y, d1z, d2x, d2y, d2z = gravity(s[2], s[3])
	applyGravity(d1x, d1y, d1z, n, 2)
	applyGravity(d2x, d2y, d2z, n, 3)
	n = velocity(n)
	return n
}

func applyGravity(dx, dy, dz int, n []P, i int) {
	ns := n[i]
	ns.vx = ns.vx + dx
	ns.vy = ns.vy + dy
	ns.vz = ns.vz + dz
	n[i] = ns
}

func gravity(p1 P, p2 P) (d1x, d1y, d1z, d2x, d2y, d2z int) {
	switch {
	case p1.x < p2.x:
		d1x++
		d2x--
	case p1.x > p2.x:
		d1x--
		d2x++
	}
	switch {
	case p1.y < p2.y:
		d1y++
		d2y--
	case p1.y > p2.y:
		d1y--
		d2y++
	}
	switch {
	case p1.z < p2.z:
		d1z++
		d2z--
	case p1.z > p2.z:
		d1z--
		d2z++
	}
	return
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func velocity(n []P) []P {
	for i, p := range n {
		p.x = p.x + p.vx
		p.y = p.y + p.vy
		p.z = p.z + p.vz
		n[i] = p
	}
	return n
}
