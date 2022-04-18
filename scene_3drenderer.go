package rayframe

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type RendererScene3D interface {
	Render3D() Scene
}

func renderScene3D(scene Scene, frame *RayFrame) Scene {
	res := scene
	if frame.Camera == nil {
		return res
	}
	if sc, ok := scene.(RendererScene3D); ok {
		raylib.BeginMode3D(*frame.Camera)
		res = initialiseScene(scene, sc.Render3D(), frame)
		raylib.EndMode3D()
	}
	return res
}
