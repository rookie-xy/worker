package command

import (
    "github.com/rookie-xy/worker/src/prototype"
    "fmt"
)

const (
    LINE = 1
    FILE = 2
)

type SetFunc func(cmd *Item, meta *Meta, obj prototype.Object) int

type Item struct {
    Meta   *Meta
    Type    int
    Scope   int
    Set     SetFunc
    Offset  uintptr
    Load    prototype.Object
}

type Meta struct {
    Flag     string
    Key      string
    Value    prototype.Object
				Details  string
}

var Items []Item

func (r *Meta)GetValueString() string {
    return ""
}

func List(commands []Item) {
    fmt.Println("list")

}

/*
type Command interface {
   	Execute() string
}

type Option struct {

}

type Command interface {
	Execute() string
}

type ToggleOnCommand struct {
	receiver *Receiver
}

func (self *ToggleOnCommand) Execute() string {
	return self.receiver.ToggleOn()
}

type ToggleOffCommand struct {
	receiver *Receiver
}

func (self *ToggleOffCommand) Execute() string {
	return self.receiver.ToggleOff()
}

type Receiver struct {
}

func (self *Receiver) ToggleOn() string {
	return "Toggle On"
}

func (self *Receiver) ToggleOff() string {
	return "Toggle Off"
}

type Invoker struct {
	commands []Command
}

func (self *Invoker) StoreCommand(command Command) {
	self.commands = append(self.commands, command)
}

func (self *Invoker) UnStoreCommand() {
	if len(self.commands) != 0 {
		self.commands = self.commands[:len(self.commands)-1]
	}
}

func (self *Invoker) Execute() string {
	var result string
	for _, command := range self.commands {
		result += command.Execute() + "\n"
	}
	return result
}
*/
