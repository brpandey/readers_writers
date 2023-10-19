package main 

import (
        "fmt"
        "github.com/brpandey/readers_writers/rw"
)

func main() {
        d := rw.NewData()
        var read_write rw.ReadWrite = &d
        sd := rw.NewShared(&read_write)

        // keys used to read out of shared data structure
        keys := []string{"caramel", "purple", "obsidian", "graphite", "atlantis"}
        var roles []rw.MiniActor // mini actors that either read or write

        messages := make(chan string)

        // add readers
        for i := 0; i < len(keys); i++ {
                r := rw.NewReader(i, keys[i], &sd)
                roles = append(roles, &r)
        }

        // add writers
        for i := 0; i < 2; i++ {
                w := rw.NewWriter(i, keys[i], &sd)
                roles = append(roles, &w)
        }

        // spawn go routines for each of the roles
        for i := 0; i < len(roles); i++ {
                go roles[i].Loop(messages)
        }

        fmt.Println("\nReader (R) and Writer (W) interplay:")

        var msg string
        finished := len(roles)
        var count int

Outer:
        for {
                select {
                case msg = <- messages:
                        if msg == "Q" {
                                count += 1
                        } else {
                                fmt.Println("~~~", msg)
                        }
                default:
                        if count == finished {
                                break Outer
                        }
                }
        }

        fmt.Println("")
}
