package codec

import (
    "github.com/rookie-xy/worker/src/prototype"
)

type Codec interface {
    Clone() Codec
    Init(configure prototype.Object) int
    Template
    Type(name string) int
}

type Template interface {
    Encode(in prototype.Object) (prototype.Object, error)
    Decode(out []byte) (prototype.Object, error)
}

var Plugins []Codec

