package led_cube_go

import (
	"log"
	"net"
)

const LED_WIDTH = 16
const LED_HEIGHT = 32
const LED_DEPTH = 8
const LED_COLOR = 3
const LED_RED = 0
const LED_GREEN = 1
const LED_BLUE = 2

var ledUrl string
var ledBuffer = make([]byte, LED_WIDTH*LED_HEIGHT*LED_DEPTH*LED_COLOR)
var sem = make(chan struct{}, 1)

func SetUrl(url string) {
	ledUrl = url
}

func SetLed(x, y, z, rgb int) {
	if x < 0 || LED_WIDTH <= x {
		log.Fatalf("invalid x : %d\n", x)
		return
	}
	if y < 0 || LED_HEIGHT <= y {
		log.Fatalf("invalid y : %d\n", y)
		return
	}
	if z < 0 || LED_DEPTH <= z {
		log.Fatalf("invalid z : %d\n", z)
		return
	}
	index := z*LED_COLOR + y*LED_DEPTH*LED_COLOR + x*LED_HEIGHT*LED_DEPTH*LED_COLOR
	sem <- struct{}{}
	ledBuffer[index+LED_RED] = byte(rgb >> 16)
	ledBuffer[index+LED_GREEN] = byte(rgb >> 8)
	ledBuffer[index+LED_BLUE] = byte(rgb >> 0)
	<-sem
}

func Clear() {
	sem <- struct{}{}
	for i, _ := range ledBuffer {
		ledBuffer[i] = 0
	}
	<-sem
}

func Show() {
	tcpAddr, err := net.ResolveUDPAddr("udp", ledUrl)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
		return
	}
	conn, err := net.DialUDP("udp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
		return
	}
	defer conn.Close()
	sem <- struct{}{}
	udpBuffer := rgb888toRGB565(ledBuffer)
	<-sem
	conn.Write(udpBuffer)
}

func getUrl() string {
	return ledUrl
}

func getLedBuffer() []byte {
	return ledBuffer
}

func rgb888toRGB565(rgb888 []byte) []byte {
	const RGB888_RGB_SIZE = 3
	const RGB565_RGB_SIZE = 2

	lengthOfRGB888 := len(rgb888)
	lengthOfRGB565 := lengthOfRGB888 / RGB888_RGB_SIZE * RGB565_RGB_SIZE
	rgb565 := make([]byte, lengthOfRGB565)
	for i := 0; i < len(rgb888); i += RGB888_RGB_SIZE {
		r := rgb888[i]
		g := rgb888[i+1]
		b := rgb888[i+2]
		indexOfRGB565 := i / RGB888_RGB_SIZE * RGB565_RGB_SIZE
		rgb565[indexOfRGB565] = r&0xF8 + g>>5
		rgb565[indexOfRGB565+1] = g&0x1C + b>>3
	}
	return rgb565
}
