package main 

import (
        "sync"
        "fmt"
        "github.com/brpandey/readers_writers/rw"
)

func main() {
        d := rw.NewData()
        var read_write rw.ReadWrite = &d
        sd := rw.NewShared(&read_write)
        var wg sync.WaitGroup

        keys := []string{"caramel", "purple", "obsidian"}
        var roles []rw.MiniActor

        messages := make(chan string)

        // add readers
        for i := 0; i < len(keys); i++ {
                wg.Add(1)
                r := rw.NewReader(i, keys[i], &sd, &wg)
                roles = append(roles, &r)
        }

        // add writers
        for i := 0; i < 2; i++ {
                wg.Add(1)
                w := rw.NewWriter(i, keys[i], &sd, &wg)
                roles = append(roles, &w)
        }

        fmt.Println("roles", roles)

        for i := 0; i < len(roles); i++ {
                fmt.Println("About to spawn go routine", i)
                actor := roles[i]
                fmt.Println("actor is", actor)

                go actor.Loop(messages)
        }

        for {
                select {
                case msg <- messages:
                        fmt.Println("~~~", msg)
                }
        }

        wg.Wait() // block until wg value down to 0
}
