package core

import (
    "time"
)

var TIMESCALE   time.Duration = 3600 * 2
var START_TIME  time.Time = time.Now()
var EPOCH       time.Time = time.Date(1852, time.January, 1, 0, 0, 0, 0, time.UTC)

/* Timescale-aware sleep function */
func sleep(duration time.Duration) {
    time.Sleep(duration / TIMESCALE)
}

/* Generates a GUID */
var __next_id int64 = 0
func nextId() int64 {
    __next_id++
    return __next_id
}
