package core

type Storage map[*Player]map[*Commodity]*Crate

func (stock Storage) Store(crate *Crate) *Crate {
    owner  := crate.Owner
    crates := stock[owner]

    if crates == nil {
        crates = make(map[*Commodity]*Crate)
        crates[crate.Type] = crate
        stock[owner] = crates
        return crate
    }

    if existing, ok := crates[crate.Type]; ok {
        existing.Qty += crate.Qty
        return existing
    } else {
        crates[crate.Type] = crate
        return crate
    }
}

func (stock Storage) Get(owner *Player, com *Commodity, quantity int64) *Crate {
    if crates, ok := stock[owner]; ok {
        if crate, ok := crates[com]; ok {
            if crate.Qty >= quantity {
                crate.Qty -= quantity
                return GetCrate(owner, com, quantity)
            }
        }
    }
    return nil
}
