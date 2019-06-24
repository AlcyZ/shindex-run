package engine

import (
	"fmt"
)

type Vector struct {
	X, Y float64
}

func NewVector(x float64, y float64) Vector {
	return Vector{X: x, Y: y}
}

type Entity struct {
	game       *Game
	position   Vector
	components map[ComponentId]Component
}

func NewEntity(game *Game, position Vector) *Entity {
	var components = make(map[ComponentId]Component)

	return &Entity{
		game:       game,
		position:   position,
		components: components,
	}
}

func (e *Entity) AddComponent(c Component) {
	comp := c
	e.components[comp.Id()] = c
}

func (e *Entity) GetComponent(id ComponentId) (Component, error) {
	if val, ok := e.components[id]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("could not find Component: \n%v, [%v]", id, e.components)
}

func (e *Entity) Update() error {
	for _, component := range e.components {
		comp := component
		if err := comp.Update(); err != nil {
			return fmt.Errorf("could not update Component %v: \n%v", comp.Id(), err)
		}
	}
	return nil
}

func (e *Entity) ChangePosition(position Vector) {
	e.position = position
}

func (e *Entity) CurrentPosition() Vector {
	return e.position
}

func (e *Entity) GetDelta() float64 {
	return e.game.Delta()
}
