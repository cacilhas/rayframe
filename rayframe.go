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
	OnResize   func(int, int, bool)
	Tick       time.Time
	WindowSize intVector2D
}

func (frame *RayFrame) Init(width, height int, title string) {
	frame.WindowSize = intVector2D{X: width, Y: height}
	raylib.InitWindow(int32(width), int32(height), title)
	if frame.FPS > 0 {
		raylib.SetTargetFPS(int32(frame.FPS))
	} else {
		raylib.SetTargetFPS(30)
	}
}

func (frame *RayFrame) Mainloop(initialScene interface{}) {
	frame.resize()
	scene := initialiseScene(nil, initialScene, frame)
	frame.Tick = time.Now()
	for !raylib.WindowShouldClose() {
		scene = frame.tic(scene)
		if frame.FPS > 0 {
			time.Sleep(time.Second/time.Duration(frame.FPS) - time.Since(frame.Tick))
		}
	}
	raylib.CloseWindow()
}

func (frame *RayFrame) tic(scene interface{}) interface{} {
	dt := time.Since(frame.Tick)
	frame.Tick = time.Now()
	if raylib.IsWindowResized() {
		frame.resize()
	}

	raylib.BeginDrawing()
	drawBackground(scene)
	scene = updateScene(scene, frame, dt)
	scene = renderScene3D(scene, frame)
	scene = renderScene2D(scene, frame)
	raylib.EndDrawing()
	return scene
}

func (frame *RayFrame) resize() {
	fullscreen := raylib.IsWindowFullscreen()
	if fullscreen {
		currentMonitor := raylib.GetCurrentMonitor()
		frame.WindowSize = intVector2D{
			X: raylib.GetMonitorWidth(currentMonitor),
			Y: raylib.GetMonitorHeight(currentMonitor),
		}
	} else {
		frame.WindowSize = intVector2D{
			X: raylib.GetScreenWidth(),
			Y: raylib.GetScreenHeight(),
		}
	}
	if frame.OnResize != nil {
		frame.OnResize(frame.WindowSize.X, frame.WindowSize.Y, fullscreen)
	}
}
