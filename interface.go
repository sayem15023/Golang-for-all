package main

import "fmt"

type Device interface {
	TurnOn() string  // Expecting a return type of string
	TurnOff() string // Expecting a return type of string
}

type SmartLight struct{} // Fixed typo here: SmartLigt -> SmartLight

func (s SmartLight) TurnOn() string {
	return "Smart light is turned on"
}

func (s SmartLight) TurnOff() string {
	return "Smart light is turned off"
}

type SmartSpeaker struct{}

func (s SmartSpeaker) TurnOn() string {
	return "Smart speaker is turned on"
}

func (s SmartSpeaker) TurnOff() string {
	return "Smart speaker is turned off"
}

// SmartDevice can operate any device that satisfies the Device interface
func SmartDevice(d Device) {
	fmt.Println(d.TurnOn())  // Call TurnOn method
	fmt.Println(d.TurnOff()) // Call TurnOff method
}

// func main() {
// 	light := SmartLight{}
// 	speaker := SmartSpeaker{}

// 	// Operate both devices using the same function
// 	fmt.Println("Operating the Smart Light:")
// 	SmartDevice(light)

// 	fmt.Println("\nOperating the Smart Speaker:")
// 	SmartDevice(speaker)
// }
