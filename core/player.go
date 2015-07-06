package core

type Player struct {
    Id      int64
    Name    string
    Transit []*Transit
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
