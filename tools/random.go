package tools

import (
	"fmt"
	"math/rand"
	"time"
)

func Random_binary(duration int, iter int) {

	if iter == 0 {
		for {
			bin := rand.Intn(2) //generate random binary
			fmt.Print(bin)
			time.Sleep(time.Duration(duration) * time.Microsecond)
		}
	} else {
		loop := 0
		for loop < iter {
			bin := rand.Intn(2) //generate random binary
			fmt.Print(bin)
			time.Sleep(time.Duration(duration) * time.Microsecond)
			loop++
			if loop == iter {
				print("\n")
			}
		}
	}
}
