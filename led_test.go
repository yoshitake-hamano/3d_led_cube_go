package led_cube_go

import (
	"reflect"
	"testing"
)

var ZERO_LED_BUFFER = make([]byte, LED_WIDTH*LED_HEIGHT*LED_DEPTH*LED_COLOR)

func TestSetUrl(t *testing.T) {
	const target_url = "127.0.0.1"
	SetUrl(target_url)
	if getUrl() != target_url {
		t.Fatal("failed to SetUrl()")
	}
}

func TestSetLed(t *testing.T) {
	Clear()
	buffer := getLedBuffer()
	if !reflect.DeepEqual(buffer, ZERO_LED_BUFFER) {
		t.Fatal("Invalid initial led buffer")
	}
	// x, y, z, rgb
	SetLed(0, 0, 0, 1)
	if reflect.DeepEqual(buffer, ZERO_LED_BUFFER) {
		t.Fatal("failed to SetLed()")
	}
	Clear()
	if !reflect.DeepEqual(buffer, ZERO_LED_BUFFER) {
		t.Fatal("Invalid initial led buffer")
	}
}
