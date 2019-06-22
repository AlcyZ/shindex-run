package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type adControl struct {
	container *entity
	speed     float64
}

func newAdControl(container *entity, speed float64) *adControl {
	return &adControl{container: container, speed: speed}
}

func (control *adControl) id() componentId {
	return "ad_control"
}

func (control *adControl) update() error {
	keys := sdl.GetKeyboardState()
	position := control.container.position

	if keys[sdl.SCANCODE_A] == 1 {
		control.container.position.x = position.x - control.speed
	}
	if keys[sdl.SCANCODE_D] == 1 {
		control.container.position.x = position.x + control.speed
	}

	return nil
}
