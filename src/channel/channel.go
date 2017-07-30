package channel

import "github.com/rookie-xy/worker/src/event"

type Factory func(*Configure) (Channel, error)

type Channel interface {
    Push
    Pull
}

type Push interface {
    Push(event.Event) int
}

type Pull interface {
    Pull() event.Event
}

var Channels = map[string]Factory{}

type Configure struct {
    Name      string
    configure string
}
