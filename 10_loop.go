package main

import "fmt"

// prints out 0 to 9
func main() {
  // i:=0

  // for i<10 {
    // fmt.Println(i)
    // i++
    // i += 1 or could be
    // i += 5

  x := 5
  for {
    fmt.Println("Do stuff", x)
    x+= 3
    if x > 25 {
      break
    }
  }
}

// := defining variable , = reassigning variable


