package rayframe

type InitScene interface {
	Init(*RayFrame) interface{}
}

func initialiseScene(before, after interface{}, frame *RayFrame) interface{} {
	if before == after {
		return after
	}
	if scene, ok := after.(InitScene); ok {
		return scene.Init(frame)
	}
	return after
}
