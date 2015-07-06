package core

import (
    "fmt"
    "time"
)

const BOAT = 1

type Vehicle struct {
    Id          int64
    Type        int64
    Capacity    int64
    Speed       int64
    Name        string
    Cargo       []*Crate
    Owner       *Player
    Orders      chan func()
}

/* Queue an order */
func (v *Vehicle) Issue(order func()) {
    v.Orders <- order
}

func (v *Vehicle) Move(city_a *City, city_b *City) bool {
    route := city_a.Routes[city_b]

    /* TODO: Make sure we're at the start point */
    if !city_a.HasVehicle(v) {
        return false
    }
    city_a.Unpark(v)

    remaining := route.Length
    for remaining > 0 {
        fmt.Printf("Moving. %dkm remaining\n", remaining)
        remaining -= v.Speed
        time.Sleep(1 * time.Second)
    }

    fmt.Printf("Boat %s move completed\n", v.Name)
    city_b.Harbor <- v
    fmt.Printf("In city b:", city_b.HasVehicle(v))
    return true
}

func (v *Vehicle) Load(city *City, com *Commodity, quantity int64) bool {
    if !city.HasVehicle(v) {
        return false
    }

    /* TODO: Check capacity */

    /* needs to be some kind of channel thingy */
    crate := city.Stock.Get(v.Owner, com, quantity)

    if crate == nil {
        fmt.Printf("Cannot load vehicle '%s' with %s: Not found in player stock\n", v.Name, com.Name)
        return false
    }

    v.Cargo = append(v.Cargo, crate)
    fmt.Printf("%s loaded %d x %s\n", v.Name, crate.Qty, crate.Type.Name)
    return true
}

func (v *Vehicle) UnloadAll(city *City) bool {
    if !city.HasVehicle(v) {
        return false
    }

    for _, crate := range v.Cargo {
        fmt.Printf("%s unloading %d x %s\n", v.Name, crate.Qty, crate.Type.Name)

        /* needs to be some kind of channel thingy */
        city.Stock.Store(crate)
    }
    v.Cargo = []*Crate { }
    return true
}

func (v *Vehicle) Unload(crate *Crate, city *City) bool {
    if !city.HasVehicle(v) {
        return false
    }

    for i, c := range v.Cargo {
        if c.Id == crate.Id {
            /* Remove crate from cargo */
            v.Cargo = append(v.Cargo[0:i], v.Cargo[i+1:]...)
            fmt.Printf("%s unloading %d x %s\n", v.Name, crate.Qty, crate.Type.Name)
            city.Stock.Store(crate)
            return true
        }
    }
    return false
}

func VehicleWorker(v *Vehicle) {
    fmt.Printf("Vehicle %v waiting for order\n", v.Id)
    for {
        order := <-v.Orders
        order()
    }
}

func NewBoat(name string, owner *Player, capacity int64) *Vehicle {
    boat := &Vehicle {
        Id:         nextId(),
        Name:       name,
        Type:       BOAT,
        Capacity:   capacity,
        Speed:      10,
        Cargo:      []*Crate { },
        Owner:      owner,
        Orders:     make(chan func(), 5),
    }
    go VehicleWorker(boat)
    return boat
}
