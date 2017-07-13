package initial

import (
    "sync"
    "github.com/rookie-xy/worker/src/command"
    "github.com/rookie-xy/worker/src/module"
)

type singleton struct {
    scope string
    name  string
}

var instance *singleton
var once sync.Once

func getInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })

    return instance
}

// flyweight
func Register(scope, name string, items []command.Item, m module.Template) {
    instance := getInstance()

    instance.scope = scope
    instance.name = name

    instance.Merge(items)

    if m == nil {
        return
    }

    if module.Pool == nil {
        module.Pool = make(map[string][]module.Template)
    }

    if modules, ok := module.Pool[scope]; !ok {
        // TODO is nothing
    }

    module.Pool[scope] = append(modules, m)
}

func (r *singleton) Merge(items []command.Item) {

    if command.Pool == nil {
        command.Pool = make(map[string][]command.Item)

    } else {
        if _, ok := command.Pool[r.scope]; !ok {
            command.Pool[r.scope] = items
        }
    }
}
