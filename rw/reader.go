package rw

import (
        "sync"
        "time"
        "fmt"
)

var MAX_READS = 4

type Reader struct {
        id int
        shared *Shared
        reads int
        key string
        wg *sync.WaitGroup
}

func NewReader(uid int, key string, sd *Shared, wg *sync.WaitGroup) Reader {
        return Reader{id: uid, key: key, shared: sd, wg: wg}
}

func (r* Reader) Read(key string) {
        r.shared.Read(key)
        r.reads += 1
}

func (r* Reader) Loop(messages chan string) {
        fmt.Println("In reader loop")
        defer r.wg.Done()

        for {
                v := r.Read(r.key)
                msg := "Read key" + r.key + "value is" + v + "r" + r.id
                messages <- msg

                if r.reads == MAX_READS {
                        return
                }

                time.Sleep(50)
        }
}
