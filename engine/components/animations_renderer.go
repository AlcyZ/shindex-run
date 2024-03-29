package components

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
)

const AnimationsRendererId = "animations_renderer"

type animationsRenderer struct {
	container *engine.Entity
	r         *sdl.Renderer
	a         *Animations
}

func NewAnimationsRenderer(container *engine.Entity, r *sdl.Renderer) (*animationsRenderer, error) {
	a, err := container.GetComponent(AnimationsId)
	if err != nil {
		return &animationsRenderer{}, fmt.Errorf("could not get component:\n%v", err)
	}

	return &animationsRenderer{
		container: container,
		r:         r,
		a:         a.(*Animations),
	}, nil
}

func (r *animationsRenderer) Id() engine.ComponentId {
	return AnimationsRendererId
}

func (r *animationsRenderer) Update() error {
	if r.container.CanRendered() {
		position := r.container.CurrentPosition()
		layout := r.container.CurrentLayout()

		dest := &sdl.Rect{
			X: int32(position.X),
			Y: int32(position.Y),
			W: layout.Width,
			H: layout.Height,
		}

		if err := r.r.CopyEx(layout.Texture, nil, dest, 0, nil, r.container.CurrentFlip()); err != nil {
			return fmt.Errorf("could not render texture to window: \n%v", err)
		}
	}
	return nil
}
