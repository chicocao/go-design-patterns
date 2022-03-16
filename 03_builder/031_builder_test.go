package builder

import "testing"

func TestAnimalBuilder(t *testing.T) {
	d := &Dog{}
	c := &Cat{}

	cases := []struct {
		a *AnimalToyBuilder
		r *Res
	}{
		{
			a: &AnimalToyBuilder{
				animal: d,
			},
			r: &Res{
				name: "dog call",
			},
		},
		{
			a: &AnimalToyBuilder{
				animal: c,
			},
			r: &Res{
				name: "cat call",
			},
		},
	}
	for i, c := range cases {
		res := c.a.BuilderCall()
		if res == nil || res.name != c.r.name {
			t.Errorf("Dog builder index: %d err: %v", i, res)
		}
	}

}
