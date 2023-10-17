package rw

import (
        "sync"
)

type Reader struct {
        id int
}

func NewReader(uid int) Reader {
        return Reader{id: uid}
}

func (r* Reader) Read(read_only *sync.Mutex, read_write *sync.Mutex, shared_data *Data) {
        read_only.Lock()

        // all readers must obtain read only lock first before they can update the read_count
        shared_data.read_count += 1

        if shared_data.read_count == 1 {
                read_write.Lock() // since reader is first to read, ensure no writer is currently writing
        }

        read_only.Unlock() // let another reader access read only as there can be multiple readers at same time

        // at this point this reader or a reader has acquired the read_write lock so its safe to read w/o concern for data races
        shared_data.Read("purple")

        read_only.Lock() // obtain read only lock to decrement read count and later unlock
        shared_data.read_count -= 1
        read_only.Unlock()

        if shared_data.read_count == 0 {
                read_write.Unlock() // this wakes up any writers that were waiting on this lock
        }

}
