package main

import (
    "fmt"
    "time"
    "net/http"

    "github.com/johanhenriksson/trade/core"
    "github.com/johanhenriksson/trade/api"
)

func main() {
    city_a := core.NewCity("London")
    city_b := core.NewCity("Amsterdam")
    city_c := core.NewCity("Hamburg")
    city_d := core.NewCity("Washington")

    city_a.AddRoute(city_b, 160)
    city_b.AddRoute(city_a, 160)
    city_b.AddRoute(city_c, 70)
    city_c.AddRoute(city_b, 70)
    city_c.AddRoute(city_a, 200)
    city_a.AddRoute(city_c, 200)

    player := &core.Player {
        Id: 123,
        Name: "jojje",
    }

    world := &core.World {
        Cities: map[int64]*core.City {
            city_a.Id: city_a,
            city_b.Id: city_b,
            city_c.Id: city_c,
            city_d.Id: city_d,
        },
        Players: map[int64]*core.Player {
            123: player,
        },
    }

    city_a.Stock.Store(core.GetCrate(player, core.GOLD,     10000000000))
    city_b.Stock.Store(core.GetCrate(player, core.COFFEE,   1000000000))
    city_c.Stock.Store(core.GetCrate(player, core.STEEL,    10000000))
    city_d.Stock.Store(core.GetCrate(player, core.WEAPONS,  10000000))

    for i := 0; i < 10; i++ {
        boat := core.NewBoat(fmt.Sprintf("HMS Boat #%d", i), player, 1000)
        boat.Journey.From = city_a
        boat.Journey.To = city_a
        player.AddVehicle(boat)
        city_a.Park(boat)
        boat.Issue(func() {
            for {
                /* Load up cash */
                boat.Load(city_a, core.GOLD, 45)

                /* Go to amsterdam */
                boat.Move(city_a, city_b)

                /* Drop gold & pick up coffee in amsterdam */
                boat.UnloadAll(city_b)
                boat.Load(city_b, core.COFFEE, 77)

                /* pickup steel in hamburg */
                boat.Move(city_b, city_c)
                boat.Load(city_c, core.STEEL, 140)

                /* Return to london, dump goods */
                boat.Move(city_c, city_a)
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
