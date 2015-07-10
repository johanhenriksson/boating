package core

var LONDON      = NewCity(1, "London")
var AMSTERDAM   = NewCity(2, "Amsterdam")
var HAMBURG     = NewCity(3, "Hamburg")
var WASHINGTON  = NewCity(4, "Washington")

type CityMap map[CityId]*City

var Cities = CityMap {
    AMSTERDAM.Id:   AMSTERDAM,
    LONDON.Id:      LONDON,
    HAMBURG.Id:     HAMBURG,
    WASHINGTON.Id:  WASHINGTON,
}

func init() {
    /* Define routes */
    AMSTERDAM.addRoute(HAMBURG, 70)
    AMSTERDAM.addRoute(LONDON, 160)

    HAMBURG.addRoute(AMSTERDAM, 70)
    HAMBURG.addRoute(LONDON, 200)

    LONDON.addRoute(AMSTERDAM, 160)
    LONDON.addRoute(HAMBURG, 200)
}
