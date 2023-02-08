package main

import "fmt"

// tinh tong lon nhat day con lien tiep
func b1(arr []int) int {
	b := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if b < sum {
				b = sum
			}
		}
	}
	return b
}

// tinh tong 2 so a vab a < b
func b2_naive(a int, b int) int {
	sum := 0
	for i := 0; i < b; i++ {
		if i < a {
			sum += 2
		} else {
			sum += 1
		}
	}
	return sum
}

func main() {
	a := []int{1, -2, 3, 4, -1}
	fmt.Println(b1(a))
	fmt.Println(b2_naive(4, 11))
}
