package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
)

// Controls
type controlEntity struct {
	*ecs.BasicEntity
	*engo.SpaceComponent
}

type ControlSystem struct {
	entities []controlEntity
}

func (c *ControlSystem) Add(basic *ecs.BasicEntity, space *engo.SpaceComponent) {
	c.entities = append(c.entities, controlEntity{basic, space})
}

func (c *ControlSystem) Remove(basic ecs.BasicEntity) {
	delete := -1

	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}

	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}
}

func (c *ControlSystem) Update(dt float32) {
	speed := 400 * dt

	for _, e := range c.entities {
		if engo.Keys.Get(engo.ArrowLeft).Down() {
			e.SpaceComponent.Position.X -= speed
		}

		if engo.Keys.Get(engo.ArrowRight).Down() {
			e.SpaceComponent.Position.X += speed
		}
	}
}
