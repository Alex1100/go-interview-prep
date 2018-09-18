package hash_table

import (
	"errors"
	"fmt"
	singly_linked_list "go-interview-prep/data_structures/linked_lists/singly_linked_list"
	"math"
	"strings"
)

type HashTable struct {
	Storage        []singly_linked_list.HashTableLinkedList
	StorageLimit   int
	Size           int
	ConstantHasher int
}

type hash_key struct {
	num_value int
	str_value string
}

func InitHashTable(initial_size int) *HashTable {
	internal_storage := make([]singly_linked_list.HashTableLinkedList, 0)

	for i := 0; i < initial_size; i++ {
		internal_storage = append(internal_storage, *singly_linked_list.InitEmptyHashTableLinkedList())
	}

	return &HashTable{
		Storage:        internal_storage,
		StorageLimit:   initial_size,
		Size:           0,
		ConstantHasher: 4,
	}
}

func (ht *HashTable) Hash(str_key string, num_key int) int {
	hash := 0
  fmt.Println("STR KEY IS: ", str_key, num_key)
	if len(str_key) > 0 {
		for i := 0; i < len(str_key); i++ {
			hash = (hash << 5) + hash
			hash = hash + int(rune(string(str_key[i])[0])-'0')
			hash = hash & hash
			hash = int(math.Abs(float64(uint(hash) << uint(ht.ConstantHasher))))
		}
	} else if len(str_key) == 0 {
		hash = num_key << 5
		hash = hash & hash
		hash = int(math.Abs(float64(uint(hash) >> uint(ht.ConstantHasher))))
	}

	return hash % ht.StorageLimit
}

func (ht *HashTable) Insert(str_key string, num_key int, data interface{}) bool {
	fmt.Println("HT IS: ", ht.Size)

	if ht.Size == int(math.Floor(float64(ht.StorageLimit)*float64(0.625))) {
		ht.Expand()
	}

	fmt.Println("NOW HT IS: ", ht.Size)

	bucket_index := ht.Hash(str_key, num_key)

	if ht.Storage[bucket_index].Size == 0 {
		return ht.InsertUtil(bucket_index, str_key, num_key, data)
	}

	current := ht.Storage[bucket_index].Head

	for i := 0; i < ht.Storage[bucket_index].Size; i++ {
		temp1 := make([]interface{}, 0)
		temp2 := make([]interface{}, 0)
		temp3 := make([]interface{}, 0)
		temp1 = append(temp1, current.Data)
		temp2 = append(temp2, str_key)
		temp3 = append(temp3, num_key)

		if sliceToString(temp1) == sliceToString(temp2) {
			updated := make([]interface{}, 0)
			updated = append(updated, str_key, data)
			current.Data = updated
			return true
		} else if sliceToString(temp1) == sliceToString(temp3) {
			updated := make([]interface{}, 0)
			updated = append(updated, num_key, data)
			current.Data = updated
			return true
		}

		current = current.Next
	}

	return ht.InsertUtil(bucket_index, str_key, num_key, data)
}

func (ht *HashTable) InsertUtil(bucket_index int, str_key string, num_key int, data interface{}) bool {
	if len(str_key) > 0 {
		item_array := make([]interface{}, 0)
		item_array = append(item_array, str_key, data)
		ht.Storage[bucket_index].Insert(item_array)
	} else {
		item_array := make([]interface{}, 0)
		item_array = append(item_array, num_key, data)
		ht.Storage[bucket_index].Insert(item_array)
	}
	ht.Size++
	return true
}

func (ht *HashTable) Remove(str_key string, num_key int) (*singly_linked_list.Node, error) {
	fmt.Println("YOOO: ", str_key, num_key)
	if ht.Size <= int(math.Floor(float64(ht.StorageLimit)*float64(0.4))) {
		ht.Shrink()
	}

  bucket_index := 0
  if len(str_key) > 0
	  bucket_index = ht.Hash(str_key, 0)
  } else {
    bucket_index = ht.Hash("", num_key)
  }

	if bucket_index > ht.StorageLimit {
		panic(errors.New("Invalid Hash"))
	}

	target := make([]interface{}, 0)

	current_bucket := ht.Storage[bucket_index]
	current := current_bucket.Head

	for current.Data != nil {
		temp1 := make([]interface{}, 0)
		temp2 := make([]interface{}, 0)
		temp3 := make([]interface{}, 0)
		counter := 0
		temp2 = append(temp2, str_key)
		temp3 = append(temp3, num_key)

		val, _ := current.Data.([]interface{})

		for _, value := range val {
			if counter == 0 {
				temp1 = append(temp1, value)
				counter++
			}
		}

		if sliceToString(temp1) == sliceToString(temp2) || sliceToString(temp1) == sliceToString(temp3) {
			target = append(target, current.Data)
			current = nil
		} else {
			current = current.Next
		}
	}

	temp4 := make([]interface{}, 0)
	temp5 := make([]interface{}, 0)
	temp6 := make([]interface{}, 0)
	counter := 0
	temp5 = append(temp5, str_key)
	temp6 = append(temp6, num_key)

	val, _ := current.Data.([]interface{})

	for _, value := range val {
		if counter == 0 {
			temp4 = append(temp4, value)
			counter++
		}
	}
	if sliceToString(temp4) == sliceToString(temp5) || sliceToString(temp4) == sliceToString(temp6) {
		removed := current_bucket.RemoveNode(target)
		ht.Size--
		return removed, nil
	} else {
		panic(errors.New("Does not exist in Hash Table"))
	}
}

