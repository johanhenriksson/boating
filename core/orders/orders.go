package orders

import (
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

/* Looks for vehicle orders and sets their vehicle field.
   Used before creating order functions */
func (orders *Orders) SetVehicle(v *core.Vehicle) {
    for _, order := range *orders {
        if vOrder, ok := order.(VehicleOrder); ok {
            vOrder.SetVehicle(v)
        }
    }
}

/* Orders an actor to execute the commands */
func (orders Orders) Execute(actor *core.Actor) {
    for _, order := range orders {
        command := order.Create()
        actor.Issue(command)
    }
}

/* Orders an actor to loop the commands indefinately */
func (orders Orders) Loop(actor *core.Actor) {
    commands := make([]func(), 0)
    for _, order := range orders {
        command := order.Create()
        commands = append(commands, command)
    }
    actor.Loop(func() {
        for _, command := range commands {
            command()
        }
    })
}

func (orders Orders) Print() {
    for _, order := range orders {
        order.Print()
    }
}
