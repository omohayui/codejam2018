package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var s = bufio.NewScanner(os.Stdin)

func main() {
	s.Split(bufio.ScanWords)
	t := nextInt()

	for i := 1; i <= t; i++ {
		n := nextInt()
		m := nextInt()
		fmt.Printf("Case #%d: %d %d\n", i, n+m, n*m)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func nextInt() int {
	s.Scan()
	i, e := strconv.Atoi(s.Text())
	if e != nil {
		panic(e)
	}
	return i
}
