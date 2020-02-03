# Trie

## What is trie?

<img align="right" width="169px" src="https://upload.wikimedia.org/wikipedia/commons/b/be/Trie_example.svg">

In computer science, a trie, also called digital tree or prefix tree, is a kind of search tree—an ordered tree data structure used to store a dynamic set or associative array where the keys are usually strings. Unlike a binary search tree, no node in the tree stores the key associated with that node; instead, its position in the tree defines the key with which it is associated. All the descendants of a node have a common prefix of the string associated with that node, and the root is associated with the empty string. Keys tend to be associated with leaves, though some inner nodes may correspond to keys of interest. Hence, keys are not necessarily associated with every node. [Read more on Wikipedia](https://en.wikipedia.org/wiki/Trie)

## Installation
To install trie package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.12+ is required**), then you can use the below Go command to install trie.

```sh
$ go get -u github.com/hariadivicky/trie
```

2. Import it in your code:

```go
import "github.com/hariadivicky/trie"
```


## Usage Example
```go
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

```