package api

import (
    "fmt"
    "strconv"
    "encoding/json"
    "github.com/johanhenriksson/trade/core"
)

type VehicleService struct {
    World *core.World
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
    player := srv.World.Players[user_id]

    response := make([]VehicleResponse, len(player.Vehicles))
    i := 0
    for _, v := range player.Vehicles {
        response[i] = VehicleResponse {
            Id:      v.Id,
            Name:    v.Name,
            Cargo:   make([]CrateResponse, 0),
            Journey: JourneyResponse {
                To:         v.Journey.To.Name,
                From:       v.Journey.From.Name,
                Distance:   v.Journey.Distance,
                Remaining:  v.Journey.Remaining,
                Start:      v.Journey.Start.Unix(),
            },
        }

        crates, ok := v.Cargo.Crates[player]
        if ok {
            for _, crate := range crates {
                response[i].Cargo = append(response[i].Cargo, CrateResponse {
                    Type:       crate.Type.Type,
                    Commodity:  crate.Type.Name,
                    Quantity:   crate.Qty,
                    Owner:      crate.Owner.Name,
                })
            }
        }

        i++
    }

    p.Writer.Header().Set("Content-Type", "application/json")
    js, _ := json.Marshal(response)
    p.Writer.Write(js)
}

type VehicleResponse struct {
    Id          int64               `json:"id"`
    Name        string              `json:"name"`
    Cargo       []CrateResponse     `json:"cargo"`
    Journey     JourneyResponse     `json:"journey"`
}

type JourneyResponse struct {
    From        string              `json:"from"`
    To          string              `json:"to"`
    Distance    int64               `json:"distance"`
    Remaining   int64               `json:"remaining"`
    Start       int64               `json:"start"`
}
