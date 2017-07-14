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
func Module(scope, name string, items []command.Item, new module.New) {
    instance := getInstance()

    instance.scope = scope
    instance.name = name

    instance.Merge(items)

    if new == nil {
        return
    }

    key := ""
    if scope != key && name != key {
        key = scope+"_"+name

    } else {
        return
    }

    if module.Pool == nil {
        module.Pool = make(map[string]*module.New)
    }

    if this, ok := module.Pool[key]; !ok {
        *this = new
    }
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
