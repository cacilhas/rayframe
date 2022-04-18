package rayframe

import (
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type RendererScene3D interface {
	Render3D(time.Duration) interface{}
}

func renderScene3D(scene interface{}, rf *RayFrame, dt time.Duration) interface{} {
	res := scene
	if rf.Camera == nil {
		return res
	}
	if sc, ok := scene.(RendererScene3D); ok {
		raylib.BeginMode3D(*rf.Camera)
		res = initialiseScene(scene, sc.Render3D(dt), rf)
		raylib.EndMode3D()
	}
	return res
}
