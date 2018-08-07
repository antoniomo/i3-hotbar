# Hotbar for i3wm

Do you use [i3wm](https://i3wm.org) with the bar set on [hide
mode](https://i3wm.org/docs/userguide.html#_display_mode)? This mode lets you
save the maximum screen space by hiding the bar unless the modifier key is
pressed. By preventing status updates with the bar hidden, it also helps you
save some battery.

All that is vanilla `i3`. What this hotbar process does is to unhide/hide the
bar when your mouse cursor is over it, without the need to press the modifier
key. I know `i3` is all about avoiding mouse usage, but to me this makes it more
comfortable to use with the hidden bar. If you are here, I hope you find it
useful as well.

This depends heavily on:

- https://github.com/BurntSushi/xgb
- https://github.com/BurntSushi/xgbutil
- https://github.com/go-vgo/robotgo

Some of the code is taken from their docs/examples, and I didn't strip comments
or anything since, well, they are informative. All credit to their authors.

## Installation

Assuming you have https://github.com/golang/dep installed:

```
go get -u github.com/antoniomo/i3-hotbar
cd $GOPATH/src/github.com/antoniomo/i3-hotbar
dep ensure
go install
```

Then on your ``~/.i3/config`:

```
# Hotbar script
exec --no-startup-id $GOPATH/bin/i3-hotbar
```

Have fun!
