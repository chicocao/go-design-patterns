package factory_pattern

import "fmt"

// Call之后Run()

// CallInterface 叫的接口
type CallInterface interface {
	Call()
}

// RunInterface 跑的接口
type RunInterface interface {
	Run()
}

// ToyFactory 抽象工厂接口
type ToyFactory interface {
	CreateCallFactory() CallInterface
	CreateRunFactory() RunInterface
}

// DogCallImpl dog
type DogCallImpl struct{}

// Call ...
func (*DogCallImpl) Call() {
	fmt.Print("dog call")
}

//DogRunImpl 为关系型数据库的OrderDetailDAO实现
type DogRunImpl struct{}

// Run ...
func (*DogRunImpl) Run() {
	fmt.Print("dog run")
}

// DogFactory Dog 抽象工厂实现
type DogFactory struct{}

func (*DogFactory) CreateCallFactory() CallInterface {
	return &DogCallImpl{}
}

func (*DogFactory) CreateRunFactory() RunInterface {
	return &DogRunImpl{}
}

// CatCallImpl cat
type CatCallImpl struct{}

// Call ...
func (*CatCallImpl) Call() {
	fmt.Print("cat call")
}

//CatRunImpl XML存储
type CatRunImpl struct{}

// Run ...
func (*CatRunImpl) Run() {
	fmt.Print("cat run")
}

// CatFactory cat 抽象工厂实现
type CatFactory struct{}

func (*CatFactory) CreateCallFactory() CallInterface {
	return &CatCallImpl{}
}

func (*CatFactory) CreateRunFactory() RunInterface {
	return &CatRunImpl{}
}
