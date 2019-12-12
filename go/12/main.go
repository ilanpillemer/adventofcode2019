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

func itoa(i int) string {
	return strconv.Itoa(i)

}

func hashx(s []P) string {
	str := ""
	for _, c := range s {
		str = str + itoa(c.x) + "," + itoa(c.vx) + ","
	}
	return str
}
func hashy(s []P) string {
	str := ""
	for _, c := range s {
		str = str + itoa(c.y) + "," + itoa(c.vy) + ","
	}
	return str
}
func hashz(s []P) string {
	str := ""
	for _, c := range s {
		str = str + itoa(c.z) + "," + itoa(c.vz) + ","
	}
	return str
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
	startx := hashx(state)
	starty := hashy(state)
	startz := hashz(state)

	var m1, m2, m3 int
	count := 0
	tick := 0
	for count < 3 {
		tick++
		state = next(state)
		nextx := hashx(state)
		nexty := hashy(state)
		nextz := hashz(state)
		if nextx == startx && m1 == 0 {
			m1 = tick
			count++
		}
		if nexty == starty && m2 == 0 {
			m2 = tick
			count++
		}
		if nextz == startz && m3 == 0 {
			m3 = tick
			count++
		}

	}

	fmt.Println(state)
	fmt.Println(m1, m2, m3)
	fmt.Println(LCM(m1, m2, m3))
	//	fmt.Println(total(state)) part 1 (calculated after a 1000 ticks)
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

//from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
