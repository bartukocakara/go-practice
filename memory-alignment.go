package main

import (
	"fmt"
	"unsafe"
)
type User struct {
	A int32 // size: 4
	B int64 // size: 8
	C int32 // size: 4
	D int64 // size: 8
	E int64 // size: 8
}
func main() {
	x := User{}
	fmt.Printf("size: %d", unsafe.Sizeof(x)) // size: 32
}