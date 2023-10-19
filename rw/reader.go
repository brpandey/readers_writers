package rw

import (
        //        "time"
        "strconv"
)

var MAX_READS = 4

type Reader struct {
        id int
        shared *Shared
        reads int
        key string
}

func NewReader(uid int, key string, sd *Shared) Reader {
        return Reader{id: uid, key: key, shared: sd}
}

func (r* Reader) Read(key string) int {
        v := r.shared.Read(key)
        r.reads += 1
        return v
}

func (r* Reader) Loop(messages chan string) {
        for {
                v := r.Read(r.key)
                msg := "R" + strconv.Itoa(r.id) + " Read " + r.key + " -> " + strconv.Itoa(v)
                messages <- msg

                if r.reads == MAX_READS {
                        messages <- "Q"
                        return
                }
        }
}
