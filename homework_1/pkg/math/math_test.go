package math

import "testing"

func TestFib(t *testing.T) {
	got := Fib(10)
	want := 55

	if got != want {
		t.Fatalf("получили %d, ожидали %d", got, want)
	}
}
