package api

import (
    "fmt"
    "sort"
    "strconv"
    "encoding/json"
    "github.com/johanhenriksson/boating/core"
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
    player := srv.World.Players[core.PlayerId(user_id)]

    response := make([]CityResponse, len(core.Cities))
    i := 0
    for _, city := range core.Cities {
        response[i] = CityResponse {
            Id: int64(city.Id),
            Name: city.Name,
            Stock: make([]CrateResponse, 0),
        }

        /* Player stock */
        if crates, ok := city.Stock.Crates[player]; ok {
            for _, crate := range crates {
                response[i].Stock = append(response[i].Stock, CrateResponse {
                    Type:       crate.Type.Type,
                    Commodity:  crate.Type.Name,
                    Quantity:   crate.Qty,
                    Owner:      crate.Owner.Name,
                    Weight:     crate.Weight().String(),
                })
            }

            sort.Sort(CrateByType(response[i].Stock))
        }

        i++
    }

    /* Sort responses by City Id */
    sort.Sort(CityById(response))

    p.Writer.Header().Set("Content-Type", "application/json")
    json, err := json.Marshal(response)
    if err == nil {
        p.Writer.Write(json)
    }
}


type CityResponse struct {
    Id          int64               `json:"id"`
    Name        string              `json:"name"`
    Stock       []CrateResponse     `json:"stock"`
}

type CityById []CityResponse
func (c CityById) Len() int           { return len(c) }
func (c CityById) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c CityById) Less(i, j int) bool { return c[i].Id < c[j].Id }

type CrateResponse struct {
    Type        core.ComType        `json:"type"`
    Quantity    int64               `json:"quantity"`
    Commodity   string              `json:"commodity"`
    Owner       string              `json:"owner"`
    Weight      string              `json:"weight"`
}

type CrateByType []CrateResponse
func (c CrateByType) Len() int           { return len(c) }
func (c CrateByType) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c CrateByType) Less(i, j int) bool { return c[i].Type < c[j].Type }
