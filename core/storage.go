/* Provides a way of storing items in cities, buildings or vehicles concurrently */
package core

import (
    "errors"
)

var NoPlayerItemsError    = errors.New("Player does not have any stored items")
var StorageDestroyedError = errors.New("Storage container has been destroyed")

/* Base data structure - each player has a map if items currently in storage */
type StorageMap map[*Player]map[*Commodity]*Crate

/* Represents the storage itself. */
type Storage struct {
    Crates      StorageMap
    load        chan StoreMessage
    unload      chan GetMessage
    destroy     chan bool
}

/* Sent to the storage worker to store a crate of goods */
type StoreMessage struct {
    Crate       *Crate
    Result      chan StoreResult
}

/* Send to the storage worker to retrieve goods */
type GetMessage struct {
    Owner       *Player
    Commodity   *Commodity
    Quantity    int64
    Result      chan StoreResult
}

/* Returned from the two storage operations */
type StoreResult struct {
    Crate       *Crate
    err         error
}

/* Initializes a new Storage including a worker goroutine */
func NewStorage() *Storage {
    stock := &Storage {
        Crates:  make(StorageMap),
        load:    make(chan StoreMessage),
        unload:  make(chan GetMessage),
        destroy: nil,
    }
    go storageWorker(stock)
    return stock
}

/* Worker goroutine. Listens for requests to store or retrieve
   goods from the container. */
func storageWorker(stock *Storage) {
    for {
        select {
        /* Request to store a crate */
        case storeMsg := <-stock.load:
            crate, err := stock.put(storeMsg.Crate)
            storeMsg.Result <- StoreResult {
                Crate: crate,
                err: err,
            }
        /* Request to retrieve goods */
        case getMsg := <-stock.unload:
            crate, err := stock.get(getMsg.Owner, getMsg.Commodity, getMsg.Quantity)
            getMsg.Result <- StoreResult {
                Crate: crate,
                err: err,
            }
        /* Stop executing */
        case <-stock.destroy:
            return
        }
    }
}

/* Provides an interface for safely storing items in the storage container. 
   Will block the caller until a result is available */
func (stock *Storage) Store(crate *Crate) (*Crate, error) {
    /* Make sure the container hasn't been destroyed */
    if stock.destroy != nil {
        return nil, StorageDestroyedError
    }

    /* Send a Store message to the worker goroutine */
    result := make(chan StoreResult)
    stock.load <- StoreMessage {
        Crate: crate,
        Result: result,
    }

    /* Block until the result is available */
    r := <-result
    return r.Crate, r.err
}

/* Provides an interface for safely retrieving items from the storage container. 
   Will block the caller until a result is available */
func (stock *Storage) Get(owner *Player, com *Commodity, qty int64) (*Crate, error) {
    /* Make sure the container hasn't been destroyed */
    if stock.destroy != nil {
        return nil, StorageDestroyedError
    }

    /* Send a Get message to the worker goroutine */
    result := make(chan StoreResult)
    stock.unload <- GetMessage {
        Owner: owner,
        Commodity: com,
        Quantity: qty,
        Result: result,
    }

    /* Block until the result is available */
    r := <-result
    return r.Crate, r.err
}

/* Destroys this storage container. Stops the worker goroutine */
func (stock *Storage) Destroy() {
    /* Clean out storage map */
    stock.Crates  = make(StorageMap, 0)

    /* Send value on the destroy channel to notify worker */
    stock.destroy = make(chan bool)
    stock.destroy <- true
}

/* Private methods */

/* Retrieves goods from the storage container. Not thread safe */
func (stock Storage) get(owner *Player, com *Commodity, quantity int64) (*Crate, error) {
    /* Check if player has any goods stored */
    if crates, ok := stock.Crates[owner]; ok {
        /* Check if player has the desired commodity... */
        if crate, ok := crates[com]; ok {
            /* We have the exact amount! Just return that crate and remove it */
            if crate.Qty == quantity {
                delete(crates, com)
                return crate, nil
            }

            /* We have more than required. Split the crate */
            return crate.Take(quantity)
        }
        return nil, CommodityNotFoundError
    }
    return nil, NoPlayerItemsError
}

/* Stores goods in the storage container. Not thread safe */
func (stock *Storage) put(crate *Crate) (*Crate, error) {
    owner  := crate.Owner

    /* Attempt to get crates from the player map */
    crates := stock.Crates[owner]
    if crates == nil {
        /* Player-crate map doesn't exist yet. Create it */
        crates = make(map[*Commodity]*Crate)
        stock.Crates[owner] = crates

        /* Store a new crate with the contents */
        newCrate, _ := crate.Take(crate.Qty)
        crates[crate.Type] = newCrate
        return crate, nil
    }

    /* Check if we already have a crate with this commodity type */
    if existing, ok := crates[crate.Type]; ok {
        existing.Add(crate)
        return existing, nil
    } else {
        /* Nope, store a new crate */
        newCrate, _ := crate.Take(crate.Qty)
        crates[crate.Type] = newCrate
        return newCrate, nil
    }
}
