package main

import "fmt"

func main() {
	var i, s = findResult(5, "BAACC")
	fmt.Printf("%d , %s", i, s)

	fmt.Printf("expect : 3 , Bruno")
}

func findResult(n int, s string) (int, string) {
	if n > 100 {
		return 0, ""
	}

	var pA, pB, pG, result string
	var a, b, g, max int

	for len(pA) < n {
		pA += "ABC"
	}
	for len(pB) < n {
		pB += "BABC"
	}
	for len(pG) < n {
		pG += "CCAABB"
	}

	for i := 0; i < n; i++ {
		if s[i] == pA[i] {
			a += 1
		}
		if s[i] == pB[i] {
			b += 1
		}
		if s[i] == pG[i] {
			g += 1
		}
	}

	if a > max {
		max = a
	}
	if b > max {
		max = b
	}
	if g > max {
		max = g
	}

	if a == max {
		result += "Adrian"
	}
	if b == max {
		result += "Bruno"
	}
	if g == max {
		result += "Goran"
	}

	return max, result
}
