package main

import(
  "fmt"
)

func main() {
  fmt.Println()
  
  // for like while
  k := 6
  for k > 1 {
    fmt.Println(k)
  }
  
  for condition := true; condition {
    fmt.Println("Hehe")
    condition = false
  }
  
  for i = 0; i < 10; i+=2 {
    fmt.Println(i)
  }
  
  // range
  for key, value := range "Bing Chilling" {
    fmt.Println(key, value)
  }
}