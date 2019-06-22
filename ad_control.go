package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const AdControlId componentId = "ad_control"

type adControl struct {
	container *entity
	speed     float64
	moved     bool
}

func newAdControl(container *entity, speed float64) *adControl {
	return &adControl{container: container, speed: speed}
}

func (control *adControl) id() componentId {
	return AdControlId
}

func (control *adControl) update() error {
	keys := sdl.GetKeyboardState()
	position := control.container.position
	leftKeyPressed := keys[sdl.SCANCODE_A] == 1
	rightKeyPressed := keys[sdl.SCANCODE_D] == 1
	nonOrBoth := (!leftKeyPressed && !rightKeyPressed) || (leftKeyPressed && rightKeyPressed)
	otherMoved := control.otherControlMoved()

	if !otherMoved {
		if leftKeyPressed && !rightKeyPressed {
			control.container.position.x = position.x - control.speed*delta
			control.changeMovementType(MovementLeft)
			control.moved = true
		}
		if rightKeyPressed && !leftKeyPressed {
			control.container.position.x = position.x + control.speed*delta
			control.changeMovementType(MovementRight)
			control.moved = true
		}
		if nonOrBoth {
			control.changeMovementType(MovementIdle)
			control.moved = false
		}
	} else {
		control.moved = false
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

func (control *adControl) otherControlMoved() bool {
	lrControl, err := control.container.getComponent(LeftRightControlId)
	if err != nil {
		return false
	}

	return lrControl.(*leftRightControl).moved
}
