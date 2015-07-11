package core

var World = NewWorld()

type WorldState struct {
    Cities      CityMap
    Players     PlayerMap
    Vehicles    VehicleMap
}

func NewWorld() *WorldState {
    world := &WorldState {
        Cities:     Cities,
        Players:    make(PlayerMap),
        Vehicles:   make(VehicleMap),
    }
    return world
}
