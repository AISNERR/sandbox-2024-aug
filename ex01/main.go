package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
3
5 1
1
2
3
4
5
1 5
40
2 99
50
1
*/

func main() {
	var in *bufio.Reader
	var out *bufio.Writer


	inputFile, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	in = bufio.NewReader(inputFile)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Fprint(out, `1
5 1
1
2
3
4
5`)

	// var a, b, c int
	// fmt.Fscan(in, &a, &b, &c)
	// fmt.Fprint(out, a+b+c)

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var pr, per int
		fmt.Fscan(in, &pr, &per)

		oshibka := 0
		for i := 0; i < pr; i++ {
			var price float64

			var res float64
			fmt.Fscan(in, &price)
			res = price * float64(per) / 100.0
			resRounded := (int(price) * per / 100) * 100
			kopeiki := int(res * 100)
			oshibka = kopeiki - resRounded
		}
		fmt.Fprintln(out, oshibka)
	}
}
