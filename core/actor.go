package core

type ActorType int64

var ACTOR           ActorType = 0
var ACTOR_VEHICLE   ActorType = 1
var ACTOR_STORAGE   ActorType = 10
var ACTOR_GARAGE    ActorType = 11

type Actor struct {
    Orders chan func()
}

func NewActor() *Actor {
    actor := &Actor {
        Orders: make(chan func()),
    }
    go ActorWorker(actor)
    return actor
}

func ActorWorker(a *Actor) {
    orderQueue := make([]func(), 0, 4)
    orderDone  := make(chan int)
    executing  := false

    for {
        /* if there is more work to do... */
        if !executing && len(orderQueue) > 0 {
            order := orderQueue[0]
            orderQueue = orderQueue[1:]
            executing = true
            /* Execute order on a separate thread to make sure the worker
               remains responsive while executing the order */
            go func() {
                order()
                orderDone <- 1
            }()
        }

        select {
        case order := <-a.Orders:
            orderQueue = append(orderQueue, order)
        case <-orderDone:
            executing = false
        }
    }
}

/* Queue an order */
func (a *Actor) Issue(order func()) {
    a.Orders <- order
}

func (a *Actor) Type() ActorType {
    return ACTOR
}
