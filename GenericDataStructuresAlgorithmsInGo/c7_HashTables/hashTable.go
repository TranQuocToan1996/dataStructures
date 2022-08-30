package hash

import (
	"hash/fnv" // Fowler–Noll–Vo algorithm
	// Other details not shown yet
)

const tableSize = 100_000

func hash(s string) uint32 {
	h := fnv.New32a() // Fowler-Noll-Vo algorithm
	h.Write([]byte(s))
	return h.Sum32()
}

var length int

type WordType struct {
	word string
	list []string
}

// At every index there is a word and slice of words
type HashTable [tableSize]WordType

func NewTable() HashTable {
	var table HashTable
	for i := 0; i < tableSize; i++ {
		table[i] = WordType{"", []string{}}
	}
	return table
}

func (table *HashTable) Insert(word string) {
	index := hash(word) % tableSize // Between 0 and tableSize - 1
	// Search table[index] for word
	if table[index].word == word {
		return // duplicates not allowed
	}
	if len(table[index].list) > 0 {
		for i := 0; i < len(table[index].list); i++ {
			if table[index].list[i] == word {
				return // duplicates not allowed
			}
		}
	}
	if table[index].word == "" {
		table[index].word = word
	} else {
		table[index].list = append(table[index].list, word)
	}
	length += 1
}

func (table HashTable) IsPresent(word string) bool {
	index := hash(word) % tableSize // Between 0 and tableSize - 1
	// Search table[index] for word
	if table[index].word == word {
		return true
	}
	if len(table[index].list) > 0 {
		for i := 0; i < len(table[index].list); i++ {
			if table[index].list[i] == word {
				return true
			}
		}
	}
	return false
}
