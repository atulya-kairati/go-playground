package main

import "fmt"

func main() {
  // in go every variable recieves a default value
  // 0 for int
  // 0.0 for float
  // "" for string
  
  // but we some types can't be provided with a concrete value
  // for example
  var ptr *int // it would be dangerous 
  // to save any address in the pointer
  
  // so instead of a value it is provided with `nil`
  // nil means absence of value
  
  // infact 
  // nil is a default value for
  // pointers
  // interfaces
  // slices
  // maps 
  // functions
  // channels ....etc
  
  fmt.Println(ptr)
}