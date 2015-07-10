package orders

import (
    "fmt"
    "errors"
    "strconv"
    "github.com/johanhenriksson/trade/core"
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
    if !unload && len(tokens) == 0 {
        return nil, errors.New(fmt.Sprintf("%d: Too few arguments", line))
    }
    if len(tokens) > 2 {
        return nil, errors.New(fmt.Sprintf("%d: Too many arguments", line))
    }

    all := true
    var ok bool
    var qty int64 = 0
    var com *core.Commodity = nil

    /* We have a commodity specifier */
    if len(tokens) > 0 {
        all = false

        com, ok = commodityMap[tokens[0]]
        if !ok {
            return nil, errors.New(fmt.Sprintf("No such commodity '%s'", tokens[0]))
        }

        /* Check for a specified quantity */
        if len(tokens) > 1 {
            /* get count */
            var err error
            qty, err = strconv.ParseInt(tokens[1], 10, 64)
            if err != nil {
                return nil, errors.New(fmt.Sprintf("%d: Expected integer value (argument 2)", line))
            }
        }
    }

    return &LoadOrder {
        Commodity: com,
        Unload: unload,
        All: all,
        Quantity: qty,
    }, nil
}
