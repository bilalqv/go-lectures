package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, World!",
		},
		{
			items:  []string{"bilal"},
			result: "Hello, bilal!",
		},
		{
			items:  []string{"bilal", "bhat"},
			result: "Hello, bilal, bhat!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted: %s {%v}, got: %s", st.result, st.items, s)
		}
	}
}
