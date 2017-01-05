package main

// #cgo LDFLAGS: -lxdo
// #include <xdo.h>
import "C"

import (
	"os/exec"
	"time"
)

const (
	refreshPeriod = 500 * time.Millisecond
	barHeight     = 30 // Bar height in pixels
)

var (
	xdo           = C.xdo_new(nil)
	y, s          C.int
	width, height C.uint
)

func main() {
	barHidden := true
	for {
		if barHidden && cursorInBar() {
			barHidden = false
			showBar()
		} else if !barHidden && !cursorInBar() {
			barHidden = true
			hideBar()
		}
		time.Sleep(refreshPeriod)
	}
}

func cursorInBar() bool {
	C.xdo_get_mouse_location2(xdo, nil, &y, &s, nil)
	C.xdo_get_viewport_dimensions(xdo, &width, &height, s)
	if y > C.int(height-barHeight) {
		return true
	}
	return false

}

func showBar() {
	showBar := exec.Command("i3-msg", "-q", "bar", "hidden_state", "show")
	showBar.Start()
}

func hideBar() {
	hideBar := exec.Command("i3-msg", "-q", "bar", "hidden_state", "hide")
	hideBar.Start()
}
