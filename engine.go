package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"time"
)

type engine struct {
	resX, resY          int32
	posX, posY          int32
	fullscreen, running bool
	window              *sdl.Window
	renderer            *sdl.Renderer

	// to put into assets struct
	font *ttf.Font
	dbgColor sdl.Color
}

func createEngine(cfgPara config) engine {
	var game engine
	game.resX = cfgPara.width
	game.resY = cfgPara.height
	game.fullscreen = cfgPara.fullscreen
	game.running = true
	game.posX = 0
	game.posY = 0
	game.window = nil

	return game
}

func (game *engine) init() {
	var err error

	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	game.window, err = sdl.CreateWindow("test", game.posX, game.posY, game.resX, game.resY, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	if game.fullscreen {
		game.window.SetBordered(false)
		game.window.SetFullscreen(sdl.WINDOW_SHOWN | sdl.WINDOW_FULLSCREEN | sdl.WINDOW_SHOWN | sdl.WINDOW_VULKAN | sdl.WINDOW_OPENGL)
	} else {
		game.window.SetFullscreen(sdl.WINDOW_SHOWN | sdl.WINDOW_VULKAN | sdl.WINDOW_OPENGL)
	}

	game.renderer, err = sdl.CreateRenderer(game.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	sdl.ShowCursor(0)

	// to put into assets struct
	game.font, err = ttf.OpenFont("assets/times.ttf", 24)
	if err != nil {
		panic(err)
	}
}

func (game *engine) run() {
	for game.running {
		game.handleEvents()
		game.render()

		time.Sleep(10 * time.Millisecond)
	}
}

func (game *engine) render() {
	game.renderer.SetDrawColor(0, 0, 0, 0)
	game.renderer.Clear()
	game.renderer.SetDrawColor(40, 40, 40, 255)
	game.renderer.FillRect(&sdl.Rect{0, 0, int32(100), int32(100)})




	// ------------------------------------------------------

	var err error
	var shadedSurface *sdl.Surface
	shadedSurface, err = game.font.RenderUTF8Shaded("Shaded Text", sdl.Color{0, 255, 0, 255}, sdl.Color{255, 0, 255, 255})
	if err != nil {
		panic(err)
	}

	var shadedTexture  *sdl.Texture
	shadedTexture, err = game.renderer.CreateTextureFromSurface(shadedSurface)
	if err != nil {
		panic(err)
	}

	shadedSurface.Free()
	game.renderer.Copy(shadedTexture, nil, &sdl.Rect{300, 20, 244, 53})

	// ------------------------------------------------------




	game.renderer.Present()
}

func (game *engine) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			game.running = false
			break
		case *sdl.KeyboardEvent:
			game.handleKeyboard(e)
			break
		}
	}
}

func (game *engine) handleKeyboard(e *sdl.KeyboardEvent) {
	if e.Type != sdl.KEYUP {
		return
	}

	switch e.Keysym.Sym {
	case sdl.K_ESCAPE:
		game.running = false
	}
}
