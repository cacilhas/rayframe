package rayframe

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type RendererScene3D interface {
	Render3D() interface{}
}

func renderScene3D(scene interface{}, rf *RayFrame) interface{} {
	res := scene
	if rf.Camera == nil {
		return res
	}
	if sc, ok := scene.(RendererScene3D); ok {
		raylib.BeginMode3D(*rf.Camera)
		res = initialiseScene(scene, sc.Render3D(), rf)
		raylib.EndMode3D()
	}
	return res
}
