package main

import(
  "fmt"
)

func main() {
  var age int8 = 56 // (u)int(16|32|64|128)   : default = 0
  var pi float32 = 3.14 // float64   : default 0.0
  var c complex128 = 5 + 5i  // default = 0 + 0i
  fmt.Println(age, pi, c)
  
  var character byte = 'M'  // default = 0
  var emoji rune =  '🫧'// default = 0
  var name string = "manus"  // default = ""
  fmt.Println(character, emoji, name)
  
  // byte & rune print integer values by default
  fmt.Println(string(character), string(emoji))
  
}