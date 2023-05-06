package main

import (
  "fmt"
)


func main() {
  
  // functions are first class citizens
  greeter := hello
  fmt.Printf("%T\n", hello)
  greet_all(greeter, "me", "you", "others")
  
  greet_all(greeter, "Manus", "Atulya", "Chootad")
  
}


func hello(name string){
  fmt.Println("Hello!", name)
}

func greet_all(greeting func(string), people ...string){
  for _, person := range people {
    hello(person)
  }
}