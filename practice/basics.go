package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hi gophers!")

  var s string = "Chunni Lal"
  fmt.Println(s)

  var arr = make([]int, 5)
  arr[0] = 1
  arr[2] = 3
  fmt.Println(arr)

  b := append(arr, 3, 4, 5)
  fmt.Println(b)
  
  var m = make(map[int]string)
  m[1] = "Monday"
  m[2] = "Tuesday"
  fmt.Println(m)
  
  var ptr = &arr
  fmt.Println(ptr)
  fmt.Println(*ptr)
  *ptr = b[4:]
  fmt.Println(arr)
  
}