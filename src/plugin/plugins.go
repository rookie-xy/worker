package plugin

import (
    "github.com/rookie-xy/worker/src/plugin/codec"
    "github.com/rookie-xy/worker/src/state"
)

func Codec(name string) codec.Template {
    for _, plugin := range codec.Plugins {
        if plugin.Type(name) == state.Ok {
            return plugin.Clone()
        }
    }

    return nil
}
