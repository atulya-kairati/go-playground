package main

import (
  "fmt"
)

type Animal struct {
  name string
  class string
  speak string
  legs uint8
}

func (a Animal) Speak() {
  fmt.Println(a.name, ":", a.speak)
}

func main() {
  // var dog = new(Animal)
  var dog Animal
  dog.name = "dog"
  dog.speak = "Bhark! Teri maa ka"
  dog.class = "Mammal"
  dog.legs = 4
  
  dog.Speak()
}