package string

import (
	"fmt"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	candidates := map[string]bool{
		"123":   true,
		"a123":  false,
		"":      false,
		"01928": true,
		"0":     true,
		"0 1 3": false,
	}

	for k, v := range candidates {
		r := IsNumeric(k)
		if r != v {
			t.Errorf("(%v) ===> expected: %v, actual: %v", k, v, r)
		}
	}
}

func BenchmarkIsNumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsNumeric(ToString(i))
	}
}

func TestIsAlpha(t *testing.T) {
	r := IsAlpha("azAZ")
	if r == false {
		t.Errorf("(%v) ===> expected: true, actual: false", "abcdef")
	}
}

func BenchmarkIsAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsAlpha(fmt.Sprintf("abcdeAZ%v", rune(i)))
	}
}
