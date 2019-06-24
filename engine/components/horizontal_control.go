package components

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

const (
	HorizontalControlId = "horizontal_control_id"
	mappingKeyIdle      = "idle"
	mappingKeyLeft      = "left"
	mappingKeyRight     = "right"
)

type HorizontalControl struct {
	container        *engine.Entity
	speed            float64
	leftKeys         []sdl.Scancode
	rightKeys        []sdl.Scancode
	animated         bool
	animationMapping map[string]AnimationType
}

func NewHorizontalControl(container *engine.Entity, speed float64, leftKeys []sdl.Scancode, rightKeys []sdl.Scancode) *HorizontalControl {
	return &HorizontalControl{
		container: container,
		speed:     speed,
		leftKeys:  leftKeys,
		rightKeys: rightKeys,
		animated:  false,
	}
}

func (c *HorizontalControl) WithAnimations(idle AnimationType, left AnimationType, right AnimationType) error {
	_, err := c.container.GetComponent(AnimationsId)
	if err != nil {
		return fmt.Errorf("animations not available on container entity: %v", err)
	}
	mapping := make(map[string]AnimationType)
	c.animationMapping = mapping

	c.animationMapping[mappingKeyIdle] = idle
	c.animationMapping[mappingKeyLeft] = left
	c.animationMapping[mappingKeyRight] = right
	c.animated = true

	return nil
}

func (c *HorizontalControl) Id() engine.ComponentId {
	return HorizontalControlId
}

func (c *HorizontalControl) Update() error {
	keys := sdl.GetKeyboardState()
	position := c.container.CurrentPosition()
	delta := c.container.GetDelta()
	leftKeyPressed := false
	rightKeyPressed := false
	nonOrBoth := false

	for _, lScanCode := range c.leftKeys {
		if !leftKeyPressed && keys[lScanCode] == 1 {
			c.container.ChangePosition(engine.NewVector(position.X-c.speed*delta, position.Y))
			leftKeyPressed = true
			c.changeAnimation(mappingKeyLeft)
		}
	}
	for _, rScanCode := range c.rightKeys {
		if !rightKeyPressed && keys[rScanCode] == 1 {
			c.container.ChangePosition(engine.NewVector(position.X+c.speed*delta, position.Y))
			rightKeyPressed = true
			c.changeAnimation(mappingKeyRight)
		}
	}

	nonOrBoth = (leftKeyPressed && rightKeyPressed) || (!leftKeyPressed && !rightKeyPressed)
	if nonOrBoth {
		// set idle Animation
		c.changeAnimation(mappingKeyIdle)
	}

	return nil
}

func (c *HorizontalControl) changeAnimation(t string) {
	if c.animated {
		//a if animated is true, the Animations component must be available
		comp, _ := c.container.GetComponent(AnimationsId)
		comp.(*Animations).ChangeAnimation(c.animationMapping[t])
	}
}
