package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const LeftRightControlId componentId = "left_right_control"

type leftRightControl struct {
	container *entity
	speed     float64
	moved     bool
}

func newLeftRightControl(container *entity, speed float64) *leftRightControl {
	return &leftRightControl{container: container, speed: speed}
}

func (control *leftRightControl) id() componentId {
	return LeftRightControlId
}

func (control *leftRightControl) update() error {
	keys := sdl.GetKeyboardState()
	position := control.container.position
	leftKeyPressed := keys[sdl.SCANCODE_LEFT] == 1
	rightKeyPressed := keys[sdl.SCANCODE_RIGHT] == 1
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

func (control *leftRightControl) changeMovementType(t movementType) {
	//check if movement could be animated
	comp, err := control.container.getComponent(MovementAnimationRendererId)
	if err == nil {
		comp.(*movementAnimationRenderer).changeMovementType(t)
	}
}

func (control *leftRightControl) otherControlMoved() bool {
	adC, err := control.container.getComponent(AdControlId)
	if err != nil {
		return false
	}

	return adC.(*adControl).moved
}
