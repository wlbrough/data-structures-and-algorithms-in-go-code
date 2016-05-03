package main

import "fmt"

type MyArray struct {
    size int
    length int
    data []int
}

func NewArray(size int) *MyArray {
    a := new(MyArray)
    a.size = size
    a.length = 0
    a.data = make([]int, a.size)
    return a
}

func main() {
    arr := NewArray(10)
}