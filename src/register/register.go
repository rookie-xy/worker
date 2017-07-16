package register

import (
    "sync"
    "github.com/rookie-xy/worker/src/command"
    "github.com/rookie-xy/worker/src/module"
)

// singleton
type singleton struct {
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
    merge := getInstance()

    key := ""
    if scope != key && name != key {
        key = scope+"_"+name

    } else {
        return
    }

    if l := len(items); l <= 0 {
        return
    } else {
        merge.Command(key, items)
    }

    if new != nil {
        merge.Module(key, new)
    }
}

func (r *singleton) Command(key string, value []command.Item) {
    if command.Pool == nil {
        command.Pool = make(map[string][]command.Item)

    } else {
        if _, ok := command.Pool[key]; !ok {
            command.Pool[key] = value
        }
    }
}

func (r *singleton) Module(key string, value module.New) {
    if module.Pool == nil {
        module.Pool = make(map[string]*module.New)

    } else {
        if this, ok := module.Pool[key]; !ok {
            *this = value
        }
    }
}
