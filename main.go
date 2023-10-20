package main 

import (
        "fmt"
        "flag"
        "strings"
        "time"
        "math/rand"
        "github.com/brpandey/readers_writers/rw"
)

const MAX_ACTORS int = 7

func main() {
        starve := flag.Bool("starve", false, "indicates whether to classically starve writers")
        nreaders := flag.Int("nr", 4, "specify num readers, max is 7")
        nwriters := flag.Int("nw", 1, "specify num writers, max is 7")

        flag.Parse()

        if *nreaders > MAX_ACTORS {
                *nreaders = MAX_ACTORS
        }

        if *nwriters > MAX_ACTORS {
                *nwriters = MAX_ACTORS
        }

        fmt.Println("Specified options: -starve:", *starve, " -nr:", *nreaders, " -nw:", *nwriters)

        d := rw.NewData()
        var read_write rw.ReadWrite = &d
        sd := rw.NewShared(&read_write, !(*starve))

        // keys used to read out of shared data structure
        keys := [MAX_ACTORS]string{"caramel", "purple", "obsidian", "graphite", "atlantis", "hexagon", "ascetic"}
        var roles []rw.MiniActor // mini actors that either read or write

        messages := make(chan string)

        // Add readers and writers based on default or cmd line options
        for i := 0; i < *nreaders; i++ {
                r := rw.NewReader(i, keys[i], &sd)
                roles = append(roles, &r)
        }
        for i := 0; i < *nwriters; i++ {
                w := rw.NewWriter(i, keys[i], &sd)
                roles = append(roles, &w)
        }

        // shuffle array so readers aren't always at beginning or started first
        rand.Seed(time.Now().UnixNano())
        rand.Shuffle(len(roles), func(i, j int) { roles[i], roles[j] = roles[j], roles[i] })

        var label string

        // Format spawn list of roles
        for i := 0; i < len(roles); i++ {
                label += fmt.Sprintf("%T", roles[i])
        }

        label = strings.Replace(label, "*rw.Reader", "R ", -1)
        label = strings.Replace(label, "*rw.Writer", "W ", -1)

        fmt.Printf("Spawn order: [ %s]\n", label)

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
