package main

// fmt is short for format.
import  ("fmt")

// argument and datatype then, the data type of the return value
func add(x,y float64) float64 {
  return x + y
}

// when returning more than one values, specify the
// data type for each
func multiple(a,b string) (string,string) {
  return a,b
}

// Go detects default data type in case of float, float64, not 32
func main() {
  num1, num2 := 5.6, 9.5
// refactored
  // var num1, num2 float64 = 5.6, 9.5
  // refactored
    // var num1 float64 = 5.6
    // var num2 float64 = 9.5

  w1, w2 := "Hey", "there"

  fmt.Println(multiple(w1, w2))
  fmt.Println(add(num1, num2))

  var a int = 62
  var x := a
  // a assigned to x as integer
  var b float64 = float64(a)
  // converts a into float64 and assign it to b
}
