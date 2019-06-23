package main

import "github.com/veandco/go-sdl2/sdl"

type movementType = string

const (
	MovementIdle  movementType = "Idle"
	MovementLeft  movementType = "left"
	MovementRight movementType = "right"
)

const MovementAnimationsId componentId = "movement_animations"

type movementAnimations struct {
	animations map[movementType]*animation
	flips      map[movementType]sdl.RendererFlip
}

func newMovementAnimations() *movementAnimations {
	animations := make(map[movementType]*animation)
	flips := make(map[movementType]sdl.RendererFlip)

	return &movementAnimations{animations: animations, flips: flips}
}

func (movement *movementAnimations) add(animation *animation, flip sdl.RendererFlip, t movementType) {
	movement.animations[t] = animation
	movement.flips[t] = flip
}

func (movement *movementAnimations) layout(t movementType) *layout {
	return movement.animations[t].layout()
}

func (movement *movementAnimations) updateAnimation(t movementType) {
	_ = movement.animations[t].update()
}

func (movement *movementAnimations) id() componentId {
	return MovementAnimationsId
}

func (movement *movementAnimations) update() error {
	return nil
}
