package compiler

import (
    "testing"
    "github.com/johanhenriksson/boating/core"
    "github.com/johanhenriksson/boating/core/orders"
)

func TestParseGo(t *testing.T) {
    order, err := parseGo(0, []string { "amsterdam", })
    if err != nil {
        t.Errorf("Parse Go error: %s", err)
    }
    goOrder, cast_ok := order.(*orders.GoOrder)
    if !cast_ok {
        t.Errorf("Expected return type to be GoOrder", err)
    }
    if goOrder.City != core.AMSTERDAM {
        t.Errorf("Expected city to be Amsterdam", err)
    }
}
