package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    file, err := os.Open("easy-input.txt")
    check(err)
    defer file.Close()

    var network =  make(map[int]int)

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        points := strings.Split(line, " ")
        if len(points) > 1 {
            origin, err := strconv.Atoi(points[0])
            check(err)

            endpoint, err := strconv.Atoi(points[1])
            check(err)

            network[origin]++
            network[endpoint]++
        }
    }

    var keys []int
    for k := range network {
        keys = append(keys, k)
    }

    sort.Ints(keys)

    for _, k := range keys {
        fmt.Printf("Node %v has a degree of %v\n", k, network[k])
    }
}
