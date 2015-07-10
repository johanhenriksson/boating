package orders

import (
    "os"
    "fmt"
    "strings"
)

type ParserFunction func(int64, []string) (Order, error)

func Compile(orders string) Orders {
    lines := strings.Split(orders, ";")
    output := make(Orders, 0)
    for i, line := range lines {
        tokens := strings.Split(strings.Trim(strings.ToLower(line), " "), " ")
        command := tokens[0]
        tokens = tokens[1:]

        if parser, ok := parseTable[command]; ok {
            order, err := parser(int64(i), tokens)
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
            output = append(output, order)
        }
    }
    output.Print()
    return output
}
