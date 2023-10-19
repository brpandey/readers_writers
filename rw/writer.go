package rw

import (
        //        "time"
        "strconv"
)

var MAX_WRITES = 3

type Writer struct {
        id int
        shared *Shared
        key string
        writes int
}

func NewWriter(uid int, key string, sd *Shared) Writer {
        return Writer{id: uid, key: key, shared: sd}
}

func (w *Writer) Write(key string, value int) {
        w.shared.Write(key, value)
        w.writes += 1
}

func (w *Writer) Incr(key string) int {
        v := w.shared.Incr(key)
        w.writes += 1
        return v
}

func (w* Writer) Loop(messages chan string) {
        for {
                v := w.Incr(w.key)
                msg := "W" + strconv.Itoa(w.id) + " Incr " + w.key + "'s value by 1 => " + strconv.Itoa(v)
                messages <- msg

                if w.writes == MAX_WRITES {
                        messages <- "Q"
                        return
                }
        }
}
