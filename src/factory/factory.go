package factory

import (
    "github.com/rookie-xy/worker/src/configure"
    "github.com/rookie-xy/worker/src/module"

   configure_module "github.com/rookie-xy/modules/configure"
    "github.com/rookie-xy/modules/inputs"
)

// abstract factory
type Factory interface {
    GetModule(name string) module.Module
}

type factory struct {
    *configure.Configure
}

func New(c *configure.Configure) *factory {
    return &factory{
        Configure: c,
    }
}

func (r *factory) GetModule(name string) module.Module {
    switch name {

    case configure_module.Name:
        return configure_module.New(r.Configure)

    case inputs.Name:
        return inputs.New(r.Log)

    }

    return nil
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
