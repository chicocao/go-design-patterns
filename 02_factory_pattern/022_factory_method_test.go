package factory_pattern

import "testing"

func Caller(a AnimalFactory,msg string) {
	toy := a.Create()
	toy.SetMsg(msg)
	toy.Call()
}

func TestFactoryMethod(t *testing.T) {
	dog := &DogOperatorFactory{}
	Caller(dog,"toy dog")
	cat :=&CatOperatorFactory{}
	Caller(cat,"toy cat")
}
