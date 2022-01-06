package main

import "fmt"
 
func main() {
 
   height := 5
   width := 20
 
   for h := 1; h <= height; h++ {
      for w := 1; w <= width; w++ {
         if h == 1 || h == height {
            fmt.Print("*")
         } else if w == 1 || w == width {
            fmt.Print("*")
         } else {
            fmt.Print(" ")
         }
      }
      fmt.Println()
   }
}