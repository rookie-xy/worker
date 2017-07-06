package builder

import (
    "github.com/rookie-xy/worker/src/worker"
    "github.com/rookie-xy/worker/src/factory"
    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/modules/inputs"
    "github.com/rookie-xy/worker/src/input"
)

type Builder interface {
    Configure(f factory.Factory) int
    Module(inputs.Input) int
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
    inputs   := inputs.New(nil)
    channels := channels.New(nil)
    outputs := outputs.New(nil)

    configure.Attach(inputs)

    //f := factory.New(configure)

    // 构建配置文件，获取配置数据
    r.build.Configure(f)

    for i,
    r.build.Module(inputs.Name)

    // TODO 解析配置数据

    // TODO 构建三大模块

    //command.
/*
    var data map[string]prototype.Object

    for _, d := range data {

    }
    */
}

type WorkerBuilder struct {
    factory.Factory
    *worker.Worker
}

func New(w *worker.Worker) *WorkerBuilder {
    return &WorkerBuilder{Worker: w}
}

func (r *WorkerBuilder) Configure(f factory.Factory) int {

    //command.Init()

    //configure := module.Init("configure", "name")

    //configure.Main()

    /*
    if f != nil {
        r.Factory = f
    } else {
        return -1
    }

    configure := f.GetModule("configure")

    configure.Init()

    configure.Main()
    */

    return 0
}

func (r *WorkerBuilder) Module(name string) int {
    if handle, ok =: module.Pool[name]; ok {
        run := handle()
        r.Load(run)
    }
    /*
    r.Load()

    module := r.GetModule(name)
    r.Load(module)

    return 0
    */
    //command.Init()
    //module.Init()

    return 0
}
/*
func (r *WorkerBuilder) Channels(f factory.Factory) int {
    return 0
}

func (r *WorkerBuilder) Outputs(f factory.Factory) int {
    return 0
}
*/
