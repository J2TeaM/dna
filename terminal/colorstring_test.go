package terminal

import (
	. "dna"
)

func ExampleColorString() {
	cs := NewColorString("Test string").Red().Inverse().Value()
	Print(cs)
	// Output is: "Test string"
	// "Test string" has black color and red background
}

func ExampleColorString_SetBackgroundColor() {
	cs := NewColorString("Test string").SetBackgroundColor(Red).Value()
	Print(cs)
	// Output is: "Test string"
	// "Test string" has red background
}

func ExampleColorString_SetTextColor() {
	cs := NewColorString("Test string").SetTextColor(Blue).Value()
	Print(cs)
	// Output is: "Test string"
	// "Test string" has blue color
}

func ExampleColorString_SetAttribute() {
	cs := NewColorString("Test string").SetTextColor(Blue).SetAttribute(ReverseCode).Value()
	Print(cs)
	// Output is: "Test string"
	// "Test string" has black color and blue background
}
