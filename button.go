package gpio0

import (
	"fmt"
	"io"
	"os"
)

// Button represents a button connected to a gpio pin
type Button struct {
	pin  int
	file *os.File
}

// NewButton creates a new Button for the gpio pin number given
func NewButton(pin int) *Button {
	filename := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return &Button{pin: pin, file: file}
}

// Pressed return true if the button is pressed (active low)
func (btn *Button) Pressed() bool {
	buffer := make([]byte, 16)

	btn.file.Seek(0, io.SeekStart)
	count, err := btn.file.Read(buffer)
	if err != nil {
		panic(err)
	}

	return !(count > 0 && buffer[0] == '1')
}
