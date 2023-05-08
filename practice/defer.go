package main

import(
  "fmt"
)


func main() {
  // any statement marked with [defer]
  // will be executed when the surrounding
  // function ends (main in this case)
  // even if the program encounters an error
  // or crashes
  
  
  defer fmt.Println(0)
  
  first()
  defer second()
  third()
}


func first()  {
  fmt.Println(1)
}

func second()  {
  fmt.Println(2)
}

func third()  {
  fmt.Println(3)
}
