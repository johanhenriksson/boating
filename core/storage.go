package core

import (
    "errors"
)

var CommodityNotFoundError = errors.New("No such commodity in storage")
var NotEnoughItemsError    = errors.New("Not enough items in storage")
var NoPlayerItemsError     = errors.New("Player does not have any stored items")

type StorageMap map[*Player]map[*Commodity]*Crate

type Storage struct {
    Crates      StorageMap
    load        chan StoreMessage
    unload      chan GetMessage
}

type StoreMessage struct {
    Crate       *Crate
    Result      chan StoreResult
}

type GetMessage struct {
    Owner       *Player
    Commodity   *Commodity
    Quantity    int64
    Result      chan StoreResult
}

type StoreResult struct {
    Crate       *Crate
    err         error
}

func NewStorage() *Storage {
    stock := &Storage {
        Crates: make(StorageMap),
        load:   make(chan StoreMessage),
        unload: make(chan GetMessage),
    }
    go storageWorker(stock)
    return stock
}

func storageWorker(stock *Storage) {
    for {
        select {
        case storeMsg := <-stock.load:
            /* Store this crate */
            crate, err := stock.put(storeMsg.Crate)
            storeMsg.Result <- StoreResult {
                Crate: crate,
                err: err,
            }
        case getMsg := <-stock.unload:
            crate, err := stock.get(getMsg.Owner, getMsg.Commodity, getMsg.Quantity)
            getMsg.Result <- StoreResult {
                Crate: crate,
                err: err,
            }
        }
    }
}

func (stock *Storage) Store(crate *Crate) (*Crate, error){
    result := make(chan StoreResult)
    stock.load <- StoreMessage {
        Crate: crate,
        Result: result,
    }
    r := <-result
    return r.Crate, r.err
}

func (stock *Storage) Get(owner *Player, com *Commodity, qty int64) (*Crate, error) {
    result := make(chan StoreResult)
    stock.unload <- GetMessage {
        Owner: owner,
        Commodity: com,
        Quantity: qty,
        Result: result,
    }
    r := <-result
    return r.Crate, r.err
}

func (stock Storage) get(owner *Player, com *Commodity, quantity int64) (*Crate, error) {
    if crates, ok := stock.Crates[owner]; ok {
        if crate, ok := crates[com]; ok {
            if crate.Qty >= quantity {
                crate.Qty -= quantity
                if crate.Qty == 0 {
                    /* Remove empty crates */
                    delete(crates, com)
                }
                return GetCrate(owner, com, quantity), nil
            } else {
                return nil, NotEnoughItemsError
            }
        } else {
            return nil, CommodityNotFoundError
        }
    }
    return nil, NoPlayerItemsError
}

func (stock *Storage) put(crate *Crate) (*Crate, error) {
    owner  := crate.Owner
    crates := stock.Crates[owner]

    if crates == nil {
        crates = make(map[*Commodity]*Crate)
        crates[crate.Type] = crate
        stock.Crates[owner] = crates
        return crate, nil
    }

    if existing, ok := crates[crate.Type]; ok {
        existing.Qty += crate.Qty
        return existing, nil
    } else {
        crates[crate.Type] = crate
        return crate, nil
    }
}
