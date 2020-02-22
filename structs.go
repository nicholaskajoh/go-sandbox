package main

import "fmt"

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel int16
	topSpeedKmh   float64
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / 65535)
}
func (c *car) newTopSpeed(s float64) {
	c.topSpeedKmh = s
}

func main() {
	myCar := car{gasPedal: 2212, brakePedal: 0, steeringWheel: 1232, topSpeedKmh: 98.0}
	fmt.Println(myCar.gasPedal, myCar.brakePedal, myCar.steeringWheel, myCar.topSpeedKmh)
	fmt.Println(myCar.kmh())
	myCar.newTopSpeed(400)
}
