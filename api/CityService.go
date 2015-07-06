package api

import (
    "fmt"
    "strconv"
    "encoding/json"
    "github.com/johanhenriksson/trade/core"
)

type CityService struct {
    World *core.World
}

func (srv CityService) Load() {
    /* get from app reference or something */
}

func (srv CityService) Path() string {
    return "/v1/{user}/city"
}

func (srv CityService) Routes() Routes {
    return Routes {
        Route {
            Name: "Get All Cities",
            Method: "GET",
            Pattern: "/",
            Handler: srv.GetAll,
        },
    }
}

func (srv *CityService) GetAll(p RouteArgs) {
    user_id, err := strconv.ParseInt(p.Vars["user"], 10, 64)
    if err != nil {
        fmt.Fprintf(p.Writer, "no such player %s", p.Vars["user"])
        return
    }
    player := srv.World.Players[user_id]

    response := make([]CityResponse, len(srv.World.Cities))
    i := 0
    for _, city := range srv.World.Cities {
        response[i] = CityResponse {
            Id: city.Id,
            Name: city.Name,
            Stock: make([]CrateResponse, 0),
        }

        /* Player stock */
        if crates, ok := city.Stock[player]; ok {
            for _, crate := range crates {
                response[i].Stock = append(response[i].Stock, CrateResponse {
                    Commodity: crate.Type.Name,
                    Quantity: crate.Qty,
                })
            }
        }

        i++
    }

    p.Writer.Header().Set("Content-Type", "application/json")
    json, err := json.Marshal(response)
    if err == nil {
        p.Writer.Write(json)
    }
}

type CityResponse struct {
    Id          int64
    Name        string
    Stock       []CrateResponse
}

type CrateResponse struct {
    Commodity   string
    Quantity    int64
}
