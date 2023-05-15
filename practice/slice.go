package main

import "fmt"

func main() {
  // slice is a homogeneous and linear ds
  // it can dynamically resize itself to
  // accomodate new values
  
  // declaration of slice is similar to arrays
  // but we do not specify size
  
  // you can access and change values in a slice 
  // in the same manner you can in array
  // and can iterate them like you do arrays
  
  // they can be multidimensional like arrays
  
  var s []int // by default s == nil
  fmt.Printf("%v: %T\n", s, s)
  
  
  s = []int{1, 2, 6}
  fmt.Printf("%v: %T\n", s, s)
  
  s = append(s, 4, 2, 0, 5)
  fmt.Printf("%v: %T\n", s, s)
  
  fmt.Println(len(s))
  fmt.Println(cap(s))
  
  // pre allocate a certain size
  s = make([]int, 45) // all indices will have default value of the type
  
  fmt.Printf("%v: %T\n", s, s)
  fmt.Println(len(s))
  fmt.Println(cap(s))
  
  // slicing
  // work similar to slicing a list in python
  
  s = []int{1, 2, 6, 12, 34, 56}
  fmt.Println(s)
  fmt.Println(s[2:5])
  fmt.Println(s[:5])
  fmt.Println(s[5:])
  
}