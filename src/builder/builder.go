package builder

import (
    "github.com/rookie-xy/worker/src/worker"
    "github.com/rookie-xy/worker/src/prototype"
    "github.com/rookie-xy/worker/src/factory"
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

func (r *Director) Construct() {
    // 构建配置，三大模块
    // 返回解析完成的kv配置数据
    f := factory.New(nil)

    configure := r.build.Configure(f)
    if configure == nil {
				    return
				}
}

type WorkerBuilder struct {
   	*worker.Worker
}

func New(w *worker.Worker) *WorkerBuilder {
    return &WorkerBuilder{Worker: w}
}

func (r *WorkerBuilder) Configure(f factory.Factory) int {
    if f == nil {
				    return nil
				}

    configure := f.GetModule("configure")

    configure.Init()

    go configure.Main()

    return nil
}
