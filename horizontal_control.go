package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	HorizontalControlId = "horizontal_control_id"
	mappingKeyIdle      = "idle"
	mappingKeyLeft      = "left"
	mappingKeyRight     = "right"
)

type horizontalControl struct {
	container        *entity
	speed            float64
	leftKeys         []sdl.Scancode
	rightKeys        []sdl.Scancode
	animated         bool
	animationMapping map[string]animationType
}

func newHorizontalControl(container *entity, speed float64, leftKeys []sdl.Scancode, rightKeys []sdl.Scancode) *horizontalControl {
	return &horizontalControl{
		container: container,
		speed:     speed,
		leftKeys:  leftKeys,
		rightKeys: rightKeys,
		animated:  false,
	}
}

func (c *horizontalControl) withAnimations(idle animationType, left animationType, right animationType) error {
	_, err := c.container.getComponent(AnimationsId)
	if err != nil {
		return fmt.Errorf("animations not available on container entity: %v", err)
	}
	mapping := make(map[string]animationType)
	c.animationMapping = mapping

	c.animationMapping[mappingKeyIdle] = idle
	c.animationMapping[mappingKeyLeft] = left
	c.animationMapping[mappingKeyRight] = right
	c.animated = true

	return nil
}

func (c *horizontalControl) id() componentId {
	return HorizontalControlId
}

func (c *horizontalControl) update() error {
	keys := sdl.GetKeyboardState()
	position := c.container.position
	leftKeyPressed := false
	rightKeyPressed := false
	nonOrBoth := false

	for _, lScanCode := range c.leftKeys {
		if !leftKeyPressed && keys[lScanCode] == 1 {
			c.container.position.x = position.x - c.speed*delta
			leftKeyPressed = true
			c.changeAnimation(mappingKeyLeft)
		}
	}
	for _, rScanCode := range c.rightKeys {
		if !rightKeyPressed && keys[rScanCode] == 1 {
			c.container.position.x = position.x + c.speed*delta
			rightKeyPressed = true
			c.changeAnimation(mappingKeyRight)
		}
	}

	nonOrBoth = (leftKeyPressed && rightKeyPressed) || (!leftKeyPressed && !rightKeyPressed)
	if nonOrBoth {
		// set idle animation
		c.changeAnimation(mappingKeyIdle)
	}

	return nil
}

func (c *horizontalControl) changeAnimation(t string) {
	if c.animated {
		//a if animated is true, the animations component must be available
		comp, _ := c.container.getComponent(AnimationsId)
		comp.(*animations).changeAnimation(c.animationMapping[t])
	}
}
