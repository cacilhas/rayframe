package rayframe

import (
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type RayFrame struct {
	Camera     *raylib.Camera
	FPS        int
	Tick       time.Time
	WindowSize raylib.Vector2
}

func (rf *RayFrame) Mainloop(initialScene interface{}) {
	rf.resize()
	if rf.FPS > 0 {
		raylib.SetTargetFPS(int32(rf.FPS))
	} else {
		raylib.SetTargetFPS(30)
	}
	scene := initialiseScene(nil, initialScene, rf)
	rf.Tick = time.Now()
	for !raylib.WindowShouldClose() {
		scene = rf.tic(scene)
		if rf.FPS > 0 {
			time.Sleep(time.Second/time.Duration(rf.FPS) - time.Since(rf.Tick))
		}
	}
}

func (rf *RayFrame) tic(scene interface{}) interface{} {
	dt := time.Since(rf.Tick)
	rf.Tick = time.Now()
	if raylib.IsWindowResized() {
		rf.resize()
	}

	raylib.BeginDrawing()
	drawBackground(scene)
	scene = renderScene3D(scene, rf, dt)
	scene = renderScene2D(scene, rf, dt)
	raylib.EndDrawing()
	return scene
}

func (rf *RayFrame) resize() {
	if raylib.IsWindowFullscreen() {
		currentMonitor := raylib.GetCurrentMonitor()
		rf.WindowSize = raylib.NewVector2(
			float32(raylib.GetMonitorWidth(currentMonitor)),
			float32(raylib.GetMonitorHeight(currentMonitor)),
		)
	} else {
		rf.WindowSize = raylib.NewVector2(
			float32(raylib.GetScreenWidth()),
			float32(raylib.GetScreenHeight()),
		)
	}
}
