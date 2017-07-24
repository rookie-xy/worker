package plugin

// adapter
type Event interface {
    MakeHeader()
    MakeBody()
    MakeFooter()
    Get()
    Set()
}

type event struct {
    Event
    name string
}
