package main

// #cgo LDFLAGS: -lxdo
// #include <xdo.h>
import "C"

import (
	"flag"
	"os/exec"
	"time"
)

var (
	xdo           = C.xdo_new(nil)
	y, s          C.int
	width, height C.uint
)

func main() {
	refreshPeriod := flag.Duration("refreshPeriod", 500*time.Millisecond, "Refresh period")
	barHeight := flag.Int("barHeight", 30, "Bar height in pixels")
	flag.Parse()

	barHidden := true
	for {
		if barHidden && cursorInBar(*barHeight) {
			barHidden = false
			showBar()
		} else if !barHidden && !cursorInBar(*barHeight) {
			barHidden = true
			hideBar()
		}
		time.Sleep(*refreshPeriod)
	}
}

func cursorInBar(barHeight int) bool {
	C.xdo_get_mouse_location2(xdo, nil, &y, &s, nil)
	C.xdo_get_viewport_dimensions(xdo, &width, &height, s)
	if y > C.int(height)-C.int(barHeight) {
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
