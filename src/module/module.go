package module

import (
    "github.com/rookie-xy/worker/src/log"
    "fmt"
)

const (
    GLOBEL = 2
    LOCAL = 3
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

type New func(log log.Log) Template
var Pool map[string]*New

type module struct {
    log.Log
    configure Module
    modules   []Module
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

func Create(name string, log log.Log) Template {
    if name != "" {
        if this, ok := Pool[name]; ok {
            if new := *this; new != nil {
                return new(log)
            } else {
                fmt.Println("New func is nil")
            }

        } else {
            fmt.Println("Not found key")
        }
    }

    return nil
}
