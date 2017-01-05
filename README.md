Hotbar for i3wm
===============

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

## Installation

Two options:

If you aren't a `go` developer or don't want to fiddle with the `C`
dependencies, just clone this somewhere and add this to your `~/.i3/config`:

```
# Hotbar script
exec --no-startup-id WHATEVER_PATH/i3-hotbar/bin/i3-hotbar
```

If you prefer a "source" install:

Dependencies:

This depends on `xdolib`. It usually comes bundled with `xdotool`, in your
distribution packages. Install that first.

After installing the dependencies, do:
```bash
$ go get github.com/antoniomo/i3-hotbar
```

To execute at i3 startup, put this in your `~/.i3/config`:

```
# Hotbar script
exec --no-startup-id $GOPATH/bin/i3-hotbar
```

Have fun!
