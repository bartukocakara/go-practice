package main

import "fmt"

type DemoStruct struct {
	Val int
}

//A.
func demo_func() DemoStruct {
	return DemoStruct{Val: 1}
}

// //B.
// func demo_func() *DemoStruct {
//     return &DemoStruct{}
// }
// //C.
// func demo_func(s *DemoStruct) {
//     s.Val = 1
// }

func main() {
	fmt.Println(demo_func())
}