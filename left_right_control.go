package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type leftRightControl struct {
	container *entity
	speed     float64
}

func newLeftRightControl(container *entity, speed float64) *leftRightControl {
	return &leftRightControl{container: container, speed: speed}
}

func (control *leftRightControl) id() componentId {
	return "left_right_control"
}

func (control *leftRightControl) update() error {
	keys := sdl.GetKeyboardState()
	position := control.container.position

	if keys[sdl.SCANCODE_LEFT] == 1 {
		control.container.position.x = position.x - control.speed
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		control.container.position.x = position.x + control.speed
	}

	return nil
}
