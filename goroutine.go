package main

import (
	"fmt"
	"time"
)

func makeCoffee(name string) {
	fmt.Println("making  coffee for ", name)
	time.Sleep(2 * time.Second)
	fmt.Println("coffee is ready", name)
}

func main() {
	go makeCoffee("Rakib")
	go makeCoffee("Sakib")
	time.Sleep(3 * time.Second)
	fmt.Println("All coffee are ready")
}
