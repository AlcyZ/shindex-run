package components

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
	"time"
)

const AnimationId engine.ComponentId = "Animation"

type Animation struct {
	container       *engine.Entity
	layouts         []*engine.Layout
	lastChange      time.Time
	changeRate      time.Duration
	animationFrames int
	duration        time.Duration
	currentIndex    int
}

func NewAnimation(container *engine.Entity, textures []*sdl.Texture, duration time.Duration, scaling float64, flip sdl.RendererFlip) (*Animation, error) {
	var layouts []*engine.Layout
	frames := len(textures)

	for _, texture := range textures {
		_, _, width, height, err := texture.Query()
		if err != nil {
			return &Animation{}, fmt.Errorf("could not query widht and height from texture: \n%v", err)
		}

		layout := engine.NewLayout(
			texture,
			int32(float64(width)*scaling),
			int32(float64(height)*scaling),
			flip,
		)
		layouts = append(layouts, layout)
	}

	return &Animation{
		container:       container,
		layouts:         layouts,
		duration:        duration,
		animationFrames: frames,
		changeRate:      duration / time.Duration(frames),
	}, nil
}

func (a *Animation) Id() engine.ComponentId {
	return AnimationId
}

func (a *Animation) Update() error {
	a.checkIndex()
	a.container.ChangeLayout(a.Layout())

	return nil
}

func (a *Animation) Layout() *engine.Layout {
	return a.layouts[a.currentIndex]
}

func (a *Animation) checkIndex() {
	if time.Since(a.lastChange) > time.Since(time.Now())+a.changeRate {
		if a.currentIndex >= a.animationFrames-1 {
			a.currentIndex = 0
		} else {
			a.currentIndex++
		}
		a.lastChange = time.Now()
	}
}
