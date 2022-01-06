package main

import "fmt"
 
func main() {
 
   for i := 0; i < 10; i++ {
      for j := 0; j < 10-i; j++ {
         fmt.Print(" ")
      }
      for k := 0; k <= i; k++ {
         fmt.Print("* ")
      }
      fmt.Println()
   }
}