package run

import (
    "github.com/rookie-xy/run/src/module"
    "github.com/rookie-xy/run/src/log"
    "github.com/rookie-xy/run/src/cycle"
)

const (
    Ok = 0
    Error = -1
)

const (
    RELOAD = 1
    RECONFIGURE = 2
    EXIT = 3
)

// facade
type Run struct {
    log.Log
    children []module.Module
}

func New(log log.Log) *Run {
    return &Run{}
}

func (r *Run) Main() {
    // 启动，监控各个模块
    for _, run := range Pool {
        go run.Main()
    }

    for ;; {
        // 监控inputs, channel, output等数据
        select {

        case cycle.Stop():

        default:

        }
    }
}

func (r *Run) Exit(code int) {
    // 重新加载
    /*
    select {

    case <- RELOAD:

    case <- RECONFIGURE:

    case <- EXIT:
        for _, module := range r.children {
            module.Exit()
        }
    }
    */
}

func (r *Run) Load(module module.Module) {
    r.children = append(r.children, module)
}
