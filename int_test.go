package dna

import (
	"testing"
)

func TestToString(t *testing.T) {
	var x Int = 123
	if x.ToString() != "123" {
		t.Errorf("%v cannot be converted to string", x)
	}
}

func TestToHex(t *testing.T) {
	var x Int = 123
	if x.ToHex() != "7b" {
		t.Errorf("%v cannot be converted to hex string", x)
	}
}

func TestInt_Rand(t *testing.T) {
	var value Int = 1000
	for i := 0; i < 10000; i++ {
		tmp := Int(value).Rand()
		if tmp < 0 && tmp >= value {
			t.Errorf("%v can not generate random number in range", tmp)
		}
	}
	// Output 0 <= number < 10
}

func ExampleInt_ToString() {
	var x Int = 66666
	Logv(x.ToString())
	// Output: "66666"
}

func ExampleInt_ToFormattedString() {
	var x Int = 66666
	Logv(x.ToFormattedString(10, false))
	Logv(x.ToFormattedString(-10, false))
	Logv(x.ToFormattedString(10, true))
	Logv(x.ToFormattedString(-10, true))
	Logv(x.ToFormattedString(-1, true))
	Logv(x.ToFormattedString(1, false))
	// Output:
	// "     66666"
	// "66666     "
	// "0000066666"
	// "66666     "
	// "66666"
	// "66666"
}

func ExampleInt_ToHex() {
	var x Int = 66666
	Logv(x.ToHex())
	// Output: "1046a"
}

func ExampleInt_ToBin() {
	var x Int = 65
	Logv(x.ToBin())
	// Output: "1000001"
}

func ExampleInt_Rand() {
	Logv(Int(10).Rand())
	// Output 0 <= number < 10
}

func ExampleInt_IsDivisibleBy() {
	var x Int = 66666
	Logv(x.IsDivisibleBy(6))
	Logv(x.IsDivisibleBy(9999))
	// Output: true
	// false
}

func ExampleInt_IsEven() {
	Logv(Int(66666).IsEven())
	Logv(Int(123).IsEven())
	// Output: true
	// false
}

func ExampleInt_IsNegative() {
	Logv(Int(-66666).IsNegative())
	Logv(Int(123).IsNegative())
	// Output: true
	// false
}

func ExampleInt_IsOdd() {
	Logv(Int(123).IsOdd())
	Logv(Int(6666).IsOdd())
	// Output: true
	// false
}

func ExampleInt_IsPositive() {
	Logv(Int(123).IsPositive())
	Logv(Int(-6666).IsPositive())
	// Output: true
	// false
}

func ExampleInt_ToBytesFormat() {
	Logv(Int(87123223).ToBytesFormat())
	Logv(Int(1234567890).ToBytesFormat())
	Logv(Int(-1234567890).ToBytesFormat())
	Logv(Int(0).ToBytesFormat())
	// Output: "87MB"
	// "1.2GB"
	// "0B"
	// "0B"
}

func ExampleInt_ToIBytesFormat() {
	Logv(Int(87123223).ToIBytesFormat())
	Logv(Int(1234567890).ToIBytesFormat())
	Logv(Int(-1234567890).ToIBytesFormat())
	// Output: "83MiB"
	// "1.1GiB"
	// "0B"
}

func ExampleInt_ToCommaFormat() {
	Logv(Int(87123223).ToCommaFormat())
	Logv(Int(1234567890).ToCommaFormat())
	// Output: "87,123,223"
	// "1,234,567,890"
}

func ExampleString_ParseBytesFormat() {
	Logv(String("42MB").ParseBytesFormat())
	Logv(String("42mib").ParseBytesFormat())
	// Output: 42000000
	// 44040192
}
