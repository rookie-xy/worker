package main

import (
    "os"
    "unsafe"

    "github.com/rookie-xy/worker/src/command"
    "github.com/rookie-xy/worker/src/module"
    "github.com/rookie-xy/worker/src/builder"
    "github.com/rookie-xy/worker/src/log"
    "fmt"
)

var (
    test = &command.Meta{ "-t", "test", false, "Test configure file" }
)

var commands = []command.Item{

    { test,
      command.LINE,
      "",
      command.SetObject,
      unsafe.Offsetof(test.Value),
      nil },

}

func init() {
    // 选项初始化
    for _, item := range commands {
        command.Pool = append(command.Pool, item)
    }

    argc, argv := len(os.Args), os.Args
    for i := 1; i < argc; i++ {
        if argv[i][0] != '-' {
            exit(-1)
        }

        fmt.Println(argv[i])
        /*
        switch argv[i] {
        }
        */
    }
}

func main() {
    core := []string{
        module.Inputs,
    }

    log := log.New()

    module := module.New(log)

    director := builder.Directors(module)
    if director == nil {
        exit(-1)
    }

    director.Construct(core)

    module.Init()

    module.Main()

    module.Exit(0)
}

func exit(code int) {
    os.Exit(code)
}
