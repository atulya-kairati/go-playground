package main

import "fmt"

func main() {
  // map is very similar to python dict
  
  // declaring a map
  var m map[int]string // map[(key type)](value type)
  // default m = []
  fmt.Println(m)
  
  // intialising maps 
  var roll map[int]string
  roll = map[int]string{
    1: "Manus Chaubey",
    2: "Chootdi Lal",
  }
  fmt.Println(roll)
  
  // modifying a value
  roll[2] = "Topi Topeshwar"
  fmt.Println(roll)
  
  // inserting a new pair
  roll[3] = "Chumban Singh"
  fmt.Println(roll)
  
  // deleting a key value pair
  delete(roll, 1)
  fmt.Println(roll)
  
  // accessing an element
  v := roll[2]
  fmt.Println(v)
  
  // accessing an element in go has one quirk
  // is the key used doesn't exist then the 
  // map will return an empty string
  // and we can't differentiate if the value was 
  // an empty or the if the key simply didn't exist
  
  // in go when we access a value we also get an
  // an additional value telling is if the key 
  // existed or not
  v1, exists := roll[21]
  if exists {
    fmt.Println(v1)
  }
  
  v2, exists := roll[3]
  if exists {
    fmt.Println(v2)
  }
  
  // iterating over a map
  for k, v:= range roll {
    fmt.Println(k, v)
  }
  
  // size of map
  fmt.Println(len(roll))
  
  // pre allocation of a map
  m = make(map[int]string, 5) // allocated space for 5 key value pairs
  
  // key type can be anything even an array
}