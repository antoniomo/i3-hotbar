package main

import (
	"flag"
	"log"
	"os/exec"
	"sync"
	"time"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/randr"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xinerama"
	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/go-vgo/robotgo"
)

var (
	headLock sync.RWMutex
	heads    xinerama.Heads
)

func main() {
	refreshPeriod := flag.Duration("refreshPeriod", 500*time.Millisecond, "Refresh period")
	barHeight := flag.Int("barHeight", 30, "Bar height in pixels")
	flag.Parse()

	go updateScreens()

	// Wait for first update
	for {
		headLock.RLock()
		h := heads
		headLock.RUnlock()
		if h == nil {
			time.Sleep(100 * time.Millisecond)
		} else {
			break
		}
	}

	barHidden := true
	for {
		inBar := cursorInBar(*barHeight)
		if barHidden && inBar {
			barHidden = false
			showBar()
		} else if !barHidden && !inBar {
			barHidden = true
			hideBar()
		}
		time.Sleep(*refreshPeriod)
	}
}

func updateScreens() {

	// Connect to the X server using the DISPLAY environment variable.
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	Xutil, err := xgbutil.NewConnXgb(X)
	if err != nil {
		log.Fatal(err)
	}

	// Every extension must be initialized before it can be used.
	err = randr.Init(X)
	if err != nil {
		log.Fatal(err)
	}
	// Get the root window on the default screen.
	root := xproto.Setup(X).DefaultScreen(X).Root

	// Tell RandR to send us events. (I think these are all of them, as of
	// 1.3.)
	err = randr.SelectInputChecked(X, root,
		randr.NotifyMaskScreenChange|
			randr.NotifyMaskCrtcChange|
			randr.NotifyMaskOutputChange|
			randr.NotifyMaskOutputProperty).Check()
	if err != nil {
		log.Fatal(err)
	}

	// Wrap the root window in a nice Window type.
	rootWin := xwindow.New(Xutil, root)

	// Get the geometry of the root window.
	rgeom, err := rootWin.Geometry()
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Get the rectangles for each of the active physical heads.
		// These are returned sorted in order from left to right and
		// then top to bottom.  But first check if Xinerama is enabled.
		// If not, use root geometry.
		var h xinerama.Heads
		if Xutil.ExtInitialized("XINERAMA") {
			h, err = xinerama.PhysicalHeads(Xutil)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			h = xinerama.Heads{rgeom}
		}

		headLock.Lock()
		heads = h
		headLock.Unlock()

		_, err := X.WaitForEvent()
		if err != nil {
			log.Fatal(err)
		}

	}
}

func cursorInBar(barHeight int) bool {
	x, y := robotgo.GetMousePos()
	headLock.RLock()
	h := heads
	headLock.RUnlock()

	i := 1
	for _, head := range h[1:] {
		if x <= head.X() {
			break
		}
		i++
	}
	if y > h[i-1].Height()-barHeight {
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
