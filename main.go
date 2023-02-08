package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InitTrie() (*Trie, error) {
	file, err := os.Open("hacking.txt")

	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	trie := NewTrie()

	i := 0
	for fileScanner.Scan() {
		var text = fileScanner.Text()
		trie.Insert(text, i)
		i++
	}
	return trie, file.Close()
}

func Solution1(K int, A []int) int {
	// Implement your solution here
	m := make(map[int][]int)
	res := 0
	for i := 0; i < len(A); i++ {
		x := K - A[i]
		m[x] = append(m[x], i)
		cc, ok := m[A[i]]
		if ok {
			for j := 0; j < len(cc); j++ {
				k := A[i]
				h := A[cc[j]]
				if k == h {
					res = res + 1
					continue
				}
				res = res + 2
			}
		}
	}
	return res
}

func Solution(A []int) int {
	m := make(map[int]int)
	count := 0
	i := 0
	for i < len(A) {
		o := i
		m[i] = i + A[i]
		i = i + A[i]
		v, ok := m[i]
		if ok && v == o {
			count = -1
			break
		}

		count++
	}
	return count
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	var n2, n1 int = 0, 1

	for i := 2; i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	fr := n2 + n1
	if fr/1000000 == 0 {
		return fr
	}
	i := 0
	x := 1
	r := 0
	for i < 6 {
		a := fr % 10
		fr = fr / 10
		r = r + a*x
		x = x * 10
		i++
	}
	return r
}

func main() {
	var f int
	var perms = make(map[string]int)
	perms["W"] = 1
	perms["R"] = 2
	perms["X"] = 4
	perms["read"] = 2
	perms["write"] = 1
	perms["execute"] = 4
	fmt.Scan(&f)
	var states = make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < f; i++ {
		var l string
		if scanner.Scan() {
			l = scanner.Text()
		}

		s := strings.Split(l, " ")
		k := 0
		for j := 1; j < len(s); j++ {
			k |= perms[s[j]]
		}
		states[s[0]] = k
	}
	var a string
	if scanner.Scan() {
		a = scanner.Text()
	}
	t, _ := strconv.Atoi(a)
	for i := 0; i < t; i++ {
		var l string
		if scanner.Scan() {
			l = scanner.Text()
		}
		s := strings.Split(l, " ")
		// fmt.Println(l)
		if states[s[1]]&perms[s[0]] != 0 {
			fmt.Println("OK")
		} else {
			fmt.Println("Access denied")
		}
	}
}
