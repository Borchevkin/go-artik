package main

import "fmt"
import "time"
import "github.com/lab409/go-artik/gpio"

func main() {
	fmt.Println("\nExample. Blink gpio125 (Arduino 8 pin)")
	fmt.Println("For exit press Ctrl+C\n")

	/* Create new pin */
	pin := gpio.NewPin(125, gpio.GPIO_OUT)
	
	/* Just to be on the safe side - pull-up the pin */
	pin.Set()

	/* Infinite loop for blinking */
	for {
		/* Blink mat completed by Set/Clear logic ... */
		pin.Set()
		time.Sleep(200 * time.Millisecond)
		pin.Clear()
		time.Sleep(200 * time.Millisecond)

		/* Or by Toggle logic */	
		pin.Toggle()
		time.Sleep(200 * time.Millisecond)
	}
}
