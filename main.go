package main

import (
    "os"
    "unsafe"

    "github.com/rookie-xy/worker/src/configure"
    "github.com/rookie-xy/worker/src/command"
    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/worker/src/worker"
    "github.com/rookie-xy/worker/src/builder"
    "github.com/rookie-xy/worker/src/log"
)

var (
    test     = command.Meta{ "-t", "test", false, "Test configure file" }
)

var Commands = []command.Item{

    { test,
      command.LINE,
      module.GLOBEL,
      configure.SetBool,
      unsafe.Offsetof(test.Value),
      nil },

}

func init() {
    // 选项初始化
    argc, argv := len(os.Args), os.Args
    for i := 1; i < argc; i++ {

        if argv[i][0] != '-' {
            exit()
            //return Error
        }

        for _, item := range Commands {
            handle := item.Set
            meta := item.Meta

            switch argv[i] {

            case meta.Flag:
                //
                i++
                handle(&item, &meta, argv[i])
                break

            //default:
            //    command.List(Commands)
            //    exit()
            }

            //fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAA")

            //break
        }
    }
}

func main() {
    log := log.New()

    worker := worker.New(log)
    build := builder.New(worker)

    director := builder.Director(build)
    if director == nil {
        exit()
    }

    director.Construct(resource, format)

    worker.Init()

    worker.Main()

    worker.Exit()
}

func exit() {
    os.Exit(-1)
}
