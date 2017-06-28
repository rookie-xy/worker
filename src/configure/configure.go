package configure

import (
    "unsafe"
    "github.com/rookie-xy/worker/src/command"
    "github.com/rookie-xy/worker/src/prototype"
    "fmt"
)

// abstract factory
type ConfigureFactory interface {
    GetConfigure() ConfigureMethod;
}

// factory method
type ConfigureMethod interface {
    GetFile()
}

type Configure struct {
}

func New() *Configure {
    return &Configure{}
}

func (r *Configure) GetConfigure(configType string) ConfigureMethod {
    return nil
}

func SetBool(cmd *command.Item, ptr *unsafe.Pointer, obj prototype.Object) int {
    if cmd == nil || ptr == nil {
        return -1
    }

    field := (*bool)(unsafe.Pointer(uintptr(*ptr) + cmd.Offset))

    if obj == nil {
        return -1
    }

    *field = obj.(bool)

    /*
    if command.Post != nil {
        post := command.Post.(DvrConfPostType);
        post.Handler(cf, post, *p);
    }
    */

    return 0
}

func SetString(item *command.Item, ptr *unsafe.Pointer, obj prototype.Object) int {
    if item == nil || ptr == nil || obj == nil {
        return -1
    }

    field := (*string)(unsafe.Pointer(uintptr(*ptr) + item.Offset))
    if obj == nil {
        return -1
    }

    *field = obj.(string)
    fmt.Println(*field)
    //fmt.Println(obj.(string), *field, item.Meta.Key, item.Meta.Value)
    //fmt.Println()

    return 0
}

func SetNumber(item *command.Item, ptr *unsafe.Pointer, obj prototype.Object) int {
    if item == nil || ptr == nil || obj == nil {
        return -1
    }

    field := (*int)(unsafe.Pointer(uintptr(*ptr) + item.Offset))

    if obj == nil {
        return -1
    }

    *field = obj.(int)

    return 0
}

/*
type Input interface {
    Listen()
    Reader()
}

type Output interface {
    Append()
    Writer()
}


type Channel interface {
    Input(o Observer)
    Output(o Observer)
    //Push() int
    //Pull() int
}

*/
/*
type Creater interface {
	CreateProduct(action string) Producter
	registerProduct(Producter)
}

type Producter interface {
	Use() string
}

type ConcreteCreator struct {
	products []*Producter
}

func (self *ConcreteCreator) CreateProduct(action string) Producter {
	var product Producter

	switch action {
	case "A":
		product = &ConcreteProductA{action}
	case "B":
		product = &ConcreteProductB{action}
	case "C":
		product = &ConcreteProductC{action}
	default:
		log.Fatalln("Unknown Action")
	}

	self.registerProduct(product)

	return product
}

func (self *ConcreteCreator) registerProduct(product Producter) {
	self.products = append(self.products, &product)
}

type ConcreteProductA struct {
	action string
}

func (self *ConcreteProductA) Use() string {
	return self.action
}

type ConcreteProductB struct {
	action string
}

func (self *ConcreteProductB) Use() string {
	return self.action
}

type ConcreteProductC struct {
	action string
}

func (self *ConcreteProductC) Use() string {
	return self.action
}
*/
