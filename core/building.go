package core

import (
    "time"
)

type Building struct {
    Id          int64
    Type        int64
    Name        string
    City        *City
    Stock       *Storage
    Owner       *Player
}

func BuildingWorker(b *Building) {
    for {

        time.Sleep(1 * time.Second)
    }
}
