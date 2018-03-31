package main

import "fmt"

func main() {
  x := 15
  a := &x /*saving memory address to a*/
  fmt.Println(a)
  fmt.Println(*a) /*read memory address*/
  *a = 5
  fmt.Println(x)
  *a = *a**a
  fmt.Println(x)
  fmt.Println(*a)

}
