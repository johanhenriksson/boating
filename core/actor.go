package core

type ActorType int64

var ACTOR           ActorType = 0
var ACTOR_VEHICLE   ActorType = 1
var ACTOR_STORAGE   ActorType = 10
var ACTOR_GARAGE    ActorType = 11

type Actor struct {
    orders chan func()
    stop   chan int
    loop   bool
}

type Orderable interface {
    Issue(func())
    Stop()
}

func NewActor() *Actor {
    actor := &Actor {
        orders: make(chan func()),
        stop:   make(chan int),
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
        case order := <-a.orders:
            orderQueue = append(orderQueue, order)

        case <-a.stop:
            /* Stop any running loop (after next order) and clear
               the order queue */
            a.loop = false
            orderQueue = nil

        case <-orderDone:
            executing = false
        }
    }
}

/* Queue an order */
func (a *Actor) Issue(order func()) {
    a.orders <- func() {
        a.loop = false
        order()
    }
}

func (a *Actor) Loop(order func()) {
    a.orders <- func() {
        a.loop = true
        for a.loop {
            order()
        }
    }
}

func (a *Actor) Stop() {
    a.stop <- 1
}

func (a *Actor) Type() ActorType {
    return ACTOR
}
