package orders

import (
    "fmt"
    "github.com/johanhenriksson/boating/core"
)

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
