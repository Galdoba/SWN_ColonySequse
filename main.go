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
		fmt.Println("Year:", i/12, "Month:", i%12+1)
		outpost.step1()
		outpost.step2()
		outpost.step3()
		outpost.step4()
		outpost.step5()
		outpost.Status()
		fmt.Println("")
		if i > 100 {
			break
		}
	}
}
