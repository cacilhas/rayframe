package rayframe

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type ExitKeyScene interface {
	ExitKey() int32
	OnKeyEscape() Scene
}

func setExitKey(scene Scene) {
	if sc, ok := scene.(ExitKeyScene); ok {
		raylib.SetExitKey(sc.ExitKey())
	} else {
		raylib.SetExitKey(raylib.KeyEscape)
	}
}

func onKeyEscape(scene Scene, frame *RayFrame) Scene {
	if sc, ok := scene.(ExitKeyScene); ok {
		if raylib.IsKeyPressed(raylib.KeyEscape) {
			return initialiseScene(scene, sc.OnKeyEscape(), frame)
		}
	}
	return scene
}
