// https://www.youtube.com/watch?v=93f9_bJQdHk&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX&index=8

// https://www.youtube.com/watch?v=i3o4G4bmqPc

package main

import "fmt"

// max value of uint16
const usixteenbitmax float64 = 65535
// kmh in mile
const kmh_multiple float64 = 1.60934

type car struct {
  // uint16 type (0 to 65535)
  gas_pedal uint16
  brake_pedal uint16
  steering_wheel int16
  top_speed_kmh float64
}

//    c variable associated with car struct
// then begin function, define return datatype
func (c car) kmh() float64 {

  return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
  return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

// * <= pointer receiver, permanently change the value
func (c *car) new_top_speed(newspeed float64) {
  c.top_speed_kmh = newspeed
}

func newer_top_speed(c car, speed float64) car {
  c.top_speed_kmh = speed
  return c
}

func main() {
  a_car := car{gas_pedal: 65000,
               brake_pedal: 0,
               steering_wheel: 12561,
               top_speed_kmh: 225.0}

  // or a_car := {22341, 0, 12562, 225.0}
  fmt.Println(a_car.gas_pedal)
  fmt.Println(a_car.kmh())
  fmt.Println(a_car.mph())
  a_car.new_top_speed(500)
  a_car = newer_top_speed(a_car, 500)
  fmt.Println(a_car.kmh())
  fmt.Println(a_car.mph())

}

