package main

import "fmt"

func main() {
	base := 21.0
	total := 0.0
	currentaward := 50.0
	for currentaward > 0{
		a := base*currentaward
		total += a
		//currentaward /= 2 除法效率低
		currentaward *= 0.5
	}
	fmt.Printf("比特币总量为%v",total)

}
