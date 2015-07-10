package main

import (
    "fmt"
    "time"
    "net/http"

    "github.com/johanhenriksson/trade/api"
    "github.com/johanhenriksson/trade/core"
    "github.com/johanhenriksson/trade/core/orders"
)

func main() {
    player := &core.Player {
        Id: 123,
        Name: "jojje",
    }

    world := &core.World {
        Cities: core.Cities,
        Players: core.PlayerMap {
            123: player,
        },
    }

    core.LONDON.Stock.Store(core.NewCrate(player,     core.GOLD,    100000))
    core.AMSTERDAM.Stock.Store(core.NewCrate(player,  core.COFFEE,  70000))
    core.HAMBURG.Stock.Store(core.NewCrate(player,    core.STEEL,   50000))
    core.WASHINGTON.Stock.Store(core.NewCrate(player, core.WEAPONS, 10000))

    for i := 0; i < 10; i++ {
        boat := core.NewBoat(player, core.LONDON, fmt.Sprintf("HMS Boat #%d", 100+i))

        orders := orders.CompileFile("scripts/merchant.txt")
        orders.Print()

        orders.SetVehicle(boat)
        orders.Execute(boat.Actor)
    }

    router := api.NewRouter()
    router.Register(&api.VehicleService {
        World: world,
    })
    router.Register(&api.CityService {
        World: world,
    })
    router.Files("/", "./html/")

    http.ListenAndServe(":8000", router.Mux())

    for {
        time.Sleep(60 * time.Second)
    }
}
