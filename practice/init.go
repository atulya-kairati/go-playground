package main

import (
  "fmt"
)

// init method are first to be called 
// even before main

// init functions of the import packages 
// will be callee first


// There can be multiple init functions
// caled from top to bottom

func init() {
  fmt.Println("Intialising 1")
}

func main() {
  fmt.Println("main at last")
}

func init() {
  fmt.Println("Intialising 2")
}

func init() {
  fmt.Println("Intialising 3")
}