package gpio

import "log"
import "io/ioutil"
import "strconv"
import "fmt"

const GPIO_LOW uint8 = 0
const GPIO_HIGH uint8 = 1
const GPIO_IN string = "in"
const GPIO_OUT string = "out"

type GPIO_Pin struct {
	number uint8
	direction string
	value uint8
}

func (pin *GPIO_Pin) exportPin() {
	/* For export pin for work with it need write to
	
	/sys/class/gpio/export

	number of the pin

	For example in bash (export 124 pin)

	echo 124 > /sys/class/gpio/export

	After this will be created file with this pin

	/sys/class/gpio/gpio124/
	*/

	buffer := []byte{pin.number}
	ioutil.WriteFile("/sys/class/gpio/export", buffer, 0644)
}

func (pin *GPIO_Pin) setDirection() {
	/* For setting direction of the pin need write to
	
	/sys/class/gpio/gpio<number>/direction

		- "in" for input (in this package for it there is GPIO_IN const)
		- "out" for output (in this package for it there is GPIO_OUT const)

	For example in bash (set 124 pin to out)

	echo out > /sys/class/gpio/gpio124/direction
	*/

	/* Set to buffer direction of the pin */
	buffer := []byte(pin.direction)
	/* Forming name of file accodring to sysfs */
	pinFile := fmt.Sprintf("/sys/class/gpio/gpio%s/direction", strconv.Itoa(int(pin.number)))
	/* Write direction stored in buffer to file */
	ioutil.WriteFile(pinFile, buffer, 0644)
}

func (pin *GPIO_Pin) init() {
	/* For init of GPIO pin need:
	1. Export the desired pin
	2. Set direction (in or out)
	*/

	/* Export pin */
	pin.exportPin()
	/* Set direction if the pin*/
	pin.setDirection()
}

func (pin *GPIO_Pin) deinit() {

}

func (pin *GPIO_Pin) setOutput() {
	/* For set out value of the pin need write to 

	/sys/class/gpio/gpio<number>/value

		- 1 for high level (in this package for it there is GPIO_HIGH const)
		- 0 for low level (in this package for it there is GPIO_LOW const)

	For example in bash (set high value on 124 pin)

	echo 1 > /sys/class/gpio/gpio124/value
	*/

	/* Set to buffer direction of the pin */
	buffer := []byte(strconv.Itoa(int(pin.value)))
	/* Forming name of file accodring to sysfs */
	pinFile := fmt.Sprintf("/sys/class/gpio/gpio%s/value", strconv.Itoa(int(pin.number)))
	/* Write direction stored in buffer to file */
	ioutil.WriteFile(pinFile, buffer, 0644)
}

func (pin *GPIO_Pin) Set() {
	pin.value = GPIO_HIGH
	pin.setOutput()
}

func (pin *GPIO_Pin) Clear() {
	pin.value = GPIO_LOW
	pin.setOutput()
}

func (pin *GPIO_Pin) Toggle() {
	if (pin.value == GPIO_LOW) {
		pin.value = GPIO_HIGH
	} else {
		pin.value = GPIO_LOW
	}
	pin.setOutput()
}

func (pin *GPIO_Pin) GetState() uint8 {
	/* For get in value of the pin need read 2 bytes to string from file

	/sys/class/gpio/gpio<number>/value

		- 1 for high level (in this package for it there is GPIO_HIGH const)
		- 0 for low level (in this package for it there is GPIO_LOW const)

	For example in bash (get value on 124 pin)

	cat /sys/class/gpio/gpio124/value
	*/

	/* Construct pin filename */
	pinFile := fmt.Sprintf("/sys/class/gpio/gpio%s/value", strconv.Itoa(int(pin.number)))
	/* Read data from pin file */
	buffer, err := ioutil.ReadFile(pinFile)
    if err != nil {
        log.Fatal(err)
    }

	result := uint8(buffer[0])
	/* Return read value */
	return result
}

func NewPin(number uint8, direction string) *GPIO_Pin {
	p := new(GPIO_Pin)
	p.number = number
	p.direction = direction

	p.init()

	return p
}
