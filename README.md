# Go Tool: Streamlining Tasks Efficiently in Golang

## Package Array

### Overview

The Array Package provides a thread-safe array data structure in Go that can store various data types as specified by the `allow` interface.

### Usage

#### Installation

To install the package, you can use the following Go command:


go get github.com/k-vanio/go-tool

```sh
package main

import (
	"fmt"
	"github.com/k-vanio/go-tool/array"
)

func main() {
	// Create a new array of integers with a capacity of 10
	arr := array.New[int](10)

	// Add some elements to the array
	arr.Push(5)
	arr.Push(10)

	// Get the elements from the array
	elem1, _ := arr.At(0)
	elem2, _ := arr.At(1)

	fmt.Println("Element 1:", elem1)
	fmt.Println("Element 2:", elem2)
}
```

