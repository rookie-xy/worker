package register

import (
    "fmt"
    "github.com/rookie-xy/worker/src/plugin/codec"
)

func Codec(name string, f codec.Factory) {
    if _, exists := codec.Codecs[name]; exists {
        panic(fmt.Sprintf("output codec '%v' already registered ", name))
    }

    codec.Codecs[name] = f
}
