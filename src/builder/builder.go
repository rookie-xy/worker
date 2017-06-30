package builder

import (
    "github.com/rookie-xy/worker/src/worker"
    "github.com/rookie-xy/worker/src/factory"
)

type Builder interface {
    Configure(f factory.Factory) int
/*
   	Inputs(f factory.Factory) int
	   Channels(f factory.Factory) int
	   Outputs(f factory.Factory) int
	   */
}

type Director struct {
   	build Builder
}

func Directors(b Builder) *Director {
    return &Director{build: b}
}

func (r *Director) Construct() {
    // 构建配置，三大模块
    // 返回解析完成的kv配置数据
    f := factory.New(nil)

    configure := r.build.Configure(f)
    if configure != 0 {
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
				    return 0
				}

    configure := f.GetModule("configure")

    configure.Init()

    configure.Main()

    return 0
}
