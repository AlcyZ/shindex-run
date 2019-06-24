package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
	"time"
)

const AnimationId engine.ComponentId = "animation"

type layout struct {
	texture *sdl.Texture
	width   int32
	height  int32
}

type animation struct {
	container       *engine.Entity
	layouts         []*layout
	lastChange      time.Time
	changeRate      time.Duration
	animationFrames int
	duration        time.Duration
	currentIndex    int
}

func newAnimation(container *engine.Entity, textures []*sdl.Texture, duration time.Duration, scaling float64) (*animation, error) {
	var layouts []*layout
	frames := len(textures)

	for _, texture := range textures {
		_, _, width, height, err := texture.Query()
		if err != nil {
			return &animation{}, fmt.Errorf("could not query widht and height from texture: \n%v", err)
		}

		layout := &layout{
			texture: texture,
			width:   int32(float64(width) * scaling),
			height:  int32(float64(height) * scaling),
		}
		layouts = append(layouts, layout)
	}

	return &animation{
		container:       container,
		layouts:         layouts,
		duration:        duration,
		animationFrames: frames,
		changeRate:      duration / time.Duration(frames),
	}, nil
}

func (a *animation) Id() engine.ComponentId {
	return AnimationId
}

func (a *animation) Update() error {
	a.checkIndex()

	return nil
}

func (a *animation) layout() *layout {
	return a.layouts[a.currentIndex]
}

func (a *animation) checkIndex() {
	if time.Since(a.lastChange) > time.Since(time.Now())+a.changeRate {
		if a.currentIndex >= a.animationFrames-1 {
			a.currentIndex = 0
		} else {
			a.currentIndex++
		}
		a.lastChange = time.Now()
	}
}
