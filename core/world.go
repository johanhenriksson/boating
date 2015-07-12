package core

var World = NewWorld()

type WorldState struct {
    Cities      CityMap
    Players     PlayerMap
    Vehicles    *Garage
}

func NewWorld() *WorldState {
    world := &WorldState {
        Cities:     Cities,
        Players:    make(PlayerMap),
        Vehicles:   NewGarage(),
    }
    return world
}
