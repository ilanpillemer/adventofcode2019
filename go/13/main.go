package main

import (
	 "os"
	"bufio"
	"log"
	"strconv"
)

type T struct {
  x int
  y int
}


const (
 empty = 0
 wall = 1
 block = 2
 hpaddle = 3
 ball = 4
)

func main () {
   tiles := map[T]int{}
   in := bufio.NewScanner(os.Stdin)
   for in.Scan() {
   x:= atoi( in.Text())
   in.Scan()
   y:= atoi( in.Text())
   in.Scan()
   typ := atoi( in.Text())
   t := T{x,y}
  tiles[t] = typ
//    log.Print(t)
   }
   log.Println("total tiles", len(tiles))
   blocks := 0
   for _, tile := range tiles {
      if tile == block {
       blocks++
     }
   }
  log.Println("total block tiles", blocks)
 }

 func atoi(a string) int {
   i,err := strconv.Atoi(a)
   if err != nil {
     panic(err)
   }
   return i
 }