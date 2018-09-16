package main

import (
	"fmt"
	trie "go-interview-prep/data_structures/tries/trie"
)

func main() {
	tt := *trie.InitTrie()
	tt.AddString("Yo")
	fmt.Println(*tt.GetChildren()["Y"])
	fmt.Println(tt.Contains("Yo"))
	fmt.Println(tt.Contains("Hello"))
}
