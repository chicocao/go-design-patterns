package builder

type Res struct {
	name string
}

type Animal interface {
	Call() *Res
}

type AnimalToyBuilder struct {
	animal Animal
}

func (a *AnimalToyBuilder) BuilderCall() *Res {
	return a.animal.Call()
}

type Dog struct {
}

func (d *Dog) Call() *Res {
	return &Res{
		name: "dog call",
	}
}

type Cat struct {
}

func (c *Cat) Call() *Res {
	return &Res{
		name: "cat call",
	}
}
