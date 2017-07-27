package factory

import (
    "fmt"
    "github.com/rookie-xy/worker/src/plugin/codec"
)

// factory method
func Codec(cfg *codec.Config) (codec.Codec, error) {
    key := "json"
    if name := cfg.Name; name != "" {
        key = name
    }

    method := codec.Codecs[key]
    if method == nil {
        return nil, fmt.Errorf("'%v' output codec is not available", key)
    }

    return method(cfg)
}


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
