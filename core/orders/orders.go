package orders

import (
    "fmt"
    "github.com/johanhenriksson/trade/core"
)

type Order interface {
    Execute(*core.Actor)
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
        order.Execute(actor)
    }
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

func (order GoOrder) Execute(actor *core.Actor) {
    if order.Vehicle == nil {
        fmt.Println("Error in order Go: Target vehicle is nil")
        return
    }
    actor.Issue(func() {
        order.Vehicle.Move(order.City)
    })
}

func (order GoOrder) Print() {
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

func (order *LoadOrder) Execute(actor *core.Actor) {
    if order.Vehicle == nil {
        fmt.Println("Error in order Load: Target vehicle is nil")
        return
    }
    actor.Issue(func() {
        if order.Unload {
            if order.All {
                order.Vehicle.UnloadAll()
            } else {
                order.Vehicle.Unload(order.Commodity, order.Quantity)
            }
        } else {
            order.Vehicle.Load(order.Commodity, order.Quantity)
        }
    })
}

func (order *LoadOrder) Print() {
    qty := order.Quantity
    com := order.Commodity.Name

    if order.Unload {
        if order.All {
            fmt.Println("Unload alln")
        } else {
            if qty > 0 {
                fmt.Println("Unload", qty, "x", com)
            } else {
                fmt.Println("Unload", com)
            }
        }
    } else {
        if qty > 0 {
            fmt.Println("Load", qty, "x", com)
        } else {
            fmt.Println("Load", com)
        }
    }
}
