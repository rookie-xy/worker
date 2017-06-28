package observer

type Subject interface {
   	Attach(observer Observer)
   	Notify()
}

type Observer interface {
   	Update(state string)
}

type Topic struct {
   	observers []Observer
   	State     string
}

func (r *Topic) Attach(observer Observer) {
   	r.observers = append(r.observers, observer)
}

func (r *Topic) Notify() {
   	for _, observer := range r.observers {
		      observer.Update(r.State)
	   }
}

type ConcreteObserver struct {
   	state string
}

func (r *ConcreteObserver) Update(state string) {
   	r.state = state
}
