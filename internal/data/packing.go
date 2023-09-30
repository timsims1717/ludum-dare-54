package data

import (
	"github.com/bytearena/ecs"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/world"
)

type Item struct {
	Key    string
	Name   string
	Shape  []world.Coords
	Object *object.Object
	Entity *ecs.Entity
	Sprite *img.Sprite
}

var ItemQueue [8]*Item
