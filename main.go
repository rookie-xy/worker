// 2017 Meng Shi

package main

import (
    "github.com/rookie-xy/worker/src/module"
)

var modules = module.New()

func init() {
    modules.Init()
}

func main() {

    modules.Main()

    exit()
}

func exit() {
    modules.Exit(-1)
}
