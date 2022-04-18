package rayframe

import "time"

type UpdateScene interface {
	Update(time.Duration) interface{}
}

func updateScene(scene interface{}, frame *RayFrame, dt time.Duration) interface{} {
	res := scene
	if sc, ok := scene.(UpdateScene); ok {
		res = initialiseScene(scene, sc.Update(dt), frame)
	}
	return res
}
