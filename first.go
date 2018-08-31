package main

import "fmt"

const usix float64 = 65535
const kmhtom float64 = 1.6

type car struct{
	gas_pedal uint16
	brake_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}

func (c car) kmh() float64{
	return float64(c.gas_pedal)*(c.top_speed_kmh/usix)
}

func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main(){
	a_car := car{gas_pedal: 22341,
	             brake_pedal: 0,
	             steering_wheel: 12561,
	             top_speed_kmh: 225.0}

	             fmt.Println(a_car.gas_pedal)
	             fmt.Println(a_car.kmh())
	             a_car.new_top_speed(500)
	             fmt.Println(a_car.kmh())
	             
}