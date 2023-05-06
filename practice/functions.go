package main

import (
  "fmt"
)


func main() {
  
  // functions are first class citizens
  greeter := hello
  fmt.Printf("%T\n", hello)
  greet_all(greeter, "me", "you", "others")
  
  greet_all(hello, "Manus", "Atulya", "Chootad")
  
  fmt.Println(logical_or(true, false))
  
  dm := get_divmod()
  fmt.Printf("%T\n", dm)
  div, mod := dm(10, 3)
  fmt.Println(div, mod)
  
  // anon func
  af := func() { fmt.Println("Me gusta anonymous") }
  
  af()
  
  greet_all(func(name string) { fmt.Println(name, "chutiya") },"atulya", "manus", "chootadi lal")
  
}


func hello(name string){
  fmt.Println("Hello!", name)
}


func greet_all(greeting func(string), people ...string){
  for _, person := range people {
    greeting(person)
  }
}


// function with return type
func logical_or(a bool, b bool) bool {
  return a || b
}


// function with multiple return type (they can be of different types)
func divmod(a int, b int) (int, int) {
  return a / b, a % b
}


// function with function as a return type
func get_divmod() func(int, int) (int, int) {
  return divmod
}