package main

import (
	"bufio"
	"fmt"
	"os"
)

func longestSubarrayWithTwoDistinct(nums []int) int {
	left := 0
	count := make(map[int]int)
	maxLength := 0

	for right := 0; right < len(nums); right++ {
		count[nums[right]]++

		for len(count) > 2 {
			count[nums[left]]--
			if count[nums[left]] == 0 {
				delete(count, nums[left])
			}
			left++
		}

		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
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
		fmt.Println(longestSubarrayWithTwoDistinct(nodes))
	}
}
