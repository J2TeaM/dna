package ns

import (
	. "dna"
	"fmt"
	"testing"
)

func TestEncryptDescrypt(t *testing.T) {
	var x Int = 1271421
	y := Decrypt(Encrypt(x)) // 1271421 => "X1pWUkddbg" => 1271421

	if x != y {
		t.Errorf("%v (IntArray) cannot encrypt id", x)
	}
}

func ExampleEncrypt() {
	var x = Encrypt(1271421)
	fmt.Printf("%#v", x)
	// Output: "X1pWUkddbg"
}

func ExampleDecrypt() {
	var x = Decrypt("X1pWUkddbg")
	fmt.Printf("%#v", x)
	// Output: 1271421
}
