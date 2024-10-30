package main

import (
	"gotchi/assets"
	"image/color"

	"machine"
	"time"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

var black = color.RGBA{1, 1, 1, 255}
var btnA, btnB, btnC, btnUp, btnDown machine.Pin

const (
	WIDTH  = 296
	HEIGHT = 128
)

func DisplayMainMenu(display uc8151.Device) {
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, HEIGHT-10, "Feed", black)
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 110, HEIGHT-10, "Clean", black)
	tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 220, HEIGHT-10, "Sleep", black)
}

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display := uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Rotation:    uc8151.ROTATION_270,
		Speed:       uc8151.MEDIUM,
		Blocking:    true,
		FlickerFree: true,
	})

	display.ClearDisplay()

	animation := [][]uint8{
		assets.Idle1,
		assets.Idle2,
		assets.Idle3,
	}

	var current int
	var pressed bool

	btnA = machine.BUTTON_A
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		display.ClearBuffer()

		DisplayMainMenu(display)

		if btnA.Get() && !pressed {
			tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, 50, "Apple", black)
			pressed = true
		} else {
			pressed = false
		}

		display.DrawBuffer(20, (WIDTH-64)/2, 64, 64, []uint8(animation[current]))

		display.Display()
		display.WaitUntilIdle()

		time.Sleep(500 * time.Millisecond)
		current += 1
		if current == len(animation) {
			current = 0
		}
	}
}
