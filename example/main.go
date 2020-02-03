package main

import (
	"fmt"

	"github.com/hariadivicky/trie"
)

func main() {
	dictionary := trie.New()
	dictionary.Insert("romane", "romane data example")
	dictionary.Insert("romanus", "romanus data example")
	dictionary.Insert("romalus", "romalus data example")
	dictionary.Insert("rubens", "rubens data example")
	dictionary.Insert("ruber", "ruber data example")
	dictionary.Insert("rubicon", "rubicon data example")
	dictionary.Insert("rubicundus", "rubicundus data example")

	node, err := dictionary.Find("rubicon")

	if err != nil {
		fmt.Printf("could not get result: %v", err)
		return
	}

	fmt.Println("matched result", node.ResultCount)

	if node.IsCompleteWord {
		fmt.Println("payload:", node.Payload)
		// Output
		// matched result 1
		// payload: rubicon data example
	}
}
