package rayframe

type InitScene interface {
	Init(*RayFrame) interface{}
}

func initialiseScene(before, after interface{}, rf *RayFrame) interface{} {
	if before == after {
		return after
	}
	if scene, ok := after.(InitScene); ok {
		return scene.Init(rf)
	}
	return after
}
