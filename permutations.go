package main

import "fmt"

var p []int
var used []bool

//Sample algorothim to generate all permutations of n size set
func main() {
	n := 5
	p = make([]int, n)
	used = make([]bool, n)
	for i := 0; i < n; i++ {
		used[i] = true
		p[0] = i
		gen(1, n)
		used[i] = false
	}
}

func gen(m, n int) {
	if m >= n {
		fmt.Println(p)
	} else {
		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				p[m] = i
				gen(m+1, n)
				used[i] = false
			}
		}
	}
}
