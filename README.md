Hotbar for i3wm
===============

Do you use [i3wm](https://i3wm.org) with the bar set on [hide
mode](https://i3wm.org/docs/userguide.html#_display_mode)? This mode lets you
save the maximum screen space by hiding the bar unless the modifier key is
pressed. By preventing status updates with the bar hidden, it also helps you
save some battery.

All that is vanilla i3. What this hotbar process does is to unhide/hide the bar
when your mouse cursor is over it, without the need to press modifier key. I
know i3 is all about avoiding mouse usage, but sometimes I use the mouse and
this way is comfortable to use. I hope you find it comfortable as well.

Dependencies:

This depends on `xdolib`. It usually comes bundled with `xdotool`, in your
distribution packages. Install that first.

To install, after the dependencies, do:
```bash
$ go get github.com/antoniomo/i3-hotbar
```

To execute at i3 startup, put this in your `.i3/config`:

```
# Hotbar script
exec --no-startup-id $GOPATH/bin/i3-hotbar
```

Have fun!
