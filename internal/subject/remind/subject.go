package remind

import (
	"github.com/haoran-mc/gcbbs/internal/subject"
)

type remindSubject struct {
	observers []subject.Observer
}

// New ...
func New() *remindSubject {
	return &remindSubject{
		observers: make([]subject.Observer, 0),
	}
}

// Attach ...
func (o *remindSubject) Attach(observer subject.Observer) {
	o.observers = append(o.observers, observer)
}

// Notify ...
func (o *remindSubject) Notify() {
	for _, s := range o.observers {
		s.Update()
	}
}
