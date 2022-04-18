package rayframe

type InitScene interface {
	Init(*RayFrame)
}

// If called with (nil, scene, frame), Init() method is called on the scene, and
// the scene is returned.
// If called with (scene, nil, frame), Init() method isn’t called, and the
// scene is returned.
// If called with (scene, scene, frame), Init() method isn’t called, and the
// scene is returned.
// If called with (sceneA, sceneB, frame), Init() method is called on the
// sceneB, and the sceneB is returned.
func initialiseScene(before, after Scene, frame *RayFrame) Scene {
	if after == nil {
		return before
	}
	if before == after {
		return after
	}
	if scene, ok := after.(InitScene); ok {
		scene.Init(frame)
		setExitKey(scene)
	}
	return after
}
