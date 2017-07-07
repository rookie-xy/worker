package main

import (
    "os"
    "unsafe"

    "github.com/rookie-xy/worker/src/command"
    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/worker/src/worker"
    "github.com/rookie-xy/worker/src/builder"
    "github.com/rookie-xy/worker/src/log"
    "github.com/rookie-xy/worker/src/configure"
    "github.com/rookie-xy/worker/src/cycle"
    "github.com/rookie-xy/worker/src/run"
)

var (
    //config = &command.Meta{ "-c", "config", "file", "This configure file path" }
    test   = &command.Meta{ "-t", "test", false, "Test configure file" }
)

var items = []command.Item{

    { test,
      command.LINE,
      module.GLOBEL,
      command.SetObject,
      unsafe.Offsetof(test.Value),
      nil },

}

func init() {
    // 选项初始化
    for _, item := range items {
        command.Pool[item.Meta.Flag] = item
    }

    argc, argv := len(os.Args), os.Args
    for i := 1; i < argc; i++ {
        if argv[i][0] != '-' {
            exit()
            //return Error
        }

        command.Pool[argv[i]]

        switch argv[i] {
        case command.Pool[]
        }

        for _, item := range command.Items {
            handle := item.Set
            meta := item.Meta
            switch argv[i] {

            case meta.Flag:
                i++
                handle(&item, meta, argv[i])
                break

            //default:
            //    command.List(Commands)
            //    exit()
            }
        }
    }
}

func main() {
    log := log.New()

    run := run.New(log)
    build := builder.New(run)

    director := builder.Directors(build)
    if director == nil {
        exit()
    }

    director.Construct()

    run.Main()

    run.Exit(0)
}

func exit() {
    os.Exit(-1)
}
