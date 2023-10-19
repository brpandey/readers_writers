package rw

type ReadWrite interface {
        Read(key string) int
        Write(key string, value int)
        Incr(key string) int
}


type MiniActor interface {
        Loop(messages chan string)
}
