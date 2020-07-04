package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	if sum != expected {
		t.Errorf("sum %d and expected %d do not match", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(3, 3)
	fmt.Println(sum)
	// Output: 6
}
