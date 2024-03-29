package components

import (
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

const AttackId engine.ComponentId = "attack"

type Attack struct {
	container *engine.Entity
	animation *Animation
}

func NewAttack(container *engine.Entity, animation *Animation) *Attack {
	return &Attack{
		container: container,
		animation: animation,
	}
}

func (a *Attack) Id() engine.ComponentId {
	return AttackId
}

func (a *Attack) Update() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_SPACE] == 1 {
		a.onAttack()
	}

	return nil
}

func (a *Attack) onAttack() {
	// do cool attack animation
	comp, _ := a.container.GetComponent(AnimationsId)
	comp.(*Animations).SingleAnimation("attack")
}
