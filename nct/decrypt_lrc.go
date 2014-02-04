package nct

import (
	"crypto/rc4"
	"dna"
	"encoding/hex"
)

func irrcrpt(_arg1 dna.String, _arg2 dna.Int) dna.String {
	var _local5 dna.Int
	var _local3 dna.String = ""
	var _local4 dna.Int
	for _local4 < _arg1.Length() {
		_local5 = _arg1.CharCodeAt(_local4)
		if (_local5 >= 48) && (_local5 <= 57) {
			_local5 = ((_local5 - _arg2) - 48)
			if _local5 < 0 {
				_local5 = (_local5 + ((57 - 48) + 1))
			}
			_local5 = ((_local5 % ((57 - 48) + 1)) + 48)
		} else {
			if (_local5 >= 65) && (_local5 <= 90) {
				_local5 = ((_local5 - _arg2) - 65)
				if _local5 < 0 {
					_local5 = (_local5 + ((90 - 65) + 1))
				}
				_local5 = ((_local5 % ((90 - 65) + 1)) + 65)
			} else {
				if (_local5 >= 97) && (_local5 <= 122) {
					_local5 = ((_local5 - _arg2) - 97)
					if _local5 < 0 {
						_local5 = (_local5 + ((122 - 97) + 1))
					}
					_local5 = ((_local5 % ((122 - 97) + 1)) + 97)
				}
			}
		}
		_local3 = (_local3 + dna.FromCharCode(_local5))
		_local4++
	}
	return (_local3)
}

//DecryptLRC returns LRC string from encrypted string.
func DecryptLRC(data dna.String) (dna.String, error) {
	keyStr := irrcrpt("Mzs2dkvtu5odu", 1).String()
	keyStrInHex := hex.EncodeToString([]byte(keyStr))

	keyStrInBytes, err := hex.DecodeString(keyStrInHex)
	if err != nil {
		return "", err
	}

	ret, err := hex.DecodeString(data.String())
	if err != nil {
		return "", err
	}

	cipher, err := rc4.NewCipher(keyStrInBytes)
	if err != nil {
		return "", err
	} else {
		cipher.XORKeyStream(ret, ret)
		return dna.String(string(ret)), nil
	}
}
