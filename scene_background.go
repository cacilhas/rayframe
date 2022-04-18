package rayframe

import "image/color"

type BackgroundScene interface {
	Background() color.RGBA
}
