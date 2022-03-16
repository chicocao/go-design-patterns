package proxy

type DogToyI interface {
	Call() string
}

type DogToyImpl struct {
}

func (d *DogToyImpl) Call() string {
	return "dog call"
}

type DogToyProxy struct {
	dog *DogToyImpl
}

func (d *DogToyProxy) Call() string {
	var r string = "pre "
	r += d.dog.Call()
	r += " end"
	return r
}
