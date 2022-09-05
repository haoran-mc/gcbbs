package subject

type Observer interface {
	Update()
}

type Subject interface {
	Attach(Observer)
	Notify()
}
