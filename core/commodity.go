package core

import (
    "errors"
)

type ComType int64

var CommodityNotFoundError = errors.New("No such commodity in storage")
var NotEnoughItemsError    = errors.New("Not enough items in storage")

type Commodity struct {
    Type            ComType
    Name            string
    UnitWeight      Weight
}

/** Returns the unit weight of this commodity */
func (c *Commodity) Weight() Weight {
    return c.UnitWeight
}

/* Commodity Definitions */

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
