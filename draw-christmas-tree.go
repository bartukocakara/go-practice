package main

import "fmt"
 
func main() {
 
   height := 10;
 
   for y := height ; y > 0; y-- {
      for x :=1 ; x < (height * 2);  x++{
         radius := height - y
         if x == height {
            fmt.Print("|")
         } else if (x >= height - radius ) && (x < height) {
            fmt.Print("\\")
         } else if (x > height) && (x <= height + radius){
            fmt.Print("/")
         } else {
            fmt.Print(" ")
         }
      }
      fmt.Println()
   }
}