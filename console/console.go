package console

import (
    "os"
    "fmt"
    "time"
    "bufio"
    "strings"
    "strconv"

    "github.com/johanhenriksson/boating/core"
    "github.com/johanhenriksson/boating/core/compiler"
)

func Run() {
    reader := bufio.NewReader(os.Stdin)
    world  := core.World
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
            fmt.Println("  vehicles")
            fmt.Println("  order <int VehicleId> <order> [; <order>]...")
            fmt.Println("  stop <int VehicleId>")
            fmt.Println("  quit")
        case "quit":
            os.Exit(0)
        case "time":
            fmt.Println("Server time:", core.Time())
        case "stop":
            if len(tokens) < 2 {
                fmt.Println("usage: stop <int VehicleId>")
                continue
            }
            for i := 1; i < len(tokens); i++ {
                id, err := strconv.ParseInt(tokens[i], 10, 64)
                if err != nil {
                    fmt.Println("Invalid number: %s", tokens[i])
                    continue
                }
                if ok, v := world.Vehicles.Find(core.VehicleId(id)); ok {
                    v.Stop()
                    fmt.Println("Stopping", v.Name)
                } else {
                    fmt.Println("Unknown vehicle", id)
                }
            }

        case "order":
            if len(tokens) < 3 {
                fmt.Println("usage: order <int VehicleId> <order> [; <order>]...")
                continue
            }
            id, err := strconv.ParseInt(tokens[1], 10, 64)
            if err != nil {
                fmt.Println("Invalid number: %s", tokens[1])
                continue
            }
            if ok, v := world.Vehicles.Find(core.VehicleId(id)); ok {
                /* Compile order */
                order := compiler.Compile(strings.Join(tokens[2:], " "))
                order.SetVehicle(v)

                fmt.Println("Ordering", v.Name, "to:")
                order.Print()

                order.Execute(v.Actor)
            } else {
                fmt.Println("Unknown vehicle #%d", id)
            }

        case "vehicles":
            fmt.Println("All vehicles")
            for _, v := range world.Vehicles.Vehicles {
                fmt.Println(" ", v)
            }

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
