package builder

import (
    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/modules/configure"
    "github.com/rookie-xy/worker/src/log"
    "github.com/rookie-xy/worker/src/observer"
    "github.com/rookie-xy-bak/worker/src/prototype"
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

func (r *Director) Construct(core map[string]prototype.Object) {

    configure := configure.New(nil)
    if configure != nil {
        for _, module := range core {
            configure.Attach(module)
            r.build.Load(module)
            /*
            if module := module.News(name, r.Log); module != nil {
                configure.Attach(value)
                r.build.Load(module)
            }
            */
        }

    } else {
        return
    }

    r.build.Configure(configure)
}
