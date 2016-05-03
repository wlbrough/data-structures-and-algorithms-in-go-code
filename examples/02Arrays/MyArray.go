package main

import "fmt"

// Define MyArray struct
// MyArray is an int array because array type cannot be dynamically assigned
// This example uses a slice for data storage rather than a true array, because
// array length cannot be set using a variable
type MyArray struct {
    size int
    length int
    data []int
}

// Constructor function for MyArray struct
// Used so each field does not have to be manually defined
func NewArray(size int) *MyArray {
    // initialize new variable `a` with `MyArray` type
    a := new(MyArray)
    // Update instance properties
    a.size = size
    a.length = 0
    a.data = make([]int, a.size)
    // return a pointer to `MyArray` struct instance
    return &a
}

func main() {
    // Create a new `MyArray` instance with length 10
    arr := NewArray(10)
}