package module

import (
    "fmt"
	"github.com/rookie-xy/worker/src/configure"
)

type Module interface {
    Init()
    Main()
    Exit(code int)
}

type module struct {
    configure.Configure
    num  int
    modules []Module
}

func New() *module {
    return &module{
        num: -1,
        Configure: configure.New(),
    }
}

func (r *module) Init() {
    if module := r.GetModule(); module != nil {
        module.Init()
        module.Main()
    }

    if r.num <= 0 {
        fmt.Println("the module is nil")
        return
    }

    for _, module := range r.modules {
        module.Init()
    }
}

func (r *module) Main() {
    if r.num <= 0 {
        fmt.Println("the module is nil")
        return
    }

    for _, module := range r.modules {
        module.Main()
    }
}

func (r *module) Exit(code int) {
    return
}
