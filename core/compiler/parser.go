package compiler

import (
    "fmt"
    "time"
    "errors"
    "strings"
    "strconv"
    "github.com/johanhenriksson/boating/core"
    "github.com/johanhenriksson/boating/core/orders"
)

var parseTable = map[string]ParserFunction {
    "go": parseGo,
    "wait": parseWait,
    "load": func(line int64, tokens []string) (orders.Order, error) {
        return parseLoad(false, line, tokens)
    },
    "unload": func(line int64, tokens []string) (orders.Order, error) {
        return parseLoad(true, line, tokens)
    },
}

func parseGo(line int64, tokens []string) (orders.Order, error) {
    if len(tokens) != 1 {
        return nil, errors.New(fmt.Sprintf("%d: Too many arguments to Go", line))
    }
    city := cityMap[tokens[0]]
    return &orders.GoOrder {
        City: city,
    }, nil
}

func parseWait(line int64, tokens []string) (orders.Order, error) {
    if len(tokens) < 1 {
        return nil, errors.New(fmt.Sprintf("%d: Too few arguments to Wait", line))
    }

    d, err := strconv.ParseInt(tokens[0], 10, 64)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("%d: Wait error: invalid duration", line))
    }

    duration := time.Duration(d)

    if len(tokens) == 2 {
        switch strings.ToLower(tokens[1]) {
        case "year":
            duration *= 365 * 24 * time.Hour
        case "years":
            duration *= 365 * 24 * time.Hour
        case "month":
            duration *= 30 * 24 * time.Hour
        case "months":
            duration *= 30 * 24 * time.Hour
        case "day":
            duration *= 24 * time.Hour
        case "days":
            duration *= 24 * time.Hour
        case "hour":
            duration *= time.Hour
        case "hours":
            duration *= time.Hour
        case "minute":
            duration *= time.Minute
        case "minutes":
            duration *= time.Minute
        case "second":
        case "seconds":
        default:
            return nil, errors.New(fmt.Sprintf("%d: Wait error: Invalid duration unit '%s'", line, tokens[1]))
        }
    }

    return &orders.WaitOrder {
        Duration: duration,
    }, nil
}

func parseLoad(unload bool, line int64, tokens []string) (orders.Order, error) {
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

    return &orders.LoadOrder {
        Commodity: com,
        Unload: unload,
        All: all,
        Quantity: qty,
    }, nil
}

