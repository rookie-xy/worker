package instance

import (
    "sync"
    "github.com/rookie-xy/worker/src/command"
    "github.com/rookie-xy/worker/src/module"
)

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
func Register(name string, items []command.Item, mod module.Module) {
    instance := getInstance()

    instance.Merge(items)

    if module.Pool == nil {
				    module.Pool = make(map[string]module.Pool)
				}

    if _, ok := module.Pool[name]; !ok {
				    module.Pool[name] =  mod
				}
}

func (r *singleton) Merge(items []command.Item) {
    for _, item := range items {
				    command.Items = append(command.Items, item)
				}
}
