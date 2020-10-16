package gpio0

import (
	"fmt"
	"io"
	"os"
)

// LED represents an LED connected to a gpio pin
type LED struct {
	pin  int
	file *os.File
}

// NewLED returns a new LED for the gpio pin given
func NewLED(pin int) *LED {
	filename := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)
	file, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		panic(err)
	}
	return &LED{pin: pin, file: file}
}

// Set sets the LED to v
func (led *LED) Set(v bool) {
	if v {
		led.On()
	} else {
		led.Off()
	}
}

// On turns on led (active low)
func (led *LED) On() {
	led.file.Seek(0, io.SeekStart)
	led.file.Write([]byte{'0'})
}

// Off turns off led (active low)
func (led *LED) Off() {
	led.file.Seek(0, io.SeekStart)
	led.file.Write([]byte{'1'})
}
