package gpio0

import (
	"fmt"
	"github.com/jaymzee/morse"
	"os"
)

// LED represents an LED connected to a gpio pin
type LED struct {
	local    bool
	activeHi bool
	pin      int
	file     *os.File
}

// OpenGPIO returns an LED for the gpio pin at /sys/class/gpio/gpio{n}
func OpenGPIO(pin int, activeHi bool) (*LED, error) {
	filename := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)
	file, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		return nil, err
	}
	return &LED{false, activeHi, pin, file}, nil
}

// OpenLocal returns an LED for the builtin LEDs
func OpenLocal(num int) (*LED, error) {
	filename := fmt.Sprintf("/sys/class/leds/led%d/brightness", num)
	file, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		return nil, err
	}
	return &LED{local: true, activeHi: true, pin: num, file: file}, nil
}

// Close closes the file in /sys/class/gpio
func (led *LED) Close() error {
	return led.file.Close()
}

// Set sets the LED to v (active low)
func (led *LED) Set(v bool) error {
	if !led.activeHi {
		v = !v
	}
	i := 0
	if v {
		if led.local {
			i = 255
		} else {
			i = 1
		}
	}
	bytes := []byte(fmt.Sprintf("%d", i))
	_, err := led.file.WriteAt(bytes, 0)
	return err
}

func (led *LED) Write(p []byte) (int, error) {
	return morse.Send(led, p)
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

// Local is true if LED is an onboard LED
func (led *LED) Local() bool {
	return led.local
}

// ActiveHi is true if the output is active high
func (led *LED) ActiveHi() bool {
	return led.activeHi
}

// Filename is the name of the underlying file used for output
func (led *LED) Filename() string {
	return led.file.Name()
}

// Pin is the gpio pin number or led number for onboard leds
func (led *LED) Pin() int {
	return led.pin
}
