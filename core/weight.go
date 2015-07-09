package core

import (
    "fmt"
)

type Weight int64

type Weighable interface {
    Weight() Weight
}

/* Units */

const GRAM  Weight = 1
const KG    Weight = 1000
const TON   Weight = 1000000

func (w Weight) Weight() Weight {
    return w
}

func (w Weight) String() string {
    if (w > TON) {
        return fmt.Sprintf("%.2f t", float32(w) / float32(TON))
    }
    if (w > KG) {
        return fmt.Sprintf("%.3f kg", float32(w) / float32(KG))
    }
    return fmt.Sprintf("%d g", w)
}
