package main

import (
	_ "embed"

	"image/color"

	"machine"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

type Food struct {
	animations map[string][]pixel.Image[pixel.Monochrome]
}

func newFood() Food {
	var f Food
	f.animations = make(map[string][]pixel.Image[pixel.Monochrome])

	var apple []pixel.Image[pixel.Monochrome]
	apple = append(apple, pixel.NewImageFromBytes[pixel.Monochrome](64, 64, Apple1))
	apple = append(apple, pixel.NewImageFromBytes[pixel.Monochrome](64, 64, Apple2))
	apple = append(apple, pixel.NewImageFromBytes[pixel.Monochrome](64, 64, Apple3))

	f.animations["apple"] = apple

	return f
}

type Gopher struct {
	animations map[string][]pixel.Image[pixel.Monochrome]
}

func newGopher() Gopher {
	var g Gopher
	g.animations = make(map[string][]pixel.Image[pixel.Monochrome])

	var idle []pixel.Image[pixel.Monochrome]
	idle = append(idle, pixel.NewImageFromBytes[pixel.Monochrome](64, 64, Idle1))
	idle = append(idle, pixel.NewImageFromBytes[pixel.Monochrome](64, 64, Idle2))
	idle = append(idle, pixel.NewImageFromBytes[pixel.Monochrome](64, 64, Idle3))

	g.animations["idle"] = idle

	var eating []pixel.Image[pixel.Monochrome]
	eating = append(eating, pixel.NewImageFromBytes[pixel.Monochrome](64, 64, Eating1))
	eating = append(eating, pixel.NewImageFromBytes[pixel.Monochrome](64, 64, Eating2))

	g.animations["eating"] = eating

	return g
}

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

	var current, currentFood int
	var pressed bool

	btnA = machine.BUTTON_A
	btnA.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	g := newGopher()
	currentAnimation := "idle"
	f := newFood()
	var eating bool

	for {
		display.ClearBuffer()

		DisplayMainMenu(display)

		if btnA.Get() && !pressed {
			//tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 10, 50, "Apple", black)
			currentAnimation = "eating"
			currentFood = 0
			current = 0
			pressed = true
			eating = true
		} else {
			pressed = false
		}

		display.DrawBitmap((WIDTH-64)/2, 20, g.animations[currentAnimation][current])
		if eating {
			display.DrawBitmap(10, 20, f.animations["apple"][currentFood])
			currentFood = (currentFood + 1) % len(f.animations["apple"])
			if currentFood == 0 {
				eating = false
				currentAnimation = "idle"
			}
		}

		display.Display()
		display.WaitUntilIdle()

		current = (current + 1) % len(g.animations[currentAnimation])
	}
}
