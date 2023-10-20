package rw

import (
        "sync"
)

// Shared Data wraps Data struct with concurrency support for multiple readers and single writer
type Shared struct {
        read_count int
        data *ReadWrite
        read_only sync.Mutex // allows multiple readers to quickly access if no writer waiting
        read_write sync.Mutex // coordination between read and write modes
        first_in_line sync.Mutex // intended to provide fair starve-free approach for writers, to make r and w are equal
        fair bool // if true don't starve writers else do
}

func NewShared(d *ReadWrite, fair bool) Shared {
        return Shared{data: d, fair: fair}
}

func (s *Shared) Read(key string) int {
        if s.fair {
                s.first_in_line.Lock() // if we are first in line whether reader or writer grab lock (here we are reader of course)
        }
        s.read_only.Lock() // grab read specific lock to update read count

        // all readers must obtain read only lock first before they can update the read_count
        s.read_count += 1

        if s.read_count == 1 {
                s.read_write.Lock() // since reader is first to read, ensure no writer is currently writing
        }

        s.read_only.Unlock() // let another reader access read only as there can be multiple readers at same time
        if s.fair {
                s.first_in_line.Unlock() // let the next r or w waiting, be the first in line by releasing lock
        }

        /* START CRITICAL SECTION */
        v := (*s.data).Read(key)
        /* END CRITICAL SECTION */

        s.read_only.Lock() // obtain read only lock to decrement read count and later unlock
        s.read_count -= 1
        s.read_only.Unlock()

        if s.read_count == 0 {
                s.read_write.Unlock() // this wakes up any writers that were waiting on this lock
        }

        return v
}

func (s *Shared) Incr(key string) int {
        if s.fair {
                s.first_in_line.Lock() // if we are first in line whether reader or writer grab lock
        }

        s.read_write.Lock()

        // since we have the exclusive r / w coordination lock now, don't need first in lock anymore, release it
        if s.fair {
                s.first_in_line.Unlock()
        }
        /* START CRITICAL SECTION */
        v := (*s.data).Incr(key)
        /* END CRITICAL SECTION */

        s.read_write.Unlock()
        return v
}

func (s *Shared) Write(key string, value int) {
        if s.fair {
                s.first_in_line.Lock() // if we are first in line whether reader or writer grab lock
        }
        s.read_write.Lock()

        // since we have the exclusive r / w coordination lock don't need first in lock anymore, release it
        if s.fair {
                s.first_in_line.Unlock()
        }

        /* START CRITICAL SECTION */
        (*s.data).Write(key, value)
        /* END CRITICAL SECTION */

        s.read_write.Unlock()
}
