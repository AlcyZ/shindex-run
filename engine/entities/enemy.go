package entities

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"shindex-run/engine"
	"shindex-run/engine/components"
	"time"
)

func NewEnemy(game *engine.Game, r *sdl.Renderer, path string) (*engine.Entity, error) {
	initPos := engine.NewVector(engine.ScreenWidth-150, engine.ScreenHeight-240)
	enemy := engine.NewEntity(game)
	enemy.ChangePosition(initPos)

	animation, err := getEnemyIdleAnimation(enemy, r)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not create enemy idle animation: \n%v", err)
	}
	enemy.AddComponent(animation)

	renderer, err := components.NewAnimationRenderer(enemy, r)
	if err != nil {
		return &engine.Entity{}, fmt.Errorf("could not create animation renderer: \n%v", err)
	}
	enemy.AddComponent(renderer)

	return enemy, nil
}

func getEnemyIdleAnimation(container *engine.Entity, r *sdl.Renderer) (*components.Animation, error) {
	var textures []*sdl.Texture

	for i := 0; i < 15; i++ {
		path := fmt.Sprintf("assets/player/male/Idle_%d.png", i)
		texture, err := img.LoadTexture(r, path)
		if err != nil {
			return &components.Animation{}, fmt.Errorf("could not load enemy idle texture: \n%v", err)
		}
		textures = append(textures, texture)
	}

	anim, err := components.NewAnimation(container, textures, time.Second, 0.25, sdl.FLIP_HORIZONTAL)
	if err != nil {
		return &components.Animation{}, fmt.Errorf("could not create enemy idle animation: \n%v", err)
	}

	return anim, nil
}
