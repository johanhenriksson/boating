package orders

import (
    "os"
    "fmt"
    "strings"
    "io/ioutil"
)

type ParserFunction func(int64, []string) (Order, error)

func CompileFile(filename string) Orders {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        return Orders { }
    }

    return Compile(string(bytes))
}

func Compile(code string) Orders {
    code = strings.Replace(code, "\n", ";", -1)
    lines := strings.Split(code, ";")
    orders := make(Orders, 0)

    for i, line := range lines {
        line = strings.Trim(line, " ")
        if len(line) == 0 || line[0] == '#' {
            continue
        }
        line = strings.ToLower(line)
        tokens := strings.Split(line, " ")
        command := tokens[0]
        tokens = tokens[1:]

        if parser, ok := parseTable[command]; ok {
            order, err := parser(int64(i+1), tokens)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
            orders = append(orders, order)
        } else {
            fmt.Println("Invalid order:", command)
            os.Exit(1)
        }
    }

    return orders
}
