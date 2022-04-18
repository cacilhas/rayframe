package rayframe

import "time"

type RendererScene2D interface {
	Render2D(dt time.Duration) interface{}
}
