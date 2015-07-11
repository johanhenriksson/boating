package core

import (
    "time"
    "sync"
)

var TIMESCALE   time.Duration = 3600 * 2
var START_TIME  time.Time = time.Now()
var EPOCH       time.Time = time.Date(1852, time.January, 1, 0, 0, 0, 0, time.UTC)

/* Timescale-aware sleep function */
func Sleep(duration time.Duration) {
    time.Sleep(duration / TIMESCALE)
}

func Time() time.Time {
    /* Changing timescale during game will fuck up dates */
    world_duration := time.Since(START_TIME) * TIMESCALE
    game_time      := EPOCH.Add(world_duration)
    return game_time
}

/* NextId state */
var _nextId   int64       = 0
var _nextLock *sync.Mutex = &sync.Mutex { }

/* Generates a GUID */
func nextId() int64 {
    _nextLock.Lock()
    defer _nextLock.Unlock()

    _nextId++
    return _nextId
}
