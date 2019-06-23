package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

const (
	screenWidth  = 1200
	screenHeight = 800

	deltaTicks = 60
	maxFps     = 240
)

var delta float64
var entities []*entity

//var sys [] systems.System

func boot() (*sdl.Window, *sdl.Renderer, error) {
	// init sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return &sdl.Window{}, &sdl.Renderer{}, fmt.Errorf("could not init SDL: \n%v", err)
	}

	// create window
	w, err := sdl.CreateWindow("Runner", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		return &sdl.Window{}, &sdl.Renderer{}, fmt.Errorf("could not create window: \n%v", err)
	}

	// create renderer
	r, err := sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return &sdl.Window{}, &sdl.Renderer{}, fmt.Errorf("could not create r: \n%v", err)
	}

	return w, r, nil
}

func main() {
	w, r, err := boot()
	if err != nil {
		fmt.Printf("could not boot: \n%v", err)
		return
	}
	defer sdl.Quit()
	defer w.Destroy()
	defer r.Destroy()

	player, err := newPlayer(r, 16, "assets/ninja/Idle__000.png")
	if err != nil {
		fmt.Println("player init failed: \n", err)
		return
	}

	enemy, err := newEnemy(r, "assets/player/male/Idle_0.png")
	if err != nil {
		fmt.Println("enemy init failed: \n", err)
		return
	}

	background, err := newBackground(r, "assets/background.jpg")
	if err != nil {
		fmt.Println("bg init failed: \n", err)
		return
	}

	entities = append(entities, background)
	entities = append(entities, enemy)
	entities = append(entities, player)

	for {
		frameRenderBegin := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("End game!")
				return
			}
		}

		_ = r.SetDrawColor(255, 255, 255, 0)
		_ = r.Clear()

		for _, entity := range entities {
			if err := entity.update(); err != nil {
				fmt.Println("could not update entity: ", err)
			}
		}

		r.Present()
		time.Sleep(time.Second / maxFps) // some hack to reduce resource usage
		delta = time.Since(frameRenderBegin).Seconds() * deltaTicks
	}
}
