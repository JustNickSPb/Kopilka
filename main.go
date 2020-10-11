package main

import (
	"fmt"
	"math/rand"
)

var kopilka float64
var counter int

func dropTheCoin(coin int) {
	fmt.Printf("Бросили %v центов\n", coin)
	kopilka += float64(coin) / 100
	counter++
}

func main() {

	fiveCents := 5
	tenCents := 10
	quater := 25

	for {
		coinDropped := rand.Intn(3)
		switch coinDropped {
		case 0:
			dropTheCoin(fiveCents)
		case 1:
			dropTheCoin(tenCents)
		case 2:
			dropTheCoin(quater)
		}
		fmt.Printf("Сейчас в копилке $%5.2f\n", kopilka)
		if kopilka >= 25 {
			fmt.Printf("В копилке %v монет.\n", counter)
			break
		}
	}
}
