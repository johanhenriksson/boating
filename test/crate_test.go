package test

import (
    "testing"
    "../core"
)

var testCom = &core.Commodity {
    Type: core.ComType(9001),
    Name: "Test Item",
    UnitWeight: 10 * core.KG,
}

var testPlayer = &core.Player {
    Name: "Test Player",
}

func TestNewCrate(t *testing.T) {
    var want_qty int64 = 50
    want_weight := core.Weight(want_qty) * testCom.UnitWeight

    c := core.NewCrate(testPlayer, testCom, want_qty)

    if c.Type != testCom {
        t.Errorf("Expected crate to be of type '%s', got '%s'", testCom.Name, c.Type.Name)
    }
    if c.Qty != want_qty {
        t.Errorf("Expected crate to contain %d items, got %d", want_qty, c.Qty)
    }
    if c.Owner != testPlayer {
        t.Errorf("Wrong owner")
    }
    if c.Weight() != want_weight {
        t.Errorf("Weight is incorrect")
    }
}

func TestNewCrateNegQty(t *testing.T) {
    c := core.NewCrate(testPlayer, testCom, -1)
    if c.Qty != 0 {
        t.Errorf("Negative quantities should probably be an error or at least zero")
    }
}

func TestCrateTake(t *testing.T) {
    qty, take := int64(100), int64(50)
    want_a, want_b := qty - take, take
    crate_a := core.NewCrate(testPlayer, testCom, qty)

    crate_b, err := crate_a.Take(take)

    if err != nil {
        t.Errorf("Crate error: %s", err)
    }
    if crate_a.Qty != want_a {
        t.Errorf("Wrong quantity in crate A. Was %d, expected %d", crate_a.Qty, want_a)
    }
    if crate_b.Qty != want_b {
        t.Errorf("Wrong quantity in crate B. Was %d, expected %d", crate_b.Qty, want_b)
    }
}

func TestCrateTakeMuch(t *testing.T) {
    qty, take := int64(100), int64(150)

    crate_a := core.NewCrate(testPlayer, testCom, qty)
    _, err := crate_a.Take(take)

    if err != core.NotEnoughItemsError {
        t.Errorf("Expected not enough items error")
    }
}
