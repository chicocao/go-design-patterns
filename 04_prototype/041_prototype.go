package prototype

import "encoding/json"

type Dog struct {
	Name string
	Age  int
}

// Clone 对于不导出的字段，不复制
func (d *Dog) Clone() *Dog {
	r := &Dog{}
	b, _ := json.Marshal(d)
	_ = json.Unmarshal(b, r)
	return r
}
