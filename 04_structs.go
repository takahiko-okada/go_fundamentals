
package main

import "fmt"

type car struct {
  // uint16 type (0 to 65535)
  gas_pedal uint16
  brake_pedal uint16
  steering_wheel int16
  top_speed_kmh float64
}


func main() {
  a_car := car{gas_pedal: 22341,
               brake_pedal: 0,
               steering_wheel: 12561,
               top_speed_kmh: 225.0}

  // or a_car := {22341, 0, 12562, 225.0}
  fmt.Println(a_car.gas_pedal)
  fmt.Println(a_car.top_speed_kmh)
}
