package core

import (
    "sync"
)

type VehicleMap  map[VehicleId]*Vehicle
type VehicleChan chan *Vehicle

type Garage struct {
    *Actor
    Vehicles VehicleMap
    lock     *sync.RWMutex
}

func NewGarage() *Garage {
    garage := &Garage {
        Actor:    NewActor(),
        Vehicles: make(VehicleMap),
        lock:     &sync.RWMutex { },
    }
    return garage
}

/* Add vehicle to garage. Blocks caller until done */
func (g *Garage) Add(vehicle *Vehicle) {
    g.park(vehicle)
}

func (g *Garage) Park(vehicle *Vehicle) {
    g.Issue(func() {
        g.park(vehicle)
    })
}

func (g *Garage) Unpark(vehicle *Vehicle) {
    g.Issue(func() {
        g.unpark(vehicle)
    })
}

/* Returns true if given vehicle is parked in this city */
func (g *Garage) Stores(vehicle *Vehicle) bool {
    g.lock.RLock()
    defer g.lock.RUnlock()

    _, exists := g.Vehicles[vehicle.Id]
    return exists
}

/* Parks a vehicle in the city */
func (g *Garage) park(vehicle *Vehicle) {
    g.lock.Lock()
    defer g.lock.Unlock()

    g.Vehicles[vehicle.Id] = vehicle
}

/* Removes a vehicle from parking in the city */
func (g *Garage) unpark(vehicle *Vehicle) {
    g.lock.Lock()
    defer g.lock.Unlock()

    delete(g.Vehicles, vehicle.Id)
}

