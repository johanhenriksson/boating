package core

import (
    "fmt"
    "time"
    "math/rand"
)

const BOAT = 1

type VehicleId  int64

type Vehicle struct {
    *Actor
    Id          VehicleId
    Type        int64
    Capacity    int64
    Speed       int64
    Name        string
    Journey     Journey
    Cargo       *Storage
    Owner       *Player
    City        *City
}

type Journey struct {
    From        *City
    To          *City
    Start       time.Time
    Distance    Distance
    Remaining   Distance
}

func NewBoat(owner *Player, city *City, name string) *Vehicle {
    boat := &Vehicle {
        Actor:      NewActor(),
        Id:         VehicleId(nextId()),
        Name:       name,
        Type:       BOAT,
        Capacity:   100,
        Speed:      25,
        Cargo:      NewStorage(),
        Owner:      owner,
        City:       city,
        Journey:    Journey {
            From:   city,
            To:     city,
            Start:  time.Now(),
            Distance:  0,
            Remaining: 0,
        },
    }
    owner.AddVehicle(boat)
    city.Vehicles.Add(boat)
    return boat
}

/* Returns true if the vehicle is currently in a city */
func (v *Vehicle) InCity() bool {
    if v.City == nil {
        return false
    }
    return v.City.Vehicles.Stores(v)
}

func (v *Vehicle) Move(city_b *City) bool {
    city_a := v.City
    route := city_a.Routes[city_b]

    if route == nil || !city_a.Vehicles.Stores(v) {
        return false
    }
    city_a.Vehicles.Unpark(v)
    v.City = city_b

    v.Journey = Journey {
        To:        route.To,
        From:      route.From,
        Distance:  route.Length,
        Remaining: route.Length,
        Start:     time.Now(),
    }

    /* Perform movement */
    for v.Journey.Remaining > 0 {
        Sleep(1 * time.Hour) /* Kilometers per hour */
        v.Journey.Remaining -= Distance(rand.Int63n(v.Speed))
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

    city_b.Vehicles.Park(v)
    return true
}

func (v *Vehicle) Load(com *Commodity, quantity int64) bool {
    if !v.InCity() {
        fmt.Println("Cannot load ship", v.Id, "- not in town")
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

        crate, err := v.City.Stock.Get(v.Owner, com, loadStep)
        if err != nil {
            return false
        }
        Sleep(1 * time.Hour)
        v.Cargo.Store(crate)
    }

    return true
}

func (v *Vehicle) UnloadAll() bool {
    if !v.InCity() {
        fmt.Println("Cannot unload ship", v.Id, "- not in town")
        return false
    }

    for _, crates:= range v.Cargo.Crates {
        for _, crate := range crates {
            v.Unload(crate.Type, crate.Qty)
        }
    }

    return true
}

func (v *Vehicle) Unload(com *Commodity, qty int64) bool {
    if !v.InCity() {
        fmt.Println("Cannot unload ship", v.Id, "- not in town")
        return false
    }

    loadAmount := qty
    var loadStep int64 = 10
    for loadAmount > 0 {
        if (loadAmount < loadStep) {
            loadStep = loadAmount
        }
        loadAmount -= loadStep

        crate, err := v.Cargo.Get(v.Owner, com, loadStep)
        if err != nil {
            return false
        }
        Sleep(1 * time.Hour)
        v.City.Stock.Store(crate)
    }

    return true
}
