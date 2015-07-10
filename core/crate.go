package core

type Crate struct {
    Id      int64
    Owner   *Player
    Type    *Commodity
    Qty     int64 /* Quantity (units) */
}

func NewCrate(owner *Player, com *Commodity, quantity int64) *Crate {
    return &Crate {
        Id:    nextId(),
        Owner: owner,
        Qty:   quantity,
        Type:  com,
    }
}

func (c *Crate) Weight() Weight {
    return Weight(c.Qty) * c.Type.Weight()
}

func (crate *Crate) Add(stuff *Crate) {
    if crate.Type != stuff.Type {
        /* Throw error? */
        return
    }
    crate.Qty += stuff.Qty
    stuff.Qty = 0
}

func (crate *Crate) Take(quantity int64) (*Crate, error) {
    if quantity > crate.Qty {
        return nil, NotEnoughItemsError
    }
    crate.Qty -= quantity
    return NewCrate(crate.Owner, crate.Type, quantity), nil
}

func (crate *Crate) Split() (*Crate, *Crate) {
    /* This method is just ridiculously awesome */
    amount_a := crate.Qty / 2
    amount_b := crate.Qty - amount_a
    crate.Qty = 0

    a := NewCrate(crate.Owner, crate.Type, amount_a)
    b := NewCrate(crate.Owner, crate.Type, amount_b)
    return a, b
}

