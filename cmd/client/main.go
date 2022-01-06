package main

import (
	"fmt"

	"github.com/cloudacademy/go-testing/pkg/rocket"
)

func main() {
	rocket := rocket.NewRocket("Saturn-V", "Nasa", 5, 25000, 100)
	fmt.Printf("%s rocket built...\n", rocket.Name)

	err := rocket.Ignite()
	if err != nil {
		panic("ignition failed!")
	}
	fmt.Printf("%s rocket ignited...\n", rocket.Name)

	currentSpeed := 0

	currentSpeed, _ = rocket.ThrottleUp(1000)
	fmt.Printf("%s current speed: %d\n", rocket.Name, currentSpeed)

	currentSpeed, _ = rocket.ThrottleUp(100)
	fmt.Printf("%s current speed: %d\n", rocket.Name, currentSpeed)

	currentSpeed, _ = rocket.ThrottleDown(10)
	fmt.Printf("%s current speed: %d\n", rocket.Name, currentSpeed)

	currentSpeed, _ = rocket.ThrottleUp(100)
	fmt.Printf("%s current speed: %d\n", rocket.Name, currentSpeed)

	currentSpeed, _ = rocket.ThrottleDown(5)
	fmt.Printf("%s current speed: %d\n", rocket.Name, currentSpeed)

	fmt.Printf("%s current speed: %d\n", rocket.Name, currentSpeed)
}
