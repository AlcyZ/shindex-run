package engine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Vector struct {
	X, Y float64
}

func NewVector(x float64, y float64) *Vector {
	return &Vector{X: x, Y: y}
}

type Layout struct {
	Texture *sdl.Texture
	Width   int32
	Height  int32
}

func NewLayout(texture *sdl.Texture, width int32, height int32) *Layout {
	return &Layout{
		Texture: texture,
		Width:   width,
		Height:  height,
	}
}

type Entity struct {
	game       *Game
	position   *Vector
	layout     *Layout
	flip       sdl.RendererFlip
	components map[ComponentId]Component
}

func NewEntity(game *Game) *Entity {
	var components = make(map[ComponentId]Component)

	return &Entity{
		game:       game,
		components: components,
		flip:       sdl.FLIP_NONE,
	}
}

func (e *Entity) AddComponent(c Component) {
	comp := c
	e.components[comp.Id()] = c
}

func (e *Entity) GetComponent(id ComponentId) (Component, error) {
	if val, ok := e.components[id]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("could not find Component: \n%v, [%v]", id, e.components)
}

func (e *Entity) Update() error {
	for _, component := range e.components {
		comp := component
		if err := comp.Update(); err != nil {
			return fmt.Errorf("could not update Component %v: \n%v", comp.Id(), err)
		}
	}
	return nil
}

func (e *Entity) ChangePosition(position *Vector) {
	e.position = position
}

func (e *Entity) ChangeLayout(layout *Layout) {
	e.layout = layout
}

func (e *Entity) ChangeFlip(flip sdl.RendererFlip) {
	e.flip = flip
}

func (e *Entity) CurrentPosition() *Vector {
	return e.position
}

func (e *Entity) CurrentLayout() *Layout {
	return e.layout
}

func (e *Entity) CurrentFlip() sdl.RendererFlip {
	return e.flip
}

func (e *Entity) CanRendered() bool {
	return e.position != &Vector{} && e.layout != nil
}

func (e *Entity) GetDelta() float64 {
	return e.game.Delta()
}
