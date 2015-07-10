package core

type Offer struct {
    Item       *Commodity
    Payment    *Commodity
    Units      int64
    Price      int64
    Count      int64
    Infinite   bool  /* Request an inifinite amount */
}

type Market struct {
    Offers      []*Offer
}

func NewMarket() *Market {
    market := &Market {
        Offers: []*Offer {
            &Offer {
                Item: GOLD,
                Units: 10,
                Payment: FOOD,
                Price: 150,
                Count: -1,
                Infinite: true,
            },
        },
    }
    go MarketWorker(market)
    return market
}

func MarketWorker(market *Market) {
    for {
    }
}

/*
Buy 150 steel for 1 gold
    Request: STEEL
    Units: 150
    Payment: GOLD
    Quantity: 1

Offer 150 steel for 1 gold
    Request: GOLD
    Units: 1
    Payment: STEEL
    Quantity: 150

*/
