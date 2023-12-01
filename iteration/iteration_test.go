package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	got := Iterate("a", 10)
	expected := "aaaaaaaaaa"
	if got != expected {
		t.Errorf("expected=%s, got=%s", expected, got)
	}

	got = Iterate("a", 1)
	expected = "a"
	if got != expected {
		t.Errorf("expected=%s, got=%s", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterate("a", 10)
	}
}

func ExampleIterate() {
	text := Iterate("t", 5)
	fmt.Println(text)
	// Output: ttttt
}
