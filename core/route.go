package core

type Route struct {
    Id      int64
    From    *City
    To      *City
    Length  Distance
    /* Type - ocean, railroad ... */
}
