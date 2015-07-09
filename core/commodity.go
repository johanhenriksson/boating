package core

import (
    "errors"
)

type ComType int64

var CommodityNotFoundError = errors.New("No such commodity in storage")
var NotEnoughItemsError    = errors.New("Not enough items in storage")

var GOLD = &Commodity {
    Type: 1,
    Name: "Gold",
    UnitWeight: 100 * GRAM,
}

var SILVER = &Commodity {
    Type: 2,
    Name: "Silver",
    UnitWeight: 1 * KG,
}

var COPPER = &Commodity {
    Type: 3,
    Name: "Copper",
    UnitWeight: 10 * KG,
}

var DIAMOND = &Commodity {
    Type: 4,
    Name: "Diamond",
    UnitWeight: 40000,
}

var STEEL = &Commodity {
    Type: 10,
    Name: "Steel",
    UnitWeight: 100 * KG,
}

var FOOD = &Commodity {
    Type: 99,
    Name: "Food",
    UnitWeight: 1 * TON,
}

var COFFEE = &Commodity {
    Type: 100,
    Name: "Coffee",
    UnitWeight: 1 * TON,
}

var WEAPONS = &Commodity {
    Type: 1337,
    Name: "GUNS",
    UnitWeight: 1 * TON,
}

type Commodity struct {
    Type        ComType
    Name        string
    UnitWeight  Weight /* Unit weight */
}

func (c *Commodity) Weight() Weight {
    return c.UnitWeight
}

type Crate struct {
    Id      int64
    Owner   *Player
    Type    *Commodity
    Qty     int64 /* Quantity (units) */
}

func (c *Crate) Weight() Weight {
    return Weight(c.Qty) * c.Type.Weight()
}

func (crate *Crate) Add(stuff *Crate) {
    if crate.Type != stuff.Type {
        /* Throw error? */
        return
    }
    crate.Qty += stuff.Qty
    stuff.Qty = 0
}

func (crate *Crate) Take(quantity int64) (*Crate, error) {
    if quantity > crate.Qty {
        return nil, NotEnoughItemsError
    }
    crate.Qty -= quantity
    return GetCrate(crate.Owner, crate.Type, quantity), nil
}

func (crate *Crate) Split() (*Crate, *Crate) {
    /* This method is just ridiculously awesome */
    amount_a := crate.Qty / 2
    amount_b := crate.Qty - amount_a
    crate.Qty = 0

    a := GetCrate(crate.Owner, crate.Type, amount_a)
    b := GetCrate(crate.Owner, crate.Type, amount_b)
    return a, b
}

func GetCrate(owner *Player, com *Commodity, quantity int64) *Crate {
    return &Crate {
        Id:    nextId(),
        Owner: owner,
        Qty:   quantity,
        Type:  com,
    }
}
