package log

type Log interface {

}

type log struct {

}

func New() *log {
    return &log{}
}

type Component interface {
	Operation() string
}

type ConcreteComponent struct {
}

func (self *ConcreteComponent) Operation() string {
	return "I am component!"
}

type ConcreteDecorator struct {
	component Component
}

func (self *ConcreteDecorator) Operation() string {
	return "<strong>" + self.component.Operation() + "</strong>"
}
