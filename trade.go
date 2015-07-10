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

    core.LONDON.Stock.Store(core.GetCrate(player, core.GOLD,     10000000000))
    core.AMSTERDAM.Stock.Store(core.GetCrate(player, core.COFFEE,   1000000000))
    core.HAMBURG.Stock.Store(core.GetCrate(player, core.STEEL,    10000000))
    core.WASHINGTON.Stock.Store(core.GetCrate(player, core.WEAPONS,  10000000))

    for i := 0; i < 1; i++ {
        boat := core.NewBoat(player, core.LONDON, fmt.Sprintf("HMS Boat #%d", 100+i))

        orders := orders.Compile("load gold 50; go hamburg; unload gold 50; go london")
        orders.SetVehicle(boat)
        orders.Execute(boat.Actor)

        boat.Issue(func() {
            for {
                /* Load up cash */
                boat.Load(core.GOLD, 45)

                /* Go to amsterdam */
                boat.Move(core.AMSTERDAM)

                /* Drop gold & pick up coffee in amsterdam */
                boat.UnloadAll()
                boat.Load(core.COFFEE, 77)

                /* pickup steel in hamburg */
                boat.Move(core.HAMBURG)
                boat.Load(core.STEEL, 140)

                /* Return to london, dump goods */
                boat.Move(core.LONDON)
                boat.UnloadAll()
            }
        })
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
