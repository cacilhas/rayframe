package rayframe

type InitScene interface {
	Init(*RayFrame)
}

func initialiseScene(before, after interface{}, frame *RayFrame) interface{} {
	if after == nil {
		return before
	}
	if before == after {
		return after
	}
	if scene, ok := after.(InitScene); ok {
		scene.Init(frame)
	}
	return after
}
