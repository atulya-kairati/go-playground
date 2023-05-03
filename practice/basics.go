package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hi gophers!")
  
  const pi = 3.14
  const (
    a = "a"
    b = "b"
  )
  fmt.Println(pi, a,b)
  
  const (
    x = 5
    y
    z
  ) // x = y = z = 5
  fmt.Println(x,y,z)
  
  const (
    d = iota // + 5
    e
    f
  ) // d = 0, e = 1, f = 2
  fmt.Println(d,e,f)
  

  var s, t string = "Chunni Lal", "Ttt"
  fmt.Println(s, t)
  
  ee := 4
  xx, yy := 3, "r"
  fmt.Println(ee, xx, yy)
  
  var (
    one = 1
    two = 2
  )
  
  fmt.Println(one, two)

  var arr = make([]int, 5)
  arr[0] = 1
  arr[2] = 3
  fmt.Println(arr)

  bb := append(arr, 3, 4, 5)
  fmt.Println(bb)
  
  var m = make(map[int]string)
  m[1] = "Monday"
  m[2] = "Tuesday"
  fmt.Println(m)
  
  var ptr = &arr
  fmt.Println(ptr)
  fmt.Println(*ptr)
  *ptr = bb[4:]
  fmt.Println(arr)
  
}