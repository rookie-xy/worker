package builder

import (
    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/modules/configure"
    "github.com/rookie-xy/worker/src/log"
)

type Builder interface {
    Configure(m module.Module) int
    module.Module
}

type Director struct {
    log.Log
    build Builder
}

func Directors(b Builder) *Director {
    return &Director{build: b}
}

func (r *Director) Construct(core []string) {

    configure := configure.New(nil)
    if configure != nil {
        for _, name := range core {
            if module := module.Create(name, r.Log); module != nil {
                configure.Attach(module)
                r.build.Load(module)
            }
        }

    } else {
        return
    }

    r.build.Configure(configure)
}
