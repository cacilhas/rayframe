package rayframe

import (
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type RayFrame struct {
	camera     *raylib.Camera
	fps        int
	tick       time.Time
	windowSize raylib.Vector2
}

func (rf *RayFrame) Mainloop(initialScene interface{}) {
	rf.resize()
	if rf.fps > 0 {
		raylib.SetTargetFPS(int32(rf.fps))
	} else {
		raylib.SetTargetFPS(30)
	}
	scene := rf.initialiseScene(initialScene)
	rf.tick = time.Now()
	for !raylib.WindowShouldClose() {
		scene = rf.tic(scene)
		if rf.fps > 0 {
			time.Sleep(time.Second/time.Duration(rf.fps) - time.Since(rf.tick))
		}
	}
}

func (rf *RayFrame) tic(scene interface{}) interface{} {
	dt := time.Since(rf.tick)
	rf.tick = time.Now()
	if raylib.IsWindowResized() {
		rf.resize()
	}

	raylib.BeginDrawing()

	if sc, ok := scene.(BackgroundScene); ok {
		raylib.ClearBackground(sc.Background())
	}
	if sc, ok := scene.(RendererScene3D); ok && rf.camera != nil {
		raylib.BeginMode3D(*rf.camera)
		newScene := sc.Render3D(dt)
		raylib.EndMode3D()
		if newScene != scene {
			scene = rf.initialiseScene(newScene)
		}
	}
	if sc, ok := scene.(RendererScene2D); ok {
		newScene := sc.Render2D(dt)
		if newScene != scene {
			scene = rf.initialiseScene(newScene)
		}
	}

	raylib.EndDrawing()
	return scene
}

func (rf *RayFrame) resize() {
	if raylib.IsWindowFullscreen() {
		currentMonitor := raylib.GetCurrentMonitor()
		rf.windowSize = raylib.NewVector2(
			float32(raylib.GetMonitorWidth(currentMonitor)),
			float32(raylib.GetMonitorHeight(currentMonitor)),
		)
	} else {
		rf.windowSize = raylib.NewVector2(
			float32(raylib.GetScreenWidth()),
			float32(raylib.GetScreenHeight()),
		)
	}
}

func (rf *RayFrame) initialiseScene(scene interface{}) interface{} {
	if sc, ok := scene.(InitScene); ok {
		return sc.Init(rf)
	}
	return scene
}
