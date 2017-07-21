package main

import (
    "os"
    "unsafe"

    "github.com/rookie-xy/worker/src/command"
    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/worker/src/builder"
    "github.com/rookie-xy/worker/src/log"
    "github.com/rookie-xy/worker/src/observer"
    "github.com/rookie-xy/worker/src/state"

    "github.com/rookie-xy/modules/inputs"
)

var (
    version = &command.Meta{ "-v", "version", "0.0.1", "version 0.0.1"    }
    help    = &command.Meta{ "-h", "help",    "", "help infomation"     }
    test    = &command.Meta{ "-t", "test",    false, "test configure file" }
)

var commands = []command.Item{

    { version,
      command.LINE,
      module.Worker,
      command.Display,
      0,
      nil },


    { help,
      command.LINE,
      module.Worker,
      command.List,
      0,
      nil },

    { test,
      command.LINE,
      module.Worker,
      command.SetObject,
      unsafe.Offsetof(test.Value),
      nil },

}

func init() {
    // ok
    for _, item := range commands {
        command.Pool = append(command.Pool, item)
    }

    argc, argv := len(os.Args), os.Args
    if argc <= 1 {
        command.Setup(help.Flag, "")
        exit(state.Done)
    }

    for i := 1; i < argc; i++ {
        if argv[i][0] != '-' {
            exit(-1)
        }

        j := i
        if j = i + 1; j >= argc {
            j = i
        }

        flag, value := argv[i], argv[j]
        if status := command.Setup(flag, value); status != state.Ok {
            exit(status)
        }

        i++
    }
}

func main() {
    log := log.New()

    core := map[string]observer.Observer{
        module.Inputs: inputs.New(log),
    }

    module := module.New(log)

    director := builder.Directors(module)
    if director == nil {
        exit(state.Error)
    }

    director.Construct(core)

    module.Init()

    module.Main()

    module.Exit(state.Ok)
}

func exit(code int) {
    os.Exit(code)
}
