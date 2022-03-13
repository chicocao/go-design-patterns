package factory_pattern

import "testing"

func getMainAndDetail(factory ToyFactory) {
	factory.CreateCallFactory().Call()
	factory.CreateRunFactory().Run()
}

func TestDogFactory(t *testing.T) {
	factory := &DogFactory{}
	getMainAndDetail(factory)
}

func TestCatFactory(t *testing.T) {
	factory := &CatFactory{}
	getMainAndDetail(factory)
}
