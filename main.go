package main

import (
	"fmt"
)

func main() {
	randomSeed()
	outpost := NewOutpost()
	fmt.Println(outpost)
	for outpost.Population > 5000 {

		outpost.step1()
		outpost.step2()
		fmt.Println("Outpost Balance:", outpost.Credits)
	}
}
