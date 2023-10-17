package rw

import (
	"fmt"
)

type Data struct {
        m map[string]int
        read_count int
}

func (s *Data) New() {
        s.m = make(map[string]int)
}

func (s *Data) Read(key string) int {
        v := s.m[key]
        fmt.Println("Read key", key, "value is", v)
        return v
}

func (s *Data) Write(key string, value int) {
        s.m[key] = value
        fmt.Println("Wrote key", key, "to value", value)
}

func (s *Data) Values() {
        for k, v := range s.m {
                fmt.Println("Key:", k, "Value:", v)
        }
}
