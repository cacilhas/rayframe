package rayframe

import (
	"time"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type intVector2D = struct {
	X, Y int
}

type RayFrame struct {
	Camera     *raylib.Camera
	FPS        int
	OnRezise   func(w, h int)
	Tick       time.Time
	WindowSize intVector2D
}

func (rf *RayFrame) Init(width, height int, title string) {
	rf.WindowSize = intVector2D{X: width, Y: height}
	raylib.InitWindow(int32(width), int32(height), title)
	if rf.FPS > 0 {
		raylib.SetTargetFPS(int32(rf.FPS))
	} else {
		raylib.SetTargetFPS(30)
	}
}

func (rf *RayFrame) Mainloop(initialScene interface{}) {
	rf.resize()
	scene := initialiseScene(nil, initialScene, rf)
	rf.Tick = time.Now()
	for !raylib.WindowShouldClose() {
		scene = rf.tic(scene)
		if rf.FPS > 0 {
			time.Sleep(time.Second/time.Duration(rf.FPS) - time.Since(rf.Tick))
		}
	}
	raylib.CloseWindow()
}

func (rf *RayFrame) tic(scene interface{}) interface{} {
	dt := time.Since(rf.Tick)
	rf.Tick = time.Now()
	if raylib.IsWindowResized() {
		rf.resize()
	}

	raylib.BeginDrawing()
	drawBackground(scene)
	scene = updateScene(scene, dt)
	scene = renderScene3D(scene, rf)
	scene = renderScene2D(scene, rf)
	raylib.EndDrawing()
	return scene
}

func (rf *RayFrame) resize() {
	if raylib.IsWindowFullscreen() {
		currentMonitor := raylib.GetCurrentMonitor()
		rf.WindowSize = intVector2D{
			X: raylib.GetMonitorWidth(currentMonitor),
			Y: raylib.GetMonitorHeight(currentMonitor),
		}
	} else {
		rf.WindowSize = intVector2D{
			X: raylib.GetScreenWidth(),
			Y: raylib.GetScreenHeight(),
		}
	}
	if rf.OnRezise != nil {
		rf.OnRezise(rf.WindowSize.X, rf.WindowSize.Y)
	}
}
