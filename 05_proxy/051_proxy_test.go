package proxy

import (
	"testing"
)

func TestCall(t *testing.T) {
	p := DogToyProxy{&DogToyImpl{}}
	c := p.Call()
	want := "pre dog call end"
	if c != want {
		t.Errorf("want: %v but get: %v", want, c)
	}
}
