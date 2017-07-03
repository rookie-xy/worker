package builder

import (
    "github.com/rookie-xy/worker/src/worker"
    "github.com/rookie-xy/worker/src/factory"
    "github.com/rookie-xy/worker/src/configure"
    "github.com/rookie-xy/worker/src/input"
    "github.com/rookie-xy/worker/src/module"
)

type Builder interface {
    Configure(f factory.Factory) int
    Module(n string) int
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

func (r *Director) Construct(configure *configure.Configure) {
    input := input.New()

    configure.Attach(input)

    f := factory.New(configure)

    // 构建配置文件，获取配置数据
    r.build.Configure(f)

    // TODO 解析配置数据
    input.Wait()

    // TODO 构建三大模块
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

    configure := module.Init("configure", "name")

    configure.Main()

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

    module := r.GetModule(name)
    r.Load(module)

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
