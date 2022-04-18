package rayframe

type RendererScene2D interface {
	Render2D() Scene
}

func renderScene2D(scene Scene, frame *RayFrame) Scene {
	res := scene
	if sc, ok := scene.(RendererScene2D); ok {
		res = initialiseScene(scene, sc.Render2D(), frame)
	}
	return res
}
