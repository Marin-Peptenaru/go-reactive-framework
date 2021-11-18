package behavior

type Behavior struct {
	OnEvent    func(interface{})
	OnError    func(error)
	OnDisposed func()
}

func New() *Behavior {
	return &Behavior{}
}
