package orders

import (
    "fmt"
    "github.com/johanhenriksson/boating/core"
)

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
