package main

import (
	"fmt"
	"sort"
)

func main() {
	IsPalindrom("Anna")
}

func IsPalindrom(str string) bool {
	for i, v := range str {
		fmt.Printf("%d %s", i, string(v))
	}
	return false
}

func TwoSort(arr []string) string {
	sort.Strings(arr)
	first := arr[0]
	var r []rune

	for _, v := range first {
		r = append(r, v)
		r = append(r, '*')
		r = append(r, '*')
		r = append(r, '*')
	}
	return fmt.Sprintf("%s", string(r[:len(r)-3]))
}
