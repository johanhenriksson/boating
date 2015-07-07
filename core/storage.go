package core

type StorageMap map[*Player]map[*Commodity]*Crate

type Storage struct {
    Crates  StorageMap
}

func NewStorage() *Storage {
    return &Storage {
        Crates: make(StorageMap),
    }
}

/* TODO: Concurrency */
func (stock *Storage) Store(crate *Crate) *Crate {
    owner  := crate.Owner
    crates := stock.Crates[owner]

    if crates == nil {
        crates = make(map[*Commodity]*Crate)
        crates[crate.Type] = crate
        stock.Crates[owner] = crates
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

func (stock Storage) Has(owner *Player, com *Commodity, quantity int64) bool {
    if crates, ok := stock.Crates[owner]; ok {
        if crate, ok := crates[com]; ok {
            return crate.Qty >= quantity
        }
    }
    return false
}

/* TODO: Concurrency */
func (stock Storage) Get(owner *Player, com *Commodity, quantity int64) *Crate {
    if crates, ok := stock.Crates[owner]; ok {
        if crate, ok := crates[com]; ok {
            if crate.Qty >= quantity {
                crate.Qty -= quantity
                if crate.Qty == 0 {
                    /* Remove empty crates */
                    delete(crates, com)
                }
                return GetCrate(owner, com, quantity)
            }
        }
    }
    return nil
}

/* TODO: Concurrency */
func (stock Storage) Remove(crate *Crate) *Crate {
    if crates, ok := stock.Crates[crate.Owner]; ok {
        crate, ok := crates[crate.Type]
        if ok {
            delete(crates, crate.Type)
            return crate
        }
    }
    return nil
}
