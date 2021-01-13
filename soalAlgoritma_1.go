package main

import "fmt"

func main() {
   var number int = 7
   var row int
   var col int

   for row = 0 ; row < number ; row++{
      for col = 0 ; col < number-row ; col++{
         if(col == 0){
            fmt.Print(row+1)
         }else{
            fmt.Print(0)
         }
      }
      fmt.Println()
   }
}