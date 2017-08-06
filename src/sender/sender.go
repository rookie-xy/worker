package sender


type Factory func(*Config) (Sender, error)

type Sender interface {
}

var Sender = map[string]Factory{}

type Config struct {
    Name      string
    configure string
}
