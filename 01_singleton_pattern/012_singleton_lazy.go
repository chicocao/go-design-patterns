package singleton_pattern

import "sync"

var (
	singleLazy *singleS
	one        = &sync.Once{}
)

// GetSingleLazy 懒汉模式：用的时候再初始化
func GetSingleLazy() *singleS {
	if singleLazy != nil {
		return singleLazy
	}
	one.Do(func() {
		singleLazy = &singleS{}
	})
	return singleLazy
}
