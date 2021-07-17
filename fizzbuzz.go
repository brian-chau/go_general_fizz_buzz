package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewScanner(os.Stdin)
    fmt.Print("> ")
    reader.Scan()
    num_args,err := strconv.Atoi(reader.Text())
    if err != nil {
        log.Fatalf("%v", err)
    }

    divisors := map[int]string{}
    for i := 0; i < num_args; i++ {
        fmt.Print("> ")
        reader.Scan()
        input := strings.Fields(reader.Text())
        divisor, err := strconv.Atoi(input[0])
        if err != nil {
            log.Fatalf("%v", err)
        }
        keyword := input[1]

        divisors[divisor] = keyword
    }

    fmt.Println(divisors)

    for i := 1; i < 20; i++ {
        found := false
        for key, val := range divisors {
            if i % key == 0 {
                fmt.Print(val)
                found = true
            }
        }
        if !found {
            fmt.Print(i)
        }
        fmt.Println()
    }
}