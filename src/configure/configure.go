package configure

import (
    "github.com/rookie-xy/worker/src/observer"
    "github.com/rookie-xy/worker/src/prototype"
    "github.com/rookie-xy/worker/src/log"
)

type Configure struct {
    log.Log
   	observers []observer.Observer
   	Data prototype.Object
}

func New(log log.Log) *Configure {
    return &Configure{
				    Log: log,
				}
}

func (r *Configure) Attach(obs observer.Observer) {
   	r.observers = append(r.observers, obs)
}

func (r *Configure) Notify() {
   	for _, observer := range r.observers {
		      observer.Update(r.Data)
	   }
}
