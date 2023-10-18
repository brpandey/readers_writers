package rw

type ReadWrite interface {
        Read(key string, uid int) int
        Write(key string, value int, uid int)
        Incr(key string, uid int) int
}


type MiniActor interface {
        Loop(messages chan string)
}
