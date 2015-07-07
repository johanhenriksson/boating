package core

import (
    "time"
)

type Building struct {
    City        *City
}

func BuildingWorker(b *Building) {
    for {

        time.Sleep(1 * time.Second)
    }
}
