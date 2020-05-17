package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type engine struct {
	resX, resY int32
	posX, posY int32
	fullscreen, running bool
	window *sdl.Window
	renderer *sdl.Renderer
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


	TTF_Font* Sans = TTF_OpenFont("Sans.ttf", 24); //this opens a font style and sets a size

	SDL_Color White = {255, 255, 255};  // this is the color in rgb format, maxing out all would give you the color white, and it will be your text's color

	SDL_Surface* surfaceMessage = TTF_RenderText_Solid(Sans, "put your text here", White); // as TTF_RenderText_Solid could only be used on SDL_Surface then you have to create the surface first

	SDL_Texture* Message = SDL_CreateTextureFromSurface(renderer, surfaceMessage); //now you can convert it into a texture


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
	if(e.Type != sdl.KEYUP) {
		return
	}

	switch e.Keysym.Sym {
	case sdl.K_ESCAPE:
		game.running = false
	}
}