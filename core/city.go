package core

import (
    "fmt"
)

type City struct {
    Id          int64
    Name        string
    Vehicles    map[int64]*Vehicle
    Routes      map[*City]*Route
    Stock       Storage
    Harbor      chan*Vehicle
}

func NewCity(name string) *City {
    city := &City {
        Id:         nextId(),
        Name:       name,
        Vehicles:   make(map[int64]*Vehicle),
        Routes:     make(map[*City]*Route),
        Stock:      make(Storage),
        Harbor:     make(chan*Vehicle),
    }
    go CityWorker(city)
    return city
}

func (city *City) AddRoute(target *City, distance int64) *Route {
    route := &Route {
        From: city,
        To: target,
        Length: distance,
    }
    city.Routes[target] = route
    return route
}

func (city *City) Park(vehicle *Vehicle) {
    city.Vehicles[vehicle.Id] = vehicle
}

func (city *City) Unpark(vehicle *Vehicle) {
    city.Vehicles[vehicle.Id] = nil
}

func (city *City) HasVehicle(vehicle *Vehicle) bool {
    _, exists := city.Vehicles[vehicle.Id]
    return exists
}

func (city *City) Print() {
    fmt.Println("------------------")
    fmt.Println("City:", city.Name)
    for player, crates := range city.Stock {
        fmt.Println("Player", player.Name)
        for _, crate := range crates {
            fmt.Printf("  %d x %s\n", crate.Qty, crate.Type.Name)
        }
        fmt.Println("")
    }
    fmt.Println("Vehicles:")
    for _, vehicle := range city.Vehicles {
        fmt.Printf("  %d %s (%s)\n", vehicle.Id, vehicle.Name, vehicle.Owner.Name)
    }
    fmt.Println("------------------")
}

func CityWorker(city *City) {
    for {
        select {
        case boat := <-city.Harbor:
            fmt.Printf("Boat %s arrived in %s\n", boat.Name, city.Name)
            city.Park(boat)

            city.Print()
        }
    }
}
