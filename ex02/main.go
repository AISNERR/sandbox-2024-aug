package main

import (
	"bufio"
	"fmt"
	"os"
)

func find_root(inp []int) int {
	var picks = make(map[int]struct{})
	var kids = make(map[int]struct{})
	if len(inp) == 2 {
		return inp[0]
	}
	var i int = 0

	for i < len(inp)-1 {
		picks[inp[i]] = struct{}{}

		i++
		var shift = 0
		if inp[i] > 0 {
			kids_count := inp[i]
			for shift < kids_count {
				i += 1
				kids[inp[i]] = struct{}{}
				shift++
			}
		}
		i++
	}
	for k := range picks {
		if _, ok := kids[k]; !ok {
			return k
		}
	}
	return 0
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var size int
		fmt.Fscan(in, &size)
		var node int
		nodes := make([]int, 0, size)
		for j := 0; j < size; j++ {
			fmt.Fscan(in, &node)
			nodes = append(nodes, node)
		}
		fmt.Println(find_root(nodes))
	}
}
