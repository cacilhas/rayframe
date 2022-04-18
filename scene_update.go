package rayframe

import "time"

type UpdateScene interface {
	Update(time.Duration) Scene
}

func updateScene(scene Scene, frame *RayFrame, dt time.Duration) Scene {
	res := scene
	if sc, ok := scene.(UpdateScene); ok {
		res = initialiseScene(scene, sc.Update(dt), frame)
	}
	return res
}
