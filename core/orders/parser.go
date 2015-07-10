package orders

import (
    "fmt"
    "errors"
    "strconv"
)

var parseTable = map[string]ParserFunction {
    "go": func(line int64, tokens []string) (Order, error) {
        if len(tokens) != 1 {
            return nil, errors.New(fmt.Sprintf("%d: Too many arguments to Go", line))
        }
        city := cityMap[tokens[0]]
        return &GoOrder {
            City: city,
        }, nil
    },
    "load": func(line int64, tokens []string) (Order, error) {
        return parseLoad(false, line, tokens)
    },
    "unload": func(line int64, tokens []string) (Order, error) {
        return parseLoad(true, line, tokens)
    },
}


func parseLoad(unload bool, line int64, tokens []string) (Order, error) {
    if len(tokens) == 0 {
        return nil, errors.New(fmt.Sprintf("%d: Too few arguments to Load", line))
    }
    if len(tokens) > 2 {
        return nil, errors.New(fmt.Sprintf("%d: Too many arguments to Load", line))
    }

    all := true
    var qty int64 = 0
    if len(tokens) == 2 {
        /* get count */
        var err error
        qty, err = strconv.ParseInt(tokens[1], 10, 64)
        if err != nil {
            return nil, errors.New(fmt.Sprintf("%d: Expected integer value (argument 2)", line))
        }

        all = false
    }

    com, ok := commodityMap[tokens[0]]
    if !ok {
        return nil, errors.New(fmt.Sprintf("No such commodity '%s'", tokens[0]))
    }

    return &LoadOrder {
        Commodity: com,
        Unload: unload,
        All: all,
        Quantity: qty,
    }, nil
}
