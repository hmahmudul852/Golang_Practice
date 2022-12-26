package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (a byLength) Len() int           { return len(a) }
func (a byLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLength) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
