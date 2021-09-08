# gpio0

gpio0 Go GPIO library for the Raspberry Pi.
It is inspired by gpiozero, the Raspberry Pi GPIO library for Python

## Types

### type [Button](/button.go#L9)

`type Button struct { ... }`

Button represents a button connected to a gpio pin

#### func OpenButton

`func OpenButton(pin int) (*Button, error)`

#### func (*Button) [Pressed](/button.go#L38)

`func (btn *Button) Pressed() bool`

Pressed returns true if the button is pressed (active low)
To keep usage of this function simple (e.g. interfaces),
it will panic if an error is incurred while writing to the gpio pin

#### func (*Button) [Value](/button.go#L26)

`func (btn *Button) Value() (int, error)`

Value returns the value read at the gpio pin

### type [LED](/led.go#L9)

`type LED struct { ... }`

LED represents an LED connected to a gpio pin

#### func OpenGPIO
`func OpenGPIO(pin int, activeHi bool) (*LED, error)`

#### func OpenLocal
`func OpenLocal(num int) (*LED, error)`

#### func (*LED) [ActiveHi](/led.go#L85)

`func (led *LED) ActiveHi() bool`

ActiveHi is true if the output is active high

#### func (*LED) [Close](/led.go#L37)

`func (led *LED) Close() error`

Close closes the file in /sys/class/gpio

#### func (*LED) [Filename](/led.go#L90)

`func (led *LED) Filename() string`

Filename is the name of the underlying file used for output

#### func (*LED) [Local](/led.go#L80)

`func (led *LED) Local() bool`

Local is true if LED is an onboard LED

#### func (*LED) [Off](/led.go#L72)

`func (led *LED) Off()`

Off turns off led (active low)

#### func (*LED) [On](/led.go#L64)

`func (led *LED) On()`

On turns on led (active low)

#### func (*LED) [Pin](/led.go#L95)

`func (led *LED) Pin() int`

Pin is the gpio pin number or led number for onboard leds

#### func (*LED) [Set](/led.go#L42)

`func (led *LED) Set(v bool) error`

Set sets the LED to v (active low)
package gpio0 // import "github.com/jaymzee/gpio0"


type Button struct{ ... }
    func OpenButton(pin int) (*Button, error)
type LED struct{ ... }
    func OpenGPIO(pin int, activeHi bool) (*LED, error)
    func OpenLocal(num int) (*LED, error)
