package rw

import (
        "sync"
)

// Shared Data wraps Data struct with concurrency support for multiple readers and single writer
type Shared struct {
        read_count int
        data *ReadWrite
        read_only sync.Mutex
        read_write sync.Mutex
}

func NewShared(d *ReadWrite) Shared {
        return Shared{data: d}
}

func (s *Shared) Read(key string) int {
        s.read_only.Lock()

        // all readers must obtain read only lock first before they can update the read_count
        s.read_count += 1

        if s.read_count == 1 {
                s.read_write.Lock() // since reader is first to read, ensure no writer is currently writing
        }

        s.read_only.Unlock() // let another reader access read only as there can be multiple readers at same time

        v := (*s.data).Read(key)

        s.read_only.Lock() // obtain read only lock to decrement read count and later unlock
        s.read_count -= 1
        s.read_only.Unlock()

        if s.read_count == 0 {
                s.read_write.Unlock() // this wakes up any writers that were waiting on this lock
        }

        return v
}

func (s *Shared) Incr(key string) int {
        s.read_write.Lock()
        v := (*s.data).Incr(key)
        s.read_write.Unlock()
        return v
}

func (s *Shared) Write(key string, value int) {
        s.read_write.Lock()
        (*s.data).Write(key, value)
        s.read_write.Unlock()
}
