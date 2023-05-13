package main

import "fmt"

func main() {
  
  // type of array depends on its element type and size
  
  // declaring the array
  // all array elements are assigned their default values 
  var arr [5]string
  fmt.Printf("%T: %s\n", arr, arr)
  fmt.Println(len(arr))
  
  // creating an array using values
  arr2 := []int{3,56,7,3}
  fmt.Printf("%T: %d\n", arr2, arr2)
  fmt.Println(len(arr2))
  
  // the values will be assigned from the start
  arr3 := [10]int{2,2,56}
  fmt.Printf("%T: %d\n", arr3, arr3)
  
  // index can be used to assign values
  arr4 := [10]int{1, 4, 5: 45, 8: 72}
  fmt.Printf("%T: %d\n", arr4, arr4)
  
  // ... can be used to assign size based on no of elements
  arr5 := [...]int{3,56,7,3}
  fmt.Printf("%T: %d\n", arr5, arr5)
  
  // iterating over array
  for index, value := range arr4 {
    fmt.Println(index, ":", value)
  }
  
  // deep copying array
  arrcopy := arr5
  
  // accessing elements
  value := arrcopy[3]
  fmt.Println(value)
  arrcopy[2] = 34
  fmt.Println(arrcopy)
  
  // multidimensional array
  matrix := [2][2]int{{1, 2}, {2, 1}}
  fmt.Println(matrix)
  
  
}

/* OUTPUT
[5]string: [    ]
5
[]int: [3 56 7 3]
4
[10]int: [2 2 56 0 0 0 0 0 0 0]
[10]int: [1 4 0 0 0 45 0 0 72 0]
[4]int: [3 56 7 3]
0 : 1
1 : 4
2 : 0
3 : 0
4 : 0
5 : 45
6 : 0
7 : 0
8 : 72
9 : 0
3
[3 56 34 3]
[[1 2] [2 1]]
*/