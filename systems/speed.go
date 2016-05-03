package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
)

// RoadSpeed system
type roadSpeedEntity struct {
	*ecs.BasicEntity
	*engo.SpaceComponent
}

type RoadSpeedSystem struct {
	entities []roadSpeedEntity
}

func (f *RoadSpeedSystem) Add(basic *ecs.BasicEntity, space *engo.SpaceComponent) {
	f.entities = append(f.entities, roadSpeedEntity{basic, space})
}

func (f *RoadSpeedSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range f.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		f.entities = append(f.entities[:delete], f.entities[delete+1:]...)
	}
}

func (f *RoadSpeedSystem) Update(dt float32) {
	for _, e := range f.entities {
		e.SpaceComponent.Position.Y += 100 * dt
	}
}
