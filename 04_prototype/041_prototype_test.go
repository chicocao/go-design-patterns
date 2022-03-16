package prototype

import (
	"testing"
)

func TestClone(t *testing.T) {
	d := &Dog{
		Name: "dog1",
		Age:  10,
	}
	c := d.Clone()

	if d.Name != c.Name || d.Age != c.Age  {
		t.Errorf("d: %+v c: %+v", d, c)
	}
}
