package main

import "fmt"
import "time"
import "github.com/lab409/go-artik/gpio"

func main() {
	fmt.Println("\nExample. Set gpio125 (Arduino 8 pin) if gpio124 (Arduino 7 pin) is HIGH")
	fmt.Println("For exit press Ctrl+C\n")

	/* Create pins for led and button */
	led := gpio.NewPin(125, gpio.OUT)
	button := gpio.NewPin(124, gpio.IN)
	
	/* Clear the led */
	led.Clear()

	/* Infinite loop with polling of the button state and setting led */
	for {
		if (button.GetState() == gpio.HIGH) {
			led.Set()
		} else {
			led.Clear()
		}

		time.Sleep(100 * time.Millisecond)
	}
}
