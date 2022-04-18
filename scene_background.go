package rayframe

import (
	"image/color"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type BackgroundScene interface {
	Background() color.RGBA
}

func drawBackground(scene interface{}) {
	if sc, ok := scene.(BackgroundScene); ok {
		raylib.ClearBackground(sc.Background())
	}
}
