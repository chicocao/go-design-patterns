package factory_pattern

import "fmt"

type HelloEnum int

const (
	HelloEnumDog = iota
	HelloEnumCat
)

type HelloInterface interface {
	SayHello(msg string)
}

type DogHello struct {
}

func (d *DogHello) SayHello(msg string) {
	fmt.Println("dog say: ", msg)
}

type CatHello struct {
}

func (c *CatHello) SayHello(msg string) {
	fmt.Println("cat say: ", msg)
}

// NewHello 根据opt创建，没有枚举到，则返回DogHello
func NewHello(opt HelloEnum) HelloInterface {
	switch opt {
	case HelloEnumDog:
		return &DogHello{}
	case HelloEnumCat:
		return &CatHello{}
	}

	return &DogHello{}
}