func (ht *HashTable) Get(str_key string, num_key int) ([]interface{}, error) {
	bucket_index := ht.Hash(str_key, num_key)
	if bucket_index > ht.StorageLimit {
		panic(errors.New("Invalid Hash"))
	}

	current_bucket := ht.Storage[bucket_index]
	current := current_bucket.Head

	for current.Data != nil {
		temp1 := make([]interface{}, 0)
		temp2 := make([]interface{}, 0)
		temp3 := make([]interface{}, 0)
		counter := 0
		temp2 = append(temp2, str_key)
		temp3 = append(temp3, num_key)

		val, _ := current.Data.([]interface{})

		for _, value := range val {
			if counter == 0 {
				temp1 = append(temp1, value)
				counter++
			}
		}

		if sliceToString(temp1) == sliceToString(temp2) || sliceToString(temp1) == sliceToString(temp3) {
			return temp1, nil
		}

		current = current.Next
	}

	panic(errors.New("Does not exist in Hash Table"))
}

func (ht *HashTable) Contains(str_key string, num_key int) bool {
	iteration_counter := 0
	bucket_index := ht.Hash(str_key, num_key)
	if bucket_index > ht.StorageLimit {
		return false
	}

	current_bucket := ht.Storage[bucket_index]
	current := current_bucket.Head

	for iteration_counter < ht.Storage[bucket_index].Size {
		temp1 := make([]interface{}, 0)
		temp2 := make([]interface{}, 0)
		temp3 := make([]interface{}, 0)
		counter := 0
		temp2 = append(temp2, str_key)
		temp3 = append(temp3, num_key)

		val, _ := current.Data.([]interface{})

		for _, value := range val {
			if counter == 0 {
				temp1 = append(temp1, value)
				counter++
			}
		}

		if sliceToString(temp1) == sliceToString(temp2) || sliceToString(temp1) == sliceToString(temp3) {
			return true
		}

		current = current.Next
		iteration_counter++
	}

	return false
}

func (ht *HashTable) Expand() {
	temp_keys := make([]interface{}, 0)
	temp_vals := make([]interface{}, 0)

	for i := 0; i < len(ht.Storage); i++ {
		if ht.Storage[i].Head.Data != nil {
			for _, key := range ht.Storage[i].ListOfKeys() {
				temp_keys = append(temp_keys, key)
			}

			for _, val := range ht.Storage[i].ListOfValues() {
				temp_vals = append(temp_vals, val)
			}
		}
	}

	internal_storage := make([]singly_linked_list.HashTableLinkedList, 0)
	ht.StorageLimit *= 2

	for j := 0; j < ht.StorageLimit; j++ {
		internal_storage = append(internal_storage, *singly_linked_list.InitEmptyHashTableLinkedList())
	}

	ht.Storage = internal_storage
	ht.Size = 0

	for v := 0; v < len(temp_keys); v++ {
		val, ok := temp_keys[v].(interface{}).(string)
		if ok {
			if len(val) > 0 {
				ht.Insert(val, 0, temp_vals[v])
			}
		} else {
			ht.Insert("", temp_keys[v].(interface{}).(int), temp_vals[v])
		}
	}
}

func (ht *HashTable) Shrink() {
	temp_keys := make([]interface{}, 0)
	temp_vals := make([]interface{}, 0)

	for i := 0; i < len(ht.Storage); i++ {
		if ht.Storage[i].Head.Data != nil {
			for _, key := range ht.Storage[i].ListOfKeys() {
				temp_keys = append(temp_keys, key)
			}

			for _, val := range ht.Storage[i].ListOfValues() {
				temp_keys = append(temp_vals, val)
			}
		}
	}

	internal_storage := make([]singly_linked_list.HashTableLinkedList, 0)
	ht.StorageLimit = int(math.Floor(float64(ht.StorageLimit) / 2.0))

	for j := 0; j < ht.StorageLimit; j++ {
		internal_storage = append(internal_storage, *singly_linked_list.InitEmptyHashTableLinkedList())
	}

	ht.Storage = internal_storage
	ht.Size = 0

	for v := 0; v < len(temp_vals); v++ {
		val, ok := temp_keys[v].(interface{}).(string)
		if ok {
			if len(val) > 0 {
				ht.Insert(val, 0, temp_vals[v])
			}
		} else {
			ht.Insert("", temp_keys[v].(interface{}).(int), temp_vals[v])
		}
	}
}

func sliceToString(values []interface{}) string {
	// Pre-allocate the right size
	s := make([]string, len(values))

	for index := range values {
		s[index] = fmt.Sprintf("%v", values[index])
	}

	fmt.Println("s IS: ", strings.Join(s, ","))

	return strings.Join(s, ",")
}
