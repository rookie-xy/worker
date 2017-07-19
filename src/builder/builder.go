package builder

import (
    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/modules/configure"
    "github.com/rookie-xy/worker/src/log"
    "github.com/rookie-xy/worker/src/observer"
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

func (r *Director) Construct(core map[string]observer.Observer) {
    scope := module.Worker

    configure := configure.New(nil)
    if configure != nil {
        for name, observer := range core {
            configure.Attach(observer)

            key := scope + "_" + name

            if module := module.Setup(key, r.Log); module != nil {
                r.build.Load(module)
            }
        }

    } else {
        return
    }

    r.build.Configure(configure)
}
