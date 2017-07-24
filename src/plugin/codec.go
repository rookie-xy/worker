package plugin

import "github.com/rookie-xy-bak/worker/src/prototype"

type Codec interface {
    Clone() Codec
    Init(configure prototype.Object) int
    Encode(in prototype.Object) (prototype.Object, error)
    Decode(out []byte) (prototype.Object, error)
    Type(name string) int
}
