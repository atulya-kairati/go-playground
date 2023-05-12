package main

import (
  "fmt"
)

func main() {
  num := 5
  var num_pointer *int
  
  num_pointer = &num
  
  fmt.Printf("value: %d, type: %T\n", num, num)
  fmt.Printf("value: %d, type: %T\n", num_pointer, num_pointer)
  
  square(&num) // pass by reference
  fmt.Printf("value: %d, type: %T\n", num, num)
  
}

func square(ptr *int)  {
  *ptr = *ptr * *ptr
}