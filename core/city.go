package core

import (
    "sync"
)

type CityId int64
type RouteMap map[*City]*Route

type City struct {
    Id          CityId
    Name        string
    Stock       *Storage
    Vehicles    *Garage
    Routes      RouteMap
    routeLock   *sync.RWMutex
}

func NewCity(id CityId, name string) *City {
    city := &City {
        Id:         id,
        Name:       name,
        Routes:     make(RouteMap),
        Stock:      NewStorage(),
        Vehicles:   NewGarage(),
    }
    return city
}

/* Add an available transport route originating from this city */
func (city *City) addRoute(target *City, distance Distance) *Route {
    route := &Route {
        From: city,
        To: target,
        Length: distance,
    }
    city.Routes[target] = route
    return route
}

func (city *City) ShortName() string {
    return city.Name
}

