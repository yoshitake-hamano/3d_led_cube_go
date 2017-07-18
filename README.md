# 3d_led_cube_go

## Install

```
$ go get github.com/yoshitake-hamano/3d_led_cube_go
```

## Sample

```
package main

import (
	"github.com/yoshitake-hamano/3d_led_cube_go"
	"log"
)

const LED_URL = "127.0.0.1:9001"
const SHOW_TIMES = 100

func main() {
	log.Println("Show black")
	showSingleColor(0)

	log.Println("Show red")
	showSingleColor(255 << 16)

	log.Println("Show green")
	showSingleColor(255 << 8)

	log.Println("Show blue")
	showSingleColor(255 << 0)
}

func showSingleColor(rgb int) {
	led_cube_go.SetUrl(LED_URL)
	for x := 0; x < led_cube_go.LED_WIDTH; x++ {
		for y := 0; y < led_cube_go.LED_HEIGHT; y++ {
			for z := 0; z < led_cube_go.LED_DEPTH; z++ {
				led_cube_go.SetLed(x, y, z, rgb)
			}
		}
	}
	for i := 0; i < SHOW_TIMES; i++ {
		led_cube_go.Show()
	}
}

```

## Run sample

```
$ cd <<3d_led_cube_receiver>>/make/receiver &
$ go run
```

## How to execute 3d_led_cube_go's unit test

```
$ cd <<3d_led_cube_receiver>>/make/receiver &
$ cd <<3d_led_cube_go>>
$ go test -v
=== RUN   TestSetZero
--- PASS: TestSetZero (0.05s)
=== RUN   TestSetRed
--- PASS: TestSetRed (2.05s)
=== RUN   TestSetGreen
--- PASS: TestSetGreen (1.06s)
=== RUN   TestSetBlue
--- PASS: TestSetBlue (2.07s)
=== RUN   TestSetUrl
--- PASS: TestSetUrl (0.00s)
=== RUN   TestSetLed
--- PASS: TestSetLed (0.01s)
PASS
ok  	_3d_led_cube/3d_led_cube_go	5.254s
```
