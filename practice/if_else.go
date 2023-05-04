package main

import (
  "fmt"
)

func main() {
  meters := 5698
  
  // km will only be present in scope of if
  if km := meters/1000; km < 2{
    fmt.Println("Near")
  } else if km < 4 { // else should be placed in the same line as the closing } bracket
    fmt.Println("Take bike")
  } else {
    fmt.Println("Far")
  }
}
