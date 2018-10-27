package main

import "fmt"

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	sterring_wheel float64
	top_speed_kmh  float64
}

func main() {
	a_car := car{gas_pedal: 1212,
		brake_pedal:    199,
		sterring_wheel: 12.21,
		top_speed_kmh:  80}

	fmt.Println(a_car.sterring_wheel)
}
