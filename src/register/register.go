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
func Module(scope, name string, items []command.Item, new module.NewFunc) {
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
    for _, e := range value {
        command.Pool = append(command.Pool, e)
    }
}

func (r *singleton) Module(key string, value module.NewFunc) {
    if module.Pool == nil {
        module.Pool = make(map[string]*module.NewFunc)
    }

    if _, ok := module.Pool[key]; !ok {
        module.Pool[key] = &value
    }
}
