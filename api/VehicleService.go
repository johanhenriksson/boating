package api

import (
    "fmt"
    "sort"
    "strconv"
    "encoding/json"
    "github.com/johanhenriksson/boating/core"
)

type VehicleService struct {
    World *core.WorldState
}

func (srv VehicleService) Load() {
    /* get from app reference or something */
}

func (srv VehicleService) Path() string {
    return "/v1/{user}/vehicle"
}

func (srv VehicleService) Routes() Routes {
    return Routes {
        Route {
            Name: "Get Vehicle",
            Method: "GET",
            Pattern: "/id/{id}/",
            Handler: srv.Get,
        },
        Route {
            Name: "Get All Vehicles",
            Method: "GET",
            Pattern: "/",
            Handler: srv.GetAll,
        },
    }
}

func (srv *VehicleService) GetAll(p RouteArgs) {
    user_id, err := strconv.ParseInt(p.Vars["user"], 10, 64)
    if err != nil {
        fmt.Fprintf(p.Writer, "no such player %s", p.Vars["user"])
        return
    }
    player := srv.World.Players[core.PlayerId(user_id)]

    response := make([]VehicleResponse, len(player.Vehicles))
    i := 0
    for _, v := range player.Vehicles {
        response[i] = NewVehicleResponse(v, player) 
        i++
    }

    p.Writer.Header().Set("Content-Type", "application/json")
    js, _ := json.Marshal(response)
    p.Writer.Write(js)
}

func (srv *VehicleService) Get(p RouteArgs) {
    user_id, err := strconv.ParseInt(p.Vars["user"], 10, 64)
    if err != nil {
        fmt.Fprintf(p.Writer, "no player id")
        return
    }
    player := srv.World.Players[core.PlayerId(user_id)]

    vehicle_id, err := strconv.ParseInt(p.Vars["vehicle"], 10, 64)
    if err != nil {
        fmt.Fprintf(p.Writer, "no vehicle id")
        return
    }
    exists, vehicle := core.World.Vehicles.Find(core.VehicleId(vehicle_id))
    if exists {
        response := NewVehicleResponse(vehicle, player)

        p.Writer.Header().Set("Content-Type", "application/json")
        js, _ := json.Marshal(response)
        p.Writer.Write(js)
    } else {
        fmt.Fprintf(p.Writer, "Vehicle %d not found", vehicle_id)
    }
}

type VehicleResponse struct {
    Id          int64               `json:"id"`
    Name        string              `json:"name"`
    Cargo       []CrateResponse     `json:"cargo"`
    Journey     JourneyResponse     `json:"journey"`
}

func NewVehicleResponse(v *core.Vehicle, player *core.Player) VehicleResponse {
    r := VehicleResponse {
        Id:      int64(v.Id),
        Name:    v.Name,
        Cargo:   make([]CrateResponse, 0),
        Journey: JourneyResponse {
            To:         v.Journey.To.Name,
            From:       v.Journey.From.Name,
            Distance:   int64(v.Journey.Distance),
            Remaining:  int64(v.Journey.Remaining),
            Start:      v.Journey.Start.Unix(),
        },
    }

    crates, ok := v.Cargo.Crates[player]
    if ok {
        for _, crate := range crates {
            r.Cargo = append(r.Cargo, CrateResponse {
                Type:       crate.Type.Type,
                Commodity:  crate.Type.Name,
                Quantity:   crate.Qty,
                Owner:      crate.Owner.Name,
                Weight:     crate.Weight().String(),
            })
        }

        sort.Sort(CrateByType(r.Cargo))
    }

    return r
}

type JourneyResponse struct {
    From        string              `json:"from"`
    To          string              `json:"to"`
    Distance    int64               `json:"distance"`
    Remaining   int64               `json:"remaining"`
    Start       int64               `json:"start"`
}
