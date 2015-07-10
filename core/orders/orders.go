package orders

import (
    "fmt"
    "time"
    "github.com/johanhenriksson/boating/core"
)

type Order interface {
    Create() func()
    Print()
}

type VehicleOrder interface {
    SetVehicle(*core.Vehicle)
}

type Orders []Order

func (orders *Orders) SetVehicle(v *core.Vehicle) {
    for _, order := range *orders {
        if vOrder, ok := order.(VehicleOrder); ok {
            vOrder.SetVehicle(v)
        }
    }
}

func (orders Orders) Execute(actor *core.Actor) {
    for _, order := range orders {
        command := order.Create()
        actor.Issue(command)
    }
}

func (orders Orders) Loop(actor *core.Actor) {
    commands := make([]func(), 0)
    for _, order := range orders {
        command := order.Create()
        commands = append(commands, command)
    }
    actor.Issue(func() {
        for {
            for _, command := range commands {
                command()
            }
        }
    })
}

func (orders Orders) Print() {
    for _, order := range orders {
        order.Print()
    }
}

type GoOrder struct {
    City    *core.City
    Vehicle *core.Vehicle
}

/* For VehicleOrder interface */
func (order *GoOrder) SetVehicle(v *core.Vehicle) {
    order.Vehicle = v
}

func (order *GoOrder) Create() func() {
    v, city := order.Vehicle, order.City
    return func() {
        v.Move(city)
    }
}

func (order *GoOrder) Print() {
    fmt.Println("Go to", order.City.Name)
}

type LoadOrder struct {
    Vehicle     *core.Vehicle
    Commodity   *core.Commodity
    Quantity    int64
    Unload      bool
    All         bool
}

/* For VehicleOrder interface */
func (order *LoadOrder) SetVehicle(v *core.Vehicle) {
    order.Vehicle = v
}

func (order *LoadOrder) Create() func() {
    v := order.Vehicle
    unload, all := order.Unload, order.All
    return func() {
        if unload {
            if all {
                v.UnloadAll()
            } else {
                v.Unload(order.Commodity, order.Quantity)
            }
        } else {
            v.Load(order.Commodity, order.Quantity)
        }
    }
}

func (order *LoadOrder) Print() {
    qty := order.Quantity

    if order.Unload {
        if order.All {
            fmt.Println("Unload all")
        } else {
            com := order.Commodity.Name
            if qty > 0 {
                fmt.Println("Unload", qty, "x", com)
            } else {
                fmt.Println("Unload", com)
            }
        }
    } else {
        com := order.Commodity.Name
        if qty > 0 {
            fmt.Println("Load", qty, "x", com)
        } else {
            fmt.Println("Load", com)
        }
    }
}

type WaitOrder struct {
    Duration    time.Duration
}

func (order *WaitOrder) Create() func() {
    duration := order.Duration
    return func() {
        core.Sleep(duration)
    }
}

func (order *WaitOrder) Print() {
    fmt.Println("Wait", order.Duration)
}
