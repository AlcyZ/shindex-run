package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
	"time"
)

func newEnemy(r *sdl.Renderer, path string) (*engine.Entity, error) {
	initPos := engine.NewVector(screenWidth-150, screenHeight-240)
	player := engine.NewEntity(initPos)

	animation, err := getEnemyIdleAnimation(player, r)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not create enemy idle animation: \n%v", err)
	}
	player.AddComponent(animation)

	renderer, err := newAnimationRenderer(player, r)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not create animation renderer: \n%v", err)
	}
	player.AddComponent(renderer)

	return player, nil
}

func getEnemyIdleAnimation(container *engine.Entity, r *sdl.Renderer) (*animation, error) {
	var textures []*sdl.Texture

	for i := 0; i < 15; i++ {
		path := fmt.Sprintf("assets/player/male/Idle_%d.png", i)
		texture, err := img.LoadTexture(r, path)
		if err != nil {
			return &animation{}, fmt.Errorf("could not load enemy idle texture: \n%v", err)
		}
		textures = append(textures, texture)
	}

	anim, err := newAnimation(container, textures, time.Second, 0.25)
	if err != nil {
		return &animation{}, fmt.Errorf("could not create enemy idle animation: \n%v", err)
	}

	return anim, nil
}
