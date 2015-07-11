package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "net/http"

    "github.com/johanhenriksson/boating/api"
    "github.com/johanhenriksson/boating/core"
    "github.com/johanhenriksson/boating/core/compiler"
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

    router := api.NewRouter()
    router.Register(&api.VehicleService {
        World: core.World,
    })
    router.Register(&api.CityService {
        World: core.World,
    })
    router.Files("/", "./html/")

    go console()
    http.ListenAndServe(":8000", router.Mux())
}

func console() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("hello?")
    for {
        fmt.Print("> ")
        text, _ := reader.ReadString('\n')
        line := strings.ToLower(strings.Trim(strings.Trim(text, "\n"), " "))
        tokens := strings.Split(line, " ")

        switch tokens[0] {
        case "help":
            fmt.Println("Helpful Help Menu:")
            fmt.Println("quit - quit.")
        case "quit":
            os.Exit(0)
        default:
            fmt.Println("Invalid command. Type help to activate helpful help menu")
        }
    }
}
