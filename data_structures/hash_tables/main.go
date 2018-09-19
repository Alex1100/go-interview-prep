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
	fmt.Println(ht.Storage[8].Head.Data)
	fmt.Println("DERRRP: ", ht.Storage[8].Size)

	removed1, err := ht.Remove("Hello", 0)
	if err == nil {
		fmt.Println("\n\nREMOVE NEEDS WORK.. CHECK REMOVE:::", removed1)
	}
	fmt.Println("DERRRP: ", ht.Storage[8].Head, ht.Storage[8].Tail)
	ht.Remove("Should have crossed Threshold", 0)
	fmt.Println("LANDING")
	ht.Remove("Yoo", 0)
	// fmt.Println(ht)
	ht.Remove("", 99)
	ht.Remove("", 100)
	ht.Remove("", 33)

	ht.Remove("Crossing Threshold", 0)
	fmt.Println(ht.Storage)
	ht.Insert("A", 0, "Alex")
	fmt.Println("\n\n\n\n")
	fmt.Println("\nHASH TABLE IS NOW: ", ht)
}
