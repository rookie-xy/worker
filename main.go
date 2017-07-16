package main

import (
    "os"
    "unsafe"

    "github.com/rookie-xy/worker/src/command"
    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/worker/src/builder"
    "github.com/rookie-xy/worker/src/log"
    "github.com/rookie-xy/modules/inputs"
)

var (
    test = &command.Meta{ "-t", "test", false, "Test configure file" }
)

var commands = []command.Item{

    { test,
      command.LINE,
      module.GLOBEL,
      command.SetObject,
      unsafe.Offsetof(test.Value),
      nil },

}

func init() {
    // 选项初始化
    for _, item := range commands {
        command.Pool[item.Meta.Flag] = item
    }

    argc, argv := len(os.Args), os.Args
    for i := 1; i < argc; i++ {
        if argv[i][0] != '-' {
            exit(-1)
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
    core := []string{
        inputs.Name,
    }

    log := log.New()

    module := module.New(log)

    director := builder.Directors(module)
    if director == nil {
        exit()
    }

    director.Construct(core)

    module.Init()

    module.Main()

    module.Exit(0)
}

func exit(code int) {
    os.Exit(code)
}
