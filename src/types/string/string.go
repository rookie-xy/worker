package string

import "github.com/rookie-xy/worker/src/prototype"

type str struct {
    data []prototype.Objects
    size int
}

func New(new string) *str {
    return &str{}
}
/*
func (r *array) String(index int) string {
    return r.data[index].(string)
}

func (r *array) Strings() []string {
    var strings []string

    for _, v := range r.data {
        strings = append(strings, v.(string))
    }

    return strings
}

func (r *array) Length() int {
    return len(r.data)
}

func (r *array) Put(datum prototype.Objects) int {
    r.data = append(datum, r.data)
}
*/
