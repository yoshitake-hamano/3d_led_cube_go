package led_cube_go

import (
	"testing"
)

const LED_URL = "127.0.0.1:9001"
const SHOW_TIMES = 100

func TestSetZero(t *testing.T) {
	showSingleColor(0)
}

func TestSetRed(t *testing.T) {
	showSingleColor(255 << 16)
}

func TestSetGreen(t *testing.T) {
	showSingleColor(255 << 8)
}

func TestSetBlue(t *testing.T) {
	showSingleColor(255)
}

func showSingleColor(rgb int) {
	SetUrl(LED_URL)
	for x := 0; x < LED_WIDTH; x++ {
		for y := 0; y < LED_HEIGHT; y++ {
			for z := 0; z < LED_DEPTH; z++ {
				SetLed(x, y, z, rgb)
			}
		}
	}
	for i := 0; i < SHOW_TIMES; i++ {
		Show()
	}
}
