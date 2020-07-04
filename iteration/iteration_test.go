package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	rep := Repeat("a", 4)
	expected := "aaaa"
	if rep != expected {
		t.Errorf("expected %q doesn't match repeated %q", expected, rep)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	rep := Repeat("a", 4)
	fmt.Println(rep)
	//Output: aaaa
}
