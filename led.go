package gpio0

import (
	"fmt"
	"os"
	"github.com/jaymzee/morse"
)

// LED represents an LED connected to a gpio pin
type LED struct {
	pin  int
	file *os.File
}

// OpenLED returns an LED for the gpio pin at /sys/class/gpio/gpio{n}
func OpenLED(pin int) (*LED, error) {
	filename := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)
	file, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		return nil, err
	}
	return &LED{pin: pin, file: file}, nil
}

// Close closes the file in /sys/class/gpio
func (led *LED) Close() error {
	return led.file.Close()
}

// Set sets the LED to v (active low)
func (led *LED) Set(v bool) error {
	var b byte = '1'
	if v {
		b = '0'
	}
	_, err := led.file.WriteAt([]byte{b}, 0)
	return err
}

func (led *LED) Write(p []byte) (int, error) {
	morse.Send(led, p)
	return len(p), nil
}

// The functions below satisfy a very simple On/Off interface.
// To keep the error handling very simple for users of this interface,
// these functions panic on error.

// On turns on led (active low)
func (led *LED) On() {
	err := led.Set(true)
	if err != nil {
		panic(err)
	}
}

// Off turns off led (active low)
func (led *LED) Off() {
	err := led.Set(false)
	if err != nil {
		panic(err)
	}
}
