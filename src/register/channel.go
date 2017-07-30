package register

import (
    "fmt"
    "github.com/rookie-xy/worker/src/channel"
)

func Channel(name string, ch channel.Cannery) {
    if _, exists := channel.Channels[name]; exists {
        panic(fmt.Sprintf("this channel '%v' already registered ", name))
    }

    channel.Channels[name] = ch
}
