package rayframe

import "time"

type RendererScene3D interface {
	Render3D(time.Duration) interface{}
}
