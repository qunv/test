package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {

	/*
		Bài toán này sử dụng cấu trúc Trie để index data, phục vụ mục đích search,
		Các lines có chứa word X sẽ được lưu trữ vào node đã đuược đánh dấu là end of word khi insert vào Trie
		Ưu điểm, chi phí tìm kiếm nhỏ, nhược điểm chi phí bộ nhớ lớn trong trường hợp file có các từ trùng nhau là ít

		Bài toàn này còn có thể tối ưu hơn bằng cách dùng cấu trúc Radix Tree để init tất cả các word trong file
	*/
	trie, err := InitTrie()
	if err != nil {
		panic(err)
	}
	lines := trie.Search("xinh")

	fmt.Println("Num of line contains word: ", lines)
}
