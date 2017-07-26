package codec

import (
    "fmt"
    "github.com/rookie-xy/worker/src/prototype"
)

type Factory func(*Config) (Codec, error)

type Codec interface {
    Encode(in prototype.Object) (prototype.Object, error)
    Decode(out []byte) (prototype.Object, error)
}

var Codecs = map[string]Factory{}

type Config struct {
    Name      string
    configure string
}

func CreateCodec(cfg *Config) (Codec, error) {
    // default to json codec
    codec := "json"
    if name := cfg.Name; name != "" {
        codec = name
    }

    factory := Codecs[codec]
    if factory == nil {
        return nil, fmt.Errorf("'%v' output codec is not available", codec)
    }

    return factory(cfg)
}
