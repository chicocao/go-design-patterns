package factory_pattern

import "fmt"

// Animal 被封装的玩具接口
type Animal interface {
	SetMsg(msg string)
	Call()
}

// AnimalFactory 生产 Animal 玩具的工厂接口，可用来在函数中传递
type AnimalFactory interface {
	Create() Animal
}

// AnimalBase 实现 Animal 接口的基类，封装公用方法
type AnimalBase struct {
	msg string
}

// SetMsg 实现 Animal 接口的基类的共用方法
func (o *AnimalBase) SetMsg(msg string) {
	o.msg = msg
}

// DogOperatorFactory 是 DogOperator 的工厂类
type DogOperatorFactory struct{}

func (DogOperatorFactory) Create() Animal {
	return &DogOperator{
		AnimalBase: &AnimalBase{},
	}
}

// DogOperator Animal 的Dog实现
type DogOperator struct {
	*AnimalBase
}

// Call 叫
func (o *DogOperator) Call() {
	fmt.Println("Dog msg: ", o.msg)
}

// CatOperatorFactory 是 CatOperator 的工厂类
type CatOperatorFactory struct{}

func (CatOperatorFactory) Create() Animal {
	return &CatOperator{
		AnimalBase: &AnimalBase{},
	}
}

// CatOperator Animal 的 Cat 实现
type CatOperator struct {
	*AnimalBase
}

// Call 获取结果
func (c *CatOperator) Call() {
	fmt.Println("Cat msg: ", c.msg)
}
