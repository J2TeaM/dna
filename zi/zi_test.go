package zi

import (
	. "dna"
	"testing"
)

func TestEncryptDescrypt(t *testing.T) {
	var x Int = 1382187828
	y := Decrypt(Encrypt(x)) // 1382187828 => "ZW6W8OOU" => 1382187828

	if x != y {
		t.Errorf("%v (IntArray) cannot encrypt id", x)
	}
}

func ExampleCheckKey() {
	var x = CheckKey("ZW6W8OOU")
	var y = CheckKey("ZW6W8OOX")
	Logv(x)
	Logv(y)
	// Output: true
	// false
}

func ExampleEncrypt() {
	var x = Encrypt(1382187828)
	Logv(x)
	// Output: "ZW6W8OOU"
}

func ExampleDecrypt() {
	var x = Decrypt("ZW6W8OOU")
	Logv(x)
	// Output: 1382187828
}

func ExampleDecodeEncodedKey() {
	var encodedKey String = "kHxmyLnaAsnHsLEtBXGybmkn"
	y := DecodeEncodedKey(encodedKey)
	Logv(y)
	// Output: "ZW67FWWF"
}
