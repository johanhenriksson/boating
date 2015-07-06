package core

import (
    "fmt"
    "time"
    "math/rand"
)

const BOAT = 1

type Vehicle struct {
    Id          int64
    Type        int64
    Capacity    int64
    Speed       int64
    Name        string
    Journey     Journey
    Cargo       []*Crate
    Owner       *Player
    Orders      chan func() 
}

type Journey struct {
    From        *City
    To          *City
    Start       time.Time
    Distance    int64
    Remaining   int64
}

/* Queue an order */
func (v *Vehicle) Issue(order func()) {
    v.Orders <- order
}

func (v *Vehicle) Move(city_a *City, city_b *City) bool {
    route := city_a.Routes[city_b]

    if route == nil || !city_a.HasVehicle(v) {
        return false
    }
    city_a.Unpark(v)

    v.Journey = Journey {
        To:        route.To,
        From:      route.From,
        Distance:  route.Length,
        Remaining: route.Length,
        Start:     time.Now(),
    }

    for v.Journey.Remaining > 0 {
        v.Journey.Remaining -= rand.Int63n(v.Speed)
        time.Sleep(1 * time.Second)
    }

    v.Journey = Journey {
        To:  route.To,
        From:  route.To,
        Distance: 0,
        Remaining: 0,
        Start: time.Now(),
    }

    city_b.Harbor <- v
    for !city_b.HasVehicle(v) { 
        time.Sleep(500 * time.Millisecond) 
    }
    return true
}

func (v *Vehicle) Load(city *City, com *Commodity, quantity int64) bool {
    if !city.HasVehicle(v) {
        return false
    }

    /* TODO: Check capacity */

    /* Load time */
    time.Sleep(2 * time.Second)

    /* needs to be some kind of channel thingy */
    crate := city.Stock.Get(v.Owner, com, quantity)

    if crate == nil {
        fmt.Printf("Cannot load vehicle '%s' with %s: Not found in player stock\n", v.Name, com.Name)
        return false
    }

    v.Cargo = append(v.Cargo, crate)
    return true
}

func (v *Vehicle) UnloadAll(city *City) bool {
    if !city.HasVehicle(v) {
        fmt.Println("Cannot unload: vehicle not in city")
        return false
    }

    for _, crate := range v.Cargo {
        /* Load time */
        time.Sleep(2 * time.Second)

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
