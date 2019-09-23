// Package screenshot captures screen-shot image as image.RGBA.
// Mac, Windows, Linux, FreeBSD, OpenBSD, NetBSD, and Solaris are supported.
package screenshot

import (
	"image"
	"github.com/BurntSushi/xgb"
)

// CaptureDisplay captures whole region of displayIndex'th display.
func CaptureDisplay(displayIndex int) (*image.RGBA, error) {
	rect := GetDisplayBounds(displayIndex)
	c, err := xgb.NewConn()
	if err != nil {
		panic(err)
	}
	return CaptureRect(rect, c)
}

// CaptureRect captures specified region of desktop.
func CaptureRect(rect image.Rectangle, c *xgb.Conn) (*image.RGBA, error) {
	return Capture(rect.Min.X, rect.Min.Y, rect.Dx(), rect.Dy(), c)
}
