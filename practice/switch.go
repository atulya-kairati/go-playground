package main

import (
  "fmt"
)


func main() {
  
  n := 4
  
  switch(n)  {
  case 1:
    fmt.Println("one")
    
  case 2, 3, 4:
    fmt.Println("2,3,4")
    
  }
  
  switch  {
  case n > 0:
    fmt.Println("+")
    
  case n < 0:
    fmt.Println("-")
  
  default:
    fmt.Println("0")
  }
  
}