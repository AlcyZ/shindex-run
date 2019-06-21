package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type adControl struct {
	container *entity
}

func newAdControl(container *entity) *adControl {
	return &adControl{container: container}
}

func (control *adControl) id() componentId {
	return "ad_control"
}

func (control *adControl) update() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_A] == 1 {
		fmt.Println("A key pressed: ", time.Now())
	}
	if keys[sdl.SCANCODE_D] == 1 {
		fmt.Println("D key pressed: ", time.Now())
	}

	return nil
}
