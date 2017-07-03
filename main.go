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
)

var (
    format = &command.Meta{ "-f", "format", "yaml", "Configure file format" }
    test   = &command.Meta{ "-t", "test", false, "Test configure file" }
)

var items = []command.Item{

    { format,
      command.LINE,
      module.GLOBEL,
      command.SetObject,
      unsafe.Offsetof(format.Value),
      nil },

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
        command.Items = append(command.Items, item)
    }

    argc, argv := len(os.Args), os.Args
    for i := 1; i < argc; i++ {
        if argv[i][0] != '-' {
            exit()
            //return Error
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

    worker := worker.New(log)
    build := builder.New(worker)

    director := builder.Directors(build)
    if director == nil {
        exit()
    }

    configure := configure.New(log)
    director.Construct(configure)

    worker.Init()

    worker.Main()

    worker.Exit()

}

func exit() {
    os.Exit(-1)
}
