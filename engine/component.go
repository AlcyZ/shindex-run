package engine

type ComponentId string

type Component interface {
	Id() ComponentId
	Update() error
}
