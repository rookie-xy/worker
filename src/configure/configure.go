package configure

import "github.com/rookie-xy/worker/src/module"

type Configure interface {
    GetModule() module.Module
}

type configure struct {

}

func New() *configure {
    return &configure{}
}

func (r *configure) GetModule() module.Module {
    return nil
}
