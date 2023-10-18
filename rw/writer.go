package rw

import (
        "sync"
        "time"
        "fmt"
)

var MAX_WRITES = 3

type Writer struct {
        id int
        shared *Shared
        key string
        writes int
        wg *sync.WaitGroup
}

func NewWriter(uid int, key string, sd *Shared, wg *sync.WaitGroup) Writer {
        return Writer{id: uid, key: key, shared: sd, wg: wg}
}

func (w *Writer) Write(key string, value int) {
        w.shared.Write(key, value)
        w.writes += 1
}

func (w *Writer) Incr(key string) {
        w.shared.Incr(key)
        w.writes += 1
}

func (w* Writer) Loop(messages chan string) {
        fmt.Println("In writer loop")
        defer w.wg.Done()

        for {
                v := w.Incr(w.key)
                msg := "Incr key value" + w.key + "value is" + v + "w" + w.id
                messages <- msg

                if w.writes == MAX_WRITES {
                        return
                }

                time.Sleep(50)
        }
}
