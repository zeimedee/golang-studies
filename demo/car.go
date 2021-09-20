package main

import "fmt"

type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

func (car *Car) Drive() Car {
	if car.battery < car.speed {
		return Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      car.battery,
			distance:     0,
		}
	}
	return Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		battery:      car.battery - car.batteryDrain,
		distance:     car.speed,
	}
}

func main() {
	car := Car{
		speed:        5,
		batteryDrain: 2,
		battery:      100,
		distance:     0,
	}
	output := car.Drive()
	fmt.Println(output)
}
