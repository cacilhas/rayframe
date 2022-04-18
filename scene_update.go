package rayframe

import "time"

type UpdateScene interface {
	Update(dt time.Duration) interface{}
}

func updateScene(scene interface{}, dt time.Duration) interface{} {
	res := scene
	if sc, ok := scene.(UpdateScene); ok {
		res = sc.Update(dt)
	}
	return res
}
