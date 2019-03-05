package main

import (
	"fmt"
)

func main() {
	randomSeed()
	outpost := NewOutpost()
	fmt.Println(outpost)
	i := 0
	for outpost.Population > 5000 {
		i++
		fmt.Println("Month", i)
		outpost.step1()
		outpost.step2()
		fmt.Println("Outpost Balance:", outpost.Credits)
	}
}
