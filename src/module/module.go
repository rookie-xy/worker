package module

const (
    GLOBEL = 2
    LOCAL = 3
)

// composite
type Module interface {
    Load(module Template)
    Template
}

// template
type Template interface {
    Init(name string) Template
    Main()
    Exit()
}

var Pool map[string][]Template

func Init(scope, name string) Template {
    if scope == "" || name == "" {
        return nil
    }

    if modules, ok := Pool[scope]; ok {
        for _, module := range modules {

            if run := module.Init(name); run != nil {
                return run
            }
        }
    }

    return nil
}
/*
type Component interface {
	Add(child Component)
	Name() string
	Childs() []Component
	Print(prefix string) string
}

type Directory struct {
	name   string
	childs []Component
}

func (self *Directory) Add(child Component) {
	self.childs = append(self.childs, child)
}

func (self *Directory) Name() string {
	return self.name
}

func (self *Directory) Childs() []Component {
	return self.childs
}

func (self *Directory) Print(prefix string) string {
	result := prefix + "/" + self.Name() + "\n"
	for _, val := range self.Childs() {
		result += val.Print(prefix + "/" + self.Name())
	}
	return result
}

type File struct {
	name   string
}

func (self *File) Add(child Component) {
}

func (self *File) Name() string {
	return self.name
}

func (self *File) Childs() []Component {
	return []Component{}
}

func (self *File) Print(prefix string) string {
	return prefix + "/" + self.Name() + "\n"
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}
*/
