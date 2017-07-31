package channel

import "github.com/rookie-xy/worker/src/event"

type Factory func(*Configure) (Channel, error)

type Channel interface {
    Push
    Pull
}

type Push interface {
    Clone() Push
    Push(event.Event) int
}

type Pull interface {
    Pull() event.Event
}

//var Channels = map[string]Factory{}
var Channels = map[string]Channel{}

type Configure struct {
    Name      string
    configure string
}
