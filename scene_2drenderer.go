package rayframe

type RendererScene2D interface {
	Render2D() interface{}
}

func renderScene2D(scene interface{}, frame *RayFrame) interface{} {
	res := scene
	if sc, ok := scene.(RendererScene2D); ok {
		res = initialiseScene(scene, sc.Render2D(), frame)
	}
	return res
}
