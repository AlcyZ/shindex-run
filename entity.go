package main

import "fmt"

type vector struct {
	x, y float64
}

type entity struct {
	position   vector
	components map[componentId]component
}

func newEntity(position vector) *entity {
	var components = make(map[componentId]component)

	return &entity{
		position:   position,
		components: components,
	}
}

func (e *entity) addComponent(c component) {
	comp := c
	e.components[comp.id()] = c
}

func (e *entity) getComponent(id componentId) (component, error) {
	if val, ok := e.components[id]; ok {
		return val, nil
	}

	return nil, fmt.Errorf("could not find component: \n%v, [%v]", id, e.components)
}

func (e *entity) update() error {
	for _, component := range e.components {
		comp := component
		if err := comp.update(); err != nil {
			return fmt.Errorf("could not update component %v: \n%v", comp.id(), err)
		}
	}
	return nil
}

func (e *entity) changePosition(position vector) {
	e.position = position
}
