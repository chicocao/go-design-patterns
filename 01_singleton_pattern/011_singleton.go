package singleton_pattern

type singleS struct{}

var single *singleS

func init() {
	single = &singleS{}
}

// GetSingle 饿汉式，普通单例模式，加载这个文件即创建
func GetSingle() *singleS {
	return single
}
