package greet

import "testing"

func TestGreet(t *testing.T) {
    expected := "Hello, Muhamad Fachri Haikal"
    input := "Muhamad Fachri Haikal"

    if Greet(input) != expected{
        t.Errorf("Unexpected output. Expected: %q, but got %q", expected, Greet(input))
    }
}