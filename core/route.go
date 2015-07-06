package core

type Route struct {
    Id      int64
    From    *City
    To      *City
    Length  int64
    /* Type - ocean, railroad ... */
}
