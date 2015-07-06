package core

type Player struct {
    Id          int64
    Name        string

    Transit     []*Transit
    Vehicles    []*Vehicle
}

func (player *Player) AddVehicle(vehicle *Vehicle) {
    player.Vehicles = append(player.Vehicles, vehicle)
}

func (player *Player) AddTransit(transit *Transit) {
    player.Transit = append(player.Transit, transit)
}

type Transit struct {
    Id          int64
    Vehicle     *Vehicle
    Route       *Route
    /* Depart timestamp */
}
