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
	fmt.Println("\n\nSTORAGE IS: ", ht.Storage)
	fmt.Println("\n\n\n", ht.Storage[6].Head)
	ht.Insert("", 99, "Charlie Alpha Bravo")
	ht.Insert("", 100, 100)
	ht.Insert("", 33, 1234)
	fmt.Println("AYOOO: ", ht.Hash("Crossing Threshold", 0))

	ht.Insert("Crossing Threshold", 0, "Woah")
	fmt.Println("YOOO: ", ht.Hash("Should have crossed Threshold", 0))
	ht.Insert("Should have crossed Threshold", 0, "Expanding")
	fmt.Println("\n\nCHECK EXPAND FUNC:: ", ht.Storage)
	fmt.Println("\n\nCHECK SIZE:: ", ht.Size)
	fmt.Println("\n\nCHECK CONTAINS Hello:: ", ht.Contains("Hello", 0))
	fmt.Println("\n\nCHECK CONTAINS DUCK:: ", ht.Contains("DUCK", 0))
	fmt.Println("\n\nREMOVE NEEDS WORK.. CHECK REMOVE:::", ht.Remove("Hello", 0))
}
