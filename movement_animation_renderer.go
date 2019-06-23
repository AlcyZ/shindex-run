package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const MovementAnimationRendererId componentId = "movement_animation_renderer"

type movementAnimationRenderer struct {
	container       *entity
	renderer        *sdl.Renderer
	movement        *movementAnimations
	currentMovement movementType
}

func newMovementAnimationRenderer(container *entity, renderer *sdl.Renderer, initType movementType) (*movementAnimationRenderer, error) {
	moveAnimations, err := container.getComponent(MovementAnimationsId)
	if err != nil {
		return &movementAnimationRenderer{}, fmt.Errorf("component movement_animation_renderer depends on other component: \n%v", err)
	}

	return &movementAnimationRenderer{
		container:       container,
		renderer:        renderer,
		movement:        moveAnimations.(*movementAnimations),
		currentMovement: initType,
	}, nil
}

func (r *movementAnimationRenderer) id() componentId {
	return MovementAnimationRendererId
}

func (r *movementAnimationRenderer) update() error {
	t := r.currentMovement
	layout := r.movement.layout(t)
	flip := r.movement.flips[t]

	dest := &sdl.Rect{
		X: int32(r.container.position.x),
		Y: int32(r.container.position.y),
		W: layout.width,
		H: layout.height,
	}

	r.movement.updateAnimation(t)

	if err := r.renderer.CopyEx(layout.texture, nil, dest, 0, nil, flip); err != nil {
		return fmt.Errorf("could not render texture to window: \n%v", err)
	}

	return nil
}

func (r *movementAnimationRenderer) changeMovementType(t movementType) {
	if _, ok := r.movement.animations[t]; ok {
		r.currentMovement = t
	}
}
