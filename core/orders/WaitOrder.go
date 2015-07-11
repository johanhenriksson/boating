package orders

import (
    "fmt"
    "time"
    "github.com/johanhenriksson/boating/core"
)

type WaitOrder struct {
    Duration    time.Duration
}

func (order *WaitOrder) Create() func() {
    duration := order.Duration
    return func() {
        core.Sleep(duration)
    }
}

func (order *WaitOrder) Print() {
    fmt.Println("Wait", order.Duration)
}
