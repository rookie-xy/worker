package register

import (
    "hubble/src/module"
    "hubble/src/command"
    "sync"
)

type singleton struct {
    pool map[string]Register
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
   	once.Do(func() {
		      instance = &singleton{}
	   })

   	return instance
}

// flyweight
type Register interface {
    GetModule() module.ModuleTemplate
    GetCommands() []command.Command
}

/*
type RegisterFactory struct {
   	pool map[string]Register
}
*/

func (r *singleton) Module(name string, mod module.ModuleTemplate, cmd []command.Command) Register {
   	if r.pool == nil {
	      	r.pool = make(map[string]Register)
	   }

	   if _, ok := r.pool[name]; !ok {
		      r.pool[name] = &item{module: mod, commands:cmd}
	   }

	   return r.pool[name]
}

type item struct {
   module   module.Module
   commands []command.Command
}

func (r *item) GetModule() module.Module {
    return r.module
}

func (r *item) GetCommands() []command.Command {
   return r.commands
}
