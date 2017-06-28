package worker

import (
    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/worker/src/prototype"
    "log"
)

const (
    Ok = 0
    Error = -1
)

// facade
type Worker struct {
    log.Log
    options   map[string]prototype.Object
    children  []module.Module
}

func New(log *log.Log) *Worker {
    return &Worker{}
}

func (r *Worker) Init() {
    // 初始化各个模块, inputs, channels, outputs
    for module := range r.child {
				    module.Init()
				}
}

func (r *Worker) Main() {
    // 启动，监控各个模块
    cycle := cycle.New()

    for module := range r.children {
        cycle.Start(module.Main, nil)
    }

    for ;; {
        // 监控inputs, channel, output等数据

        select {

        case cycle.Stop():

        default:

        }
    }
}

func (r *Worker) Exit() {
    // 重新加载
    select {

    case RELOAD:

    case RECONFIGURE:

    case EXIT:
        for module := range r.children {
            module.Exit()
        }
    }
}

func (r *Worker) Load(module module.Module, value map[string]prototype.Object) {
    r.options[r.children] = value
    r.children = append(r.children, module)
}
