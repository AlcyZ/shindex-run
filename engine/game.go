package engine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type Game struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	entities []*Entity
	delta    float64
}

func CreateSimpleGame(window *sdl.Window, renderer *sdl.Renderer) *Game {
	return &Game{
		window:   window,
		renderer: renderer,
	}
}

func (g *Game) Start() {
	for {
		frameRenderBegin := time.Now()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("End game!")
				return
			}
		}

		_ = g.renderer.SetDrawColor(255, 255, 255, 0)
		_ = g.renderer.Clear()

		for _, entity := range g.entities {
			if err := entity.Update(); err != nil {
				fmt.Println("could not update entity: ", err)
			}
		}

		g.renderer.Present()
		time.Sleep(time.Second / MaxFps) // some hack to reduce resource usage
		g.delta = time.Since(frameRenderBegin).Seconds() * DeltaTicks
	}
}

func (g *Game) Delta() float64 {
	return g.delta
}

func (g *Game) AddEntity(e *Entity) {
	g.entities = append(g.entities, e)
}
