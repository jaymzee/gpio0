package gpio0

import (
	"fmt"
	"os"
)

// Button represents a button connected to a gpio pin
type Button struct {
	file *os.File
}

// OpenButton creates a new Button for the gpio pin number given
func OpenButton(pin int) (*Button, error) {
	filename := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return &Button{file}, nil
}

// Value returns the value read at the gpio pin
func (btn *Button) Value() (int, error) {
	buf := make([]byte, 1)
	_, err := btn.file.ReadAt(buf, 0)
	if err != nil {
		return 0, err
	}
	return int(buf[0] - '0'), nil
}

// Pressed returns true if the button is pressed (active low)
// To keep usage of this function simple (e.g. interfaces),
// it will panic if an error is incurred while writing to the gpio pin
func (btn *Button) Pressed() bool {
	val, err := btn.Value()
	if err != nil {
		panic(err)
	}
	return val == 0
}
