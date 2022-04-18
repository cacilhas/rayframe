package rayframe

import "time"

type RendererScene2D interface {
	Render2D(dt time.Duration) interface{}
}

func renderScene2D(scene interface{}, rf *RayFrame, dt time.Duration) interface{} {
	res := scene
	if sc, ok := scene.(RendererScene2D); ok {
		res = initialiseScene(scene, sc.Render2D(dt), rf)
	}
	return res
}
