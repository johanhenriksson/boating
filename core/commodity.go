package core

type ComType int64

var GOLD = &Commodity {
    Type: 1,
    Name: "Gold",
    Weight: 20000,
}

var COFFEE = &Commodity {
    Type: 100,
    Name: "Coffee",
    Weight: 20000,
}

type Commodity struct {
    Type    ComType
    Name    string
    Weight  int64 /* Unit weight */
}

type Crate struct {
    Id      int64
    Owner   *Player
    Type    *Commodity
    Qty     int64 /* Quantity (units) */
}

func (crate *Crate) Split() (*Crate, *Crate) {
    /* This method is just ridiculously awesome */
    amount_a := crate.Qty / 2
    amount_b := crate.Qty - amount_a
    crate.Qty = 0

    a := GetCrate(crate.Owner, crate.Type, amount_a)
    b := GetCrate(crate.Owner, crate.Type, amount_b)
    return a, b
}

func GetCrate(owner *Player, com *Commodity, quantity int64) *Crate {
    return &Crate {
        Id:    nextId(),
        Owner: owner,
        Qty:   quantity,
        Type:  com,
    }
}
