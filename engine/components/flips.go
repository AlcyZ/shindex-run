package components

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

const FlipsId engine.ComponentId = "flips"

type FlipType string

type Flips struct {
	container *engine.Entity
	flips     map[FlipType]sdl.RendererFlip
	active    FlipType
}

func NewFlips(container *engine.Entity) *Flips {
	flips := make(map[FlipType]sdl.RendererFlip)

	return &Flips{
		container: container,
		flips:     flips,
	}
}

func (f *Flips) Id() engine.ComponentId {
	return FlipsId
}

func (f *Flips) Update() error {
	f.container.ChangeFlip(f.Active())
	return nil
}

func (f *Flips) Add(flip sdl.RendererFlip, t FlipType) {
	f.flips[t] = flip
	f.active = t
}

func (f *Flips) Switch(t FlipType) error {
	if _, ok := f.flips[t]; ok {
		f.active = t

		return nil
	}

	return fmt.Errorf("could not switch flip, fliptype is not added to list: (%v)\n", t)
}

func (f *Flips) Active() sdl.RendererFlip {
	return f.flips[f.active]
}
