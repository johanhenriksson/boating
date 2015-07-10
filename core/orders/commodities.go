package orders

import (
    "github.com/johanhenriksson/trade/core"
)

/* Commodity translations for order scripting */
var commodityMap = map[string]*core.Commodity {
    "gold":     core.GOLD,
    "silver":   core.SILVER,
    "steel":    core.STEEL,
    "coffee":   core.COFFEE,
}
