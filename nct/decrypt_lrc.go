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
//
//For example: Given a song id 2882720 with a key 9Fd4zVvPMIbf
//and an XML source file http://www.nhaccuatui.com/flash/xml?key1=8bfe9db992afaff5dc931dffca1b5c7b
//
//The encrypted lyric file url lies in <lyric> tags
// http://lrc.nct.nixcdn.com/2013/12/29/0/a/5/b/1388320374102.lrc
func DecryptLRC(data dna.String) (dna.String, error) {
	// CODE_SPECIAL "M z s 2 d k v t u 5 o d u" separated by space
	keyStr := irrcrpt(dna.Sprintf("%s", []byte{0x4d, 0x7a, 0x73, 0x32, 0x64, 0x6b, 0x76, 0x74, 0x75, 0x35, 0x6f, 0x64, 0x75}), 1)
	keyStrInHex := hex.EncodeToString(keyStr.ToBytes())

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
