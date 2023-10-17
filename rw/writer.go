package rw

import (
        "sync"
)

type Writer struct {
        id int
}

func NewWriter(uid int) Writer {
        return Writer{id: uid}
}

func (w *Writer) Write(read_write *sync.Mutex, shared_data *Data) {
        read_write.Lock()

        v := shared_data.Read("purple")
        shared_data.Write("purple", v + 1)

        read_write.Unlock()
}
