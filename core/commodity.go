package core

import (
    "errors"
)

type ComType int64

/* Commodity Definitions */

var GOLD        = NewCommodity(1,   "Gold",       100 * GRAM)
var SILVER      = NewCommodity(2,   "Silver",     1   * KG)
var COPPER      = NewCommodity(3,   "Copper",     10  * KG)
var DIAMOND     = NewCommodity(4,   "Diamond",    10  * GRAM)
var COAL        = NewCommodity(5,   "Coal",       1   * TON)
var STEEL       = NewCommodity(10,  "Steel",      100 * KG)
var WOOD        = NewCommodity(11,  "Wood",       100 * KG)
var FOOD        = NewCommodity(99,  "Food",       1   * TON)
var COFFEE      = NewCommodity(100, "Coffee",     1   * TON)
var LEMONS      = NewCommodity(101, "Lemons",     1   * TON)
var BEER        = NewCommodity(102, "Beer",       100 * KG)
var GARBAGE     = NewCommodity(400, "Garbage",    1   * TON)
var MAIL        = NewCommodity(500, "Mail",       1   * TON)
var EXPLOSIVES  = NewCommodity(300, "Explosives", 1   * TON)

/* Commodity Errors */

var CommodityNotFoundError = errors.New("No such commodity in storage")
var NotEnoughItemsError    = errors.New("Not enough items in storage")


type Tradable interface {
    Is(ComType) bool
    Type() ComType
}

type Commodity struct {
    Name            string
    UnitWeight      Weight
    ctype           ComType
}

func (c *Commodity) Type() ComType {
    return c.ctype
}

func (c *Commodity) Is(ctype ComType) bool {
    return c.Type() == ctype
}

/** Returns the unit weight of this commodity */
func (c *Commodity) Weight() Weight {
    return c.UnitWeight
}

func NewCommodity(ctype ComType, name string, weight Weight) *Commodity {
    return &Commodity {
        Name: name,
        UnitWeight: weight,
        ctype: ctype,
    }
}
