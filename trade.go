package main

import (
    "os"
    "fmt"
    "time"
    "bufio"
    "strings"
    "strconv"
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

    fmt.Println("boat server")

    /* Stdio console */
    go console()

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

func console() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("> ")

        text, _ := reader.ReadString('\n')
        line := strings.ToLower(strings.Trim(strings.Trim(text, "\n"), " "))
        tokens := strings.Split(line, " ")

        switch tokens[0] {
        case "help":
            fmt.Println("Helpful Help Menu:")
            fmt.Println("  time")
            fmt.Println("  timescale <int scale>")
            fmt.Println("  quit")
        case "quit":
            os.Exit(0)
        case "time":
            fmt.Println("Server time:", core.Time())
        case "timescale":
            if len(tokens) < 2 {
                fmt.Println("usage: timescale <int scale>")
                continue
            }
            if tokens[1] == "reset" {
                fmt.Println("Timescale: 1hr per second")
                core.TIMESCALE = time.Duration(3600)
                continue
            }
            scale, err := strconv.ParseInt(tokens[1], 10, 64)
            if err != nil {
                fmt.Println("Invalid number")
                continue
            }
            core.TIMESCALE = time.Duration(scale)
            t := core.TIMESCALE * time.Second
            fmt.Println("Server timescale:", t, "per second") 

        default:
            fmt.Printf("Invalid command '%s'. Type 'help' to show a helpful help menu\n", tokens[0])
        }
    }
}
