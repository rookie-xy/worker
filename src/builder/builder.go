package builder

import (
    "github.com/rookie-xy/worker/src/worker"
    "github.com/rookie-xy/worker/src/prototype"
    "github.com/rookie-xy/worker/src/configure"
    "github.com/rookie-xy/worker/src/command"

    "github.com/rookie-xy/modules/inputs"
    "github.com/rookie-xy/modules/channels"
    "github.com/rookie-xy/modules/outputs"
)

type Builder interface {
    Configure() map[string]prototype.Object
   	Inputs()
	   Channels()
	   Outputs()
}

type Director struct {
   	build Builder
}

func Director(b Builder) *Director {
    return &Director{build: b}
}

func (r *Director) Construct(resource command.Meta) {
    // 构建配置，三大模块
    name := configure.Name
    if v := resource.Value; v != nil {
				    name = v.GetString()
				}

    configure := r.build.Configure(name)
    if configure == nil {
				    return
				}

    for key, value := range configure {

				    switch key {

				    case inputs.Name:
				        r.build.Inputs(value)

				    case channels.Name:
				        r.build.Channels(value)

								case outputs.Name:
				        r.build.Outputs(value)

				    default:
				    }
				}
}

type WorkerBuilder struct {
   	*worker.Worker
}

func New(w *worker.Worker) *WorkerBuilder {
   return &WorkerBuilder{Worker: w}
}

func (r *WorkerBuilder) Configure(resource string) map[string]prototype.Object {
    //r.Load(configure.New(), nil)
    c := configure.New()

    c.GetConfigure

    select {

				case

				}

    return nil
}

func (r *WorkerBuilder) Inputs(object prototype.Object) {
    r.Load(inputs.New(), object)
}

func (r *WorkerBuilder) Channels(object prototype.Object) {
}

func (r *WorkerBuilder) Outputs(object prototype.Object) {
}
