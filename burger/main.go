package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var s = bufio.NewScanner(os.Stdin)

func main() {
	s.Split(bufio.ScanWords)
	t := nextInt() // test case 数

	for i := 1; i <= t; i++ {
		k := nextInt()       // burger の中の要素数
		ds := make([]int, k) // バンズまでの距離のslice
		bs := make([]int, k) // 最適なバンズまでの距離のslice
		for j := 0; j < k; j++ {
			ds[j] = distance(k, j) // バンズまでの距離
			bs[j] = nextInt()      // 最適なバンズまでの距離
		}

		// 最小誤差を出すために並び替える
		sort.Sort(sort.IntSlice(ds))
		sort.Sort(sort.IntSlice(bs))

		es := 0 // 誤差合計
		for j := 0; j < k; j++ {
			d := square(ds[j] - bs[j]) // 誤差
			log.Print(d)
			es += d
		}
		fmt.Printf("Case #%d: %d\n", i, es)
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

// 次乗
func square(x int) int {
	return x * x
}

// バンズからの距離
func distance(k, j int) int {
	d := k - 1
	if (d - j) < j {
		return d - j
	}
	return j
}
