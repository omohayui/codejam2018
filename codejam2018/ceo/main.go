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
	t := nextInt() // test case 数

	for i := 1; i <= t; i++ {
		l := nextInt() // experience level の数
		ceo := 0       // ceo 最低レベル
		tmpNum := 0    // manger を必要とするtmp人数
		for j := 0; j < l; j++ {
			num := nextInt()
			level := nextInt()
			if tmpNum-level*num <= 0 {
				tmpNum = 0
			} else {
				tmpNum -= level * num
			}
			tmpNum += num
			if ceo <= level {
				ceo = level + 1
			}
		}

		if ceo < tmpNum {
			ceo = tmpNum
		}

		fmt.Printf("Case #%d: %d\n", i, ceo)
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
