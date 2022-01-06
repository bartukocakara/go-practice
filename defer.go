package main

import "fmt"

// Functions
func mul(a1, a2 int) int {
 
    res := a1 * a2
    fmt.Println("Result: ", res)
    return 0
}
 
func show() {
    fmt.Println("Hello!, GeeksforGeeks")
}
 
// Main function
func main() {
 
    // Alttaki methodlar çalışmadan defer devreye girmez
    defer mul(23, 2)

    mul(23, 45)

    show()
}