package main

import "fmt"

var p []int
var used []bool

// Sample code to generate all permutations of k elements from a set of size n.
func main() {
	n := 5
	k := 3
	p = make([]int, k)
	used = make([]bool, n)
	for i := 0; i < n; i++ {
		used[i] = true
		p[0] = i
		gen(1, n, k)
		used[i] = false
	}

}

func gen(m, n, k int) {
	if m >= k {
		fmt.Println(p)
	} else {
		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				p[m] = i
				gen(m+1, n, k)
				used[i] = false
			}
		}
	}
}
