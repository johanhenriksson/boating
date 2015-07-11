package core

type PlayerId int64
type PlayerMap map[PlayerId]*Player

type Player struct {
    Id          PlayerId
    Name        string
    Vehicles    []*Vehicle
}

func NewPlayer(id PlayerId, name string) *Player {
    player := &Player {
        Id: id,
        Name: name,
        Vehicles: make([]*Vehicle, 0),
    }
    World.Players[player.Id] = player
    return player
}

func (player *Player) AddVehicle(vehicle *Vehicle) {
    player.Vehicles = append(player.Vehicles, vehicle)
}
