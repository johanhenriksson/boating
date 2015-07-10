package core

import (
    "fmt"
)

type Distance int64

var METER   Distance  = 1
var KM      Distance  = 1000

func (d Distance) String() string {
    return fmt.Sprintf("%d m", d)
}
