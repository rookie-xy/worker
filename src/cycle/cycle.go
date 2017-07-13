package cycle

import "github.com/rookie-xy/worker/src/prototype"

type Cycle interface {
    Start(func(), prototype.Object)
    Stop()
}

type cycle struct {

}

func New() *cycle {
    return &cycle{}
}

func (r *cycle) Start(func(), prototype.Object) {

}

func (r *cycle) Stop() {

}
