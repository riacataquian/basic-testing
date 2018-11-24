package greeter

import "testing"

func TestGreet(t *testing.T) {
	in := "john"
	got := Greet(in)
	want := "Hello, john!"

	if got != want {
		t.Errorf("Greet(%s) = %q, want %q", in, got, want)
	}
}
