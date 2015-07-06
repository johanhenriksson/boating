package main

import (
    "fmt"
    "time"

    "github.com/johanhenriksson/trade/core"
)

func main() {
    city_a := core.NewCity("City A")
    city_b := core.NewCity("City B")
    city_c := core.NewCity("City C")

    city_a.AddRoute(city_b, 50)
    city_a.AddRoute(city_c, 70)
    city_b.AddRoute(city_a, 50)

    player := &core.Player {
        Name: "jojje",
    }

    boat := core.NewBoat("HMS Boat", player, 1000)
    city_a.Park(boat)

    city_a.Stock.Store(core.GetCrate(player, core.GOLD, 100))
    city_b.Stock.Store(core.GetCrate(player, core.COFFEE, 100))

    city_a.Print()

    boat.Issue(func() {
        boat.Load(city_a, core.GOLD, 50)
        boat.Move(city_a, city_b)
        boat.UnloadAll(city_b)
        boat.Load(city_b, core.COFFEE, 100)
        boat.Move(city_b, city_a)
        boat.UnloadAll(city_a)
        boat.Load(city_a, core.GOLD, 50)
        boat.Move(city_a, city_b)
        boat.UnloadAll(city_b)
        boat.Move(city_b, city_a)
        city_b.Print()
    })

    fmt.Println("concurrent?")

    for {
        time.Sleep(60 * time.Second)
    }
}
