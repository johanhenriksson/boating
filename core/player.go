package core

type PlayerId int64
type PlayerMap map[PlayerId]*Player

type Player struct {
    Id          PlayerId
    Name        string
    Vehicles    []*Vehicle
}

func (player *Player) AddVehicle(vehicle *Vehicle) {
    player.Vehicles = append(player.Vehicles, vehicle)
}
