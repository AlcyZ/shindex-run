package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func newPlayer(r *sdl.Renderer, speed float64, path string) (*entity, error) {
	initPos := vector{x: 50, y: screenHeight - 220}
	player := newEntity(initPos)

	idleAnimation, err := getPlayerIdleAnimation(player, r)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create player idle animation: \n%v", err)
	}
	runAnimation, err := getPlayerRunAnimation(player, r)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create player run animation: \n%v", err)
	}

	var idle = "idle"
	var left = "left"
	var right = "right"

	animations := newAnimations()
	animations.add(idleAnimation, sdl.FLIP_NONE, idle)
	animations.add(runAnimation, sdl.FLIP_HORIZONTAL, left)
	animations.add(runAnimation, sdl.FLIP_NONE, right)
	player.addComponent(animations)

	control := newHorizontalControl(player, speed, getLeftKeys(), getRightKeys())
	err = control.withAnimations(idle, left, right) // animations must be attach to player component first
	if err != nil {
		return &entity{}, fmt.Errorf("could not add animations to horizontal control: \n%v", err)
	}

	player.addComponent(control)

	// the render component should be the last attached, because its very likely that other components updates the
	// internal state to be rendered
	renderer, err := newAnimationsRenderer(player, r)
	if err != nil {
		return &entity{}, fmt.Errorf("could not create animations renderer: \n%v", err)
	}
	player.addComponent(renderer)

	return player, nil
}

func getPlayerIdleAnimation(container *entity, r *sdl.Renderer) (*animation, error) {
	var idleTxt []*sdl.Texture

	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("assets/ninja/Idle__00%d.png", i)
		txt, err := img.LoadTexture(r, path)
		if err != nil {
			return &animation{}, fmt.Errorf("could not load player idle texture %v, \n%v", i, err)
		}

		idleTxt = append(idleTxt, txt)
	}

	anim, err := newAnimation(container, idleTxt, time.Second, 0.25)
	if err != nil {
		return &animation{}, fmt.Errorf("could not create idle animation: \n%v", err)
	}

	return anim, nil
}

func getPlayerRunAnimation(container *entity, r *sdl.Renderer) (*animation, error) {
	var runTxt []*sdl.Texture

	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("assets/ninja/Run__00%d.png", i)
		txt, err := img.LoadTexture(r, path)
		if err != nil {
			return &animation{}, fmt.Errorf("could not load player run texture %v, \n%v", i, err)
		}

		runTxt = append(runTxt, txt)
	}

	anim, err := newAnimation(container, runTxt, time.Second, 0.25)
	if err != nil {
		return &animation{}, fmt.Errorf("could not create run animation: \n%v", err)
	}

	return anim, nil
}

func getLeftKeys() []sdl.Scancode {
	keys := make([]sdl.Scancode, 2)

	keys = append(keys, sdl.SCANCODE_A)
	keys = append(keys, sdl.SCANCODE_LEFT)

	return keys
}

func getRightKeys() []sdl.Scancode {
	keys := make([]sdl.Scancode, 2)

	keys = append(keys, sdl.SCANCODE_D)
	keys = append(keys, sdl.SCANCODE_RIGHT)

	return keys
}
