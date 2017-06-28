package builder

import (
    "github.com/rookie-xy/worker/src/worker"
    "github.com/rookie-xy/worker/src/prototype"
    "github.com/rookie-xy/worker/src/configure"
    "github.com/rookie-xy/worker/src/command"

    "github.com/rookie-xy/modules/inputs"

    "github.com/rookie-xy/plugins"
)

type Builder interface {
    Configure(resource string) map[string]prototype.Object
   	Inputs(object prototype.Object)
	   Channels(object prototype.Object)
	   Outputs(object prototype.Object)
}

type Director struct {
   	build Builder
}

func Director(b Builder) *Director {
    return &Director{build: b}
}

func (r *Director) Construct(resource command.Meta) {
    // 构建配置，三大模块
/*
    name := configure.Name
    if v := resource.Value; v != nil {
				    name = v.GetString()
				}
				*/

    // 返回解析完成的kv配置数据
    configure := r.build.Configure("")
    if configure == nil {
				    return
				}

    for key, value := range configure {

				    switch key {

				    case inputs.Name:
				        r.build.Inputs(value)

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
    // 通过插件获取配置方法
/*
    cfg := configure.New(nil)

    if p := plugins.Configure(resource); p != nil {
				    cfg = configure.New(p)
				}

    method := cfg.Method
    */

    return nil
}

func (r *WorkerBuilder) Inputs(object prototype.Object) {
    input := inputs.New(r.Log)
    input.Options = object

    r.Load(input, nil)
}

func (r *WorkerBuilder) Channels(object prototype.Object) {
}

func (r *WorkerBuilder) Outputs(object prototype.Object) {
}
