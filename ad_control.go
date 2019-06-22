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
	leftKeyPressed := keys[sdl.SCANCODE_A] == 1
	rightKeyPressed := keys[sdl.SCANCODE_D] == 1
	nonOrBoth := (!leftKeyPressed && !rightKeyPressed) || (leftKeyPressed && rightKeyPressed)

	if leftKeyPressed && !rightKeyPressed {
		control.container.position.x = position.x - control.speed
		control.changeMovementType(MovementLeft)
	}
	if rightKeyPressed && !leftKeyPressed {
		control.container.position.x = position.x + control.speed
		control.changeMovementType(MovementRight)
	}

	if nonOrBoth {
		control.changeMovementType(MovementIdle)
	}

	return nil
}

func (control *adControl) changeMovementType(t movementType) {
	//check if movement could be animated
	comp, err := control.container.getComponent(MovementAnimationRendererId)
	if err == nil {
		comp.(*movementAnimationRenderer).changeMovementType(t)
	}
}
