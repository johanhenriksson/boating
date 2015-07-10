package orders

import (
    "github.com/johanhenriksson/trade/core"
)

/* City translations for order scripting */
var cityMap = map[string]*core.City {
    "london":       core.LONDON,
    "amsterdam":    core.AMSTERDAM,
    "hamburg":      core.HAMBURG,
    "washington":   core.WASHINGTON,
}
