package rw

type Data struct {
        m map[string]int
}

func NewData() Data {
        var d Data 
        d.m = make(map[string]int)
        return d
}

func (d *Data) Read(key string) int {
        return d.m[key]
}

func (d *Data) Write(key string, value int) {
        d.m[key] = value
}

func (d *Data) Incr(key string) int {
        d.m[key]++
        return d.m[key]
}
