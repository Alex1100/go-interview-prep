package trie

type Trie struct {
	children    map[string]*Trie
	end_of_word bool
}

func InitTrie() *Trie {
	return &Trie{
		children:    make(map[string]*Trie),
		end_of_word: false,
	}
}

func (t *Trie) AddString(str string) {
	string_char_count := 0
	current_char := string(str[string_char_count])
	current_trie_node := t

	for string_char_count < len(str) {
		if current_trie_node.children[current_char] == nil {
			current_trie_node.children[current_char] = InitTrie()
		} else {
			current_trie_node.end_of_word = false
			current_trie_node = current_trie_node.children[current_char]
			string_char_count++
			if string_char_count < len(str) {
				current_char = string(str[string_char_count])
			} else {
				break
			}
		}
	}

	current_trie_node.end_of_word = true
}

func (t *Trie) Contains(str string) bool {
	string_char_count := 0
	current_char := string(str[string_char_count])
	current_trie_node := t

	for string_char_count < len(str) {
		_, ok := current_trie_node.children[current_char]

		if ok {
			if string_char_count != len(str)-1 {
				current_trie_node = current_trie_node.children[current_char]
				string_char_count++
				current_char = string(str[string_char_count])
			} else if string_char_count == len(str)-1 {
				return true
			}
		} else {
			return false
		}
	}

	return false
}

func (t *Trie) GetChildren() map[string]*Trie {
	return t.children
}
