package module

import (
    "github.com/rookie-xy/worker/src/log"
    "fmt"
)

const (
    Worker = "worker"
    Configure = "configure"
    Inputs = "inputs"
    Outputs = "outputs"
    Channels = "channels"
)

// composite
type Module interface {
    Load(module Template)
    Template
}

// template
type Template interface {
    Init()
    Main()
    Exit(code int)
}

type NewFunc func(log log.Log) Template
var Pool map[string]*NewFunc

type module struct {
    log.Log
    configure Module
    modules   []Template
}

func New(log log.Log) *module {
    return &module{
        Log: log,
    }
}

func (r *module) Init() {
    for _, module := range r.modules {
        if module != nil {
            module.Init()
        }
    }
}

func (r *module) Main() {
    for _, module := range r.modules {
        if module != nil {
            go module.Main()
        }
    }

    for {
        select {

        }
    }
}

func (r *module) Exit(code int) {
    for _, module := range r.modules {
        module.Exit(code)
    }
}

func (r *module) Load(module Template) {
    if module != nil {
        r.modules = append(r.modules, module)
    }
}

func (r *module) Configure(configure Module) int {
    if configure != nil {
        r.configure = configure

    } else {
        return -1
    }

    r.configure.Init()
    r.configure.Main()

    return 0
}

func Setup(key string, log log.Log) Template {
    if key == "" {
        goto J_RET
    }

    if this, ok := Pool[key]; ok {
        if new := *this; new != nil {
            return new(log)
        } else {
            fmt.Println("New func is nil")
        }

    } else {
        fmt.Println("Not found key: ", key)
    }

J_RET:
    return nil
}
