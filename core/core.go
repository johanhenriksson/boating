package core

var __next_id int64 = 0
func nextId() int64 {
    __next_id++
    return __next_id
}
