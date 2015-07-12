package main

import (
    "fmt"
    "net/http"

    "github.com/johanhenriksson/boating/core"
    "github.com/johanhenriksson/boating/core/compiler"
    "github.com/johanhenriksson/boating/console"
    "github.com/johanhenriksson/boating/api"
)

func main() {
    player := core.NewPlayer(123, "jojje")

    core.LONDON.Stock.Store(core.NewCrate(player,     core.GOLD,    100000))
    core.AMSTERDAM.Stock.Store(core.NewCrate(player,  core.COFFEE,  70000))
    core.HAMBURG.Stock.Store(core.NewCrate(player,    core.STEEL,   50000))
    core.WASHINGTON.Stock.Store(core.NewCrate(player, core.EXPLOSIVES, 10000))

    orders := compiler.CompileFile("scripts/amsterdam_coffee.txt")
    for i := 0; i < 10; i++ {
        boat := core.NewBoat(player, core.LONDON, fmt.Sprintf("HMS Boat #%d", 100+i))

        orders.SetVehicle(boat)
        orders.Loop(boat.Actor)
    }

    fmt.Println("boat server")

    /* Stdio console */
    go console.Run()

    /* Run http interface server on main thread */
    httpServe(core.World)
}

func httpServe(world *core.WorldState) {
    router := api.NewRouter()
    router.Register(&api.VehicleService {
        World: core.World,
    })
    router.Register(&api.CityService {
        World: core.World,
    })
    router.Files("/", "./html/")

    http.ListenAndServe(":8000", router.Mux())
}
