package rayframe

import "time"

type UpdateScene interface {
	Update(time.Duration) Scene
}

func updateScene(scene Scene, frame *RayFrame, dt time.Duration) Scene {
	if sc, ok := scene.(UpdateScene); ok {
		return initialiseScene(scene, sc.Update(dt), frame)
	}
	return scene
}
