package main

type componentId string

type component interface {
	id() componentId
	update() error
}
