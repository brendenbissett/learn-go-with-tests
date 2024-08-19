package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	t.Run("adding two numbers", func(t *testing.T) {
		sum := 4
		got := Add(2, 2)
		assertCorrectMessage(t, got, sum)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper() // This ensures that if the test fails, the calling function name is printed(not the helper function name)
	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
