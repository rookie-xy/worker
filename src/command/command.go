package command

import (
    "github.com/rookie-xy/worker/src/prototype"
    "fmt"
)

const (
    LINE = 1
    FILE = 2
)

type SetFunc func(cmd *Item, meta *Meta, val prototype.Object) int

type Item struct {
    Meta   *Meta
    Type    int
    Scope   string
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

var Pool []Item

func List() {
    fmt.Println("list")
}

func Setup(flag, value string) int {
    for _, item := range Pool {

        if item.Type != LINE || item.Meta.Flag != flag {
            continue
        }

        if item.Set(&item, item.Meta, value) == -1 {
            fmt.Println("error")
            return -1
        }

        return 0
    }

    return -1
}

func SetObject(cmd *Item, meta *Meta, value prototype.Object) int {
    meta.Value = value
    return 0
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
