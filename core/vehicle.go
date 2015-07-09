package core

import (
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
    Cargo       *Storage
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
    city_a.Embark <- v

    v.Journey = Journey {
        To:        route.To,
        From:      route.From,
        Distance:  route.Length,
        Remaining: route.Length,
        Start:     time.Now(),
    }

    /* Perform movement */
    for v.Journey.Remaining > 0 {
        time.Sleep(1 * time.Second)
        v.Journey.Remaining -= rand.Int63n(v.Speed)
    }
    v.Journey.Remaining = 0

    /* Reset journey */
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
    loadAmount := quantity
    var loadStep int64 = 10
    for loadAmount > 0 {
        if (loadAmount < loadStep) {
            loadStep = loadAmount
        }
        loadAmount -= loadStep

        crate := city.Stock.Get(v.Owner, com, loadStep)
        if crate == nil {
            return false
        }
        time.Sleep(1 * time.Second)
        v.Cargo.Store(crate)
    }

    return true
}

func (v *Vehicle) UnloadAll(city *City) bool {
    if !city.HasVehicle(v) {
        return false
    }
    for _, crates:= range v.Cargo.Crates {
        for _, crate := range crates {
            v.Unload(crate, city)
        }
    }
    return true
}

func (v *Vehicle) Unload(crate *Crate, city *City) bool {
    if !city.HasVehicle(v) {
        return false
    }

    loadAmount := crate.Qty
    var loadStep int64 = 10
    for loadAmount > 0 {
        if (loadAmount < loadStep) {
            loadStep = loadAmount
        }
        loadAmount -= loadStep

        crate := v.Cargo.Get(v.Owner, crate.Type, loadStep)
        if crate == nil {
            return false
        }
        time.Sleep(1 * time.Second)
        city.Stock.Store(crate)
    }

    return true
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
        Cargo:      NewStorage(),
        Owner:      owner,
        Orders:     make(chan func(), 5),
    }
    go VehicleWorker(boat)
    return boat
}
