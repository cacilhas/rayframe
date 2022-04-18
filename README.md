[bsd-3-clause]: https://opensource.org/licenses/BSD-3-Clause
[email]: mailto:batalema@cacilhas.info
[raylib]: https://www.raylib.com/
[raylib-go]: https://github.com/gen2brain/raylib-go

# Rayframe

Rayframe is a [`raylib`][raylib] framework for [Go][raylib-go].

## Use

### Creating a framework

Instantiate a `*rayframe.RayFrame`:

```go
frame := &rayframe.RayFrame{}
```

The fields you can set are:

- `Camera`: a `*raylib.Camera` for 3D rendering
- `FPS`: how many frames per second
- `OnResize`: a callback (`func(int, int)`) called whenever the window is resized

The fields you can read:

- `Tick`: last tick time
- `WindowSize`: current window size (`struct { X, Y int }`)

### Starting the framework

Initialise the framework by calling:

```go
frame.Init(1280, 720, "My Application")
```

The parameters are:

1. `width int`: window initial width
1. `height int`: window initial height
1. `title string`: window title

Then start the main loop by calling:

```go
frame.Mainloop(scene)
```

`scene` can be any structure point.

### Scene

Each scene may implement any of the following methods:

- `Init(*rayframe.RayFrame)`: called when the framework shifts to
  the scene.
- `Background() color.RBGA`: inform the framework which colour to use when
  painting the background. If not implemented, the framework won’t call
  `raylib.ClearBackground()`.
- `Update(time.Duration) interface{}`: called each tick and receives the
  timedelta since last tick.
- `Render2D() interface{}`: used to render 2D assets.
- `Render3D() interface{}`: used to render 3D assets under 3D mode, **if** the
  framework’s `Camera` isn’t `nil`.

If any `interface{}`-return method returns anything but the calling scene
itself, the framework will change the scene by the return value and will call
`Init()` (if it’s implemented).

## License

- Copyright ©2022 [Arĥimedeς Montegasppα Cacilhας][email]
- [3-Clause BSD License][bsd-3-clause]
- [COPYING](/cacilhas/rayframe/blob/master/COPYING)