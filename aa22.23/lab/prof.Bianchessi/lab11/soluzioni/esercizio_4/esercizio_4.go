package main

import (
	. "fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	n := os.Args[1]
	d, _ := strconv.Atoi(os.Args[2])

	var nmin uint64 = math.MaxUint64
	for i := uint64(0); i < uint64(math.Pow(2, float64(len(n))))-1; i++ {
		nnew := ""
		count := 0
		bin := toBinary(i, len(n))
		for pos, val := range bin {
			if val == 0 {
				count++
			} else {
				nnew = nnew + string(n[pos])
			}
		}
		if count == d {
			num, _ := strconv.ParseUint(nnew, 10, 64)
			if num < nmin {
				nmin = num
			}
		}
	}
	Println("numero migliore:", nmin)

}

func toBinary(n uint64, l int) []int {
	bin := make([]int, l)
	for i := l - 1; n > 0; i-- {
		bin[i] = int(n % 2)
		n = n / 2
	}
	return bin
}
