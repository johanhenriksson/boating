package main

import (
//    "fmt"
    "time"
    "net/http"

    "github.com/johanhenriksson/trade/core"
    "github.com/johanhenriksson/trade/api"
)

func main() {
    city_a := core.NewCity("London")
    city_b := core.NewCity("Amsterdam")
    city_c := core.NewCity("Hamburg")

    city_a.AddRoute(city_b, 160)
    city_a.AddRoute(city_c, 70)
    city_b.AddRoute(city_a, 160)

    player := &core.Player {
        Id: 123,
        Name: "jojje",
    }

    world := &core.World {
        Cities: map[int64]*core.City {
            city_a.Id: city_a,
            city_b.Id: city_b,
            city_c.Id: city_c,
        },
        Players: map[int64]*core.Player {
            123: player,
        },
    }


    city_a.Stock.Store(core.GetCrate(player, core.GOLD, 10000000000))
    city_b.Stock.Store(core.GetCrate(player, core.COFFEE, 1000000000))

    for i := 0; i < 1000000; i++ {
        boat := core.NewBoat("HMS Boat", player, 1000)
        player.AddVehicle(boat)
        city_a.Park(boat)
        boat.Issue(func() {
            for {
                boat.Load(city_a, core.GOLD, 45)
                boat.Move(city_a, city_b)
                boat.UnloadAll(city_b)
                boat.Load(city_b, core.COFFEE, 77)
                boat.Move(city_b, city_a)
                boat.UnloadAll(city_a)
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
