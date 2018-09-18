package main

import (
	"fmt"
	hash_table "go-interview-prep/data_structures/hash_tables/hash_table"
)

func main() {
	ht := *hash_table.InitHashTable(10)

	fmt.Println(ht.Hash("", 123))
	fmt.Println(ht.Hash("Hello", 0))
	ht.Insert("Hello", 0, "Alex")
	ht.Insert("Yoo", 0, "Jon")
	fmt.Println("\n\nHT STILL NEEDS WORK. IT IS: ", ht.Storage)
}
