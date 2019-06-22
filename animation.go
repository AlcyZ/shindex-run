package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type animation struct {
	container       *entity
	textures        []*sdl.Texture
	lastChange      time.Time
	changeRate      time.Duration
	animationFrames int
	duration        time.Duration
	currentIndex    int
}

func newAnimation(container *entity, textures []*sdl.Texture, duration time.Duration) *animation {
	frames := len(textures)

	return &animation{
		container:       container,
		textures:        textures,
		duration:        duration,
		animationFrames: frames,
		changeRate:      duration / time.Duration(frames),
	}
}

func (a *animation) id() componentId {
	return "animation"
}

func (a *animation) update() error {
	a.checkIndex()

	return nil
}

func (a *animation) texture() *sdl.Texture {
	return a.textures[a.currentIndex]
}

// later for multiple animations needed
//func (a *animation) resetIndex() {
//	a.currentIndex = 0
//}

func (a *animation) checkIndex() {
	println("easy")
	if time.Since(a.lastChange) > time.Since(time.Now())+a.changeRate {
		if a.currentIndex >= a.animationFrames-1 {
			a.currentIndex = 0
		} else {
			a.currentIndex++
		}
		a.lastChange = time.Now()
	}
}
