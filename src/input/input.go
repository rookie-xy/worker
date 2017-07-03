package input

import "github.com/rookie-xy/worker/src/prototype"

type Input struct {
    Configure chan prototype.Object
}

func New() *Input {
    return &Input{}
}

func (r *Input) Update(configure prototype.Object) {
    if configure != nil {
        r.Configure <- configure
    }
}

func (r *Input) Wait() prototype.Object {
    //c := <- r.Configure
    //return c
    return nil
}
