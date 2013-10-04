package dna

func ExampleFloat_ToString() {
	Logv(Float(123.123).ToString())
	// Output: "123.123"
}

func ExampleFloat_ToFormattedString() {
	Logv(Float(12345.678901).ToFormattedString(6, 2))
	Logv(Float(12345.678901).ToFormattedString(4, 5))
	Logv(Float(12345.678901).ToFormattedString(10, 3))
	Logv(Float(12345.678901).ToFormattedString(15, 3))
	Logv(Float(12345.678901).ToFormattedString(-15, 3))
	// Output: "12345.68"
	// "12345.67890"
	// " 12345.679"
	// "      12345.679"
	// "12345.679      "

}

func ExampleFloat_Ceil() {
	Logv(Float(123.123).Ceil())
	Logv(Float(123.678).Ceil())
	// Output: 124
	// 124
}

func ExampleFloat_Floor() {
	Logv(Float(123.123).Floor())
	Logv(Float(123.678).Floor())
	// Output: 123
	// 123
}
