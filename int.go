package dna

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Redefine new Int Type
type Int int

// ToString returns String from Int
func (i Int) ToString() String {
	return String(fmt.Sprint(i))
}

// ToFormattedString returns a new string from Int with given width.
//
// If width is greater than the length of all digits of the number,
// it will fill with either zero or whitespace.
// If width is positive, the padding is filled from left to right.
// Otherwise, it's filled from right to left.
//
// Special case: when zero padding enable, it's only filled to the left
func (i Int) ToFormattedString(width Int, hasZeroPadding Bool) String {
	if hasZeroPadding {
		return String(fmt.Sprintf("%0[2]*[1]d", i, int(width)))
	} else {
		return String(fmt.Sprintf("%[2]*[1]d", i, int(width)))
	}
}

// Value returns primitive int type
func (i Int) ToPrimitiveType() int {
	return int(i)
}

// Value implements the Valuer interface in database/sql/driver package.
func (i Int) Value() (driver.Value, error) {
	return driver.Value(int(i)), nil
}

// Scan implements the Scanner interface in database/sql package.
// Default value for nil is 0
func (i *Int) Scan(src interface{}) error {
	var source Int
	switch src.(type) {
	case int64:
		source = Int(int(src.(int64)))
	case nil:
		source = 0
	default:
		return errors.New("Incompatible type for dna.Int type")
	}
	*i = source
	return nil
}

// ToFloat returns Float from Int
func (i Int) ToFloat() Float {
	return Float(float64(int(i)))
}

// ToHex returns hex string from Int
func (i Int) ToHex() String {
	return String(fmt.Sprintf("%x", i))
}

// ToTimeFormat returns int as seconds from 1970-01-01 to format "2013-05-14 11:00:00"
// It is used to work with postgresql
func (i Int) ToTimeFormat() String {
	tm := time.Unix(int64(i), 0)
	year := fmt.Sprint(tm.Year())
	month := fmt.Sprintf("%d", tm.Month())
	day := fmt.Sprint(tm.Day())
	hour := fmt.Sprint(tm.Hour())
	min := fmt.Sprint(tm.Minute())
	sec := fmt.Sprint(tm.Second())
	ret := year + "-" + month + "-" + day + " " + hour + ":" + min + ":" + sec
	return String(ret)
}

func (i Int) ToTime() time.Time {
	return time.Unix(int64(i), 0)
}

// ToBin returns binary string from Int
func (i Int) ToBin() String {
	return String(fmt.Sprintf("%b", i))
}

// Rand returns random number from [0,n)
func (i Int) Rand() Int {
	rand.Seed(time.Now().Unix())
	return Int(rand.Int31n(int32(i)))
}

// IsDivisibleBy returns true if the integer visible by x
func (i Int) IsDivisibleBy(x Int) Bool {
	if i%x == 0 {
		return true
	} else {
		return false
	}
}

// IsEven returns true if a number is even
func (i Int) IsEven() Bool {
	return i.IsDivisibleBy(2)
}

// IsNegative returns true if a number is negative
func (i Int) IsNegative() Bool {
	if i < 0 {
		return true
	} else {
		return false
	}
}

// IsOdd returns true if a number is odd
func (i Int) IsOdd() Bool {
	return !i.IsDivisibleBy(2)
}

// IsPositive returns true if a number is positive
func (i Int) IsPositive() Bool {
	return !i.IsNegative()
}

// CONVERTS INTEGER TO HUMAN READABLE STRING. CREDIT GOES TO http://godoc.org/github.com/dustin/go-humanize

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%dB", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+"%s", val, suffix)
}

// ToBytesFormat produces a human readable representation of an SI size.
// Bytes(67343643) -> 67MB
func (i Int) ToBytesFormat() String {
	if i.IsPositive() {
		sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
		return String(humanateBytes(uint64(i), 1000, sizes))
	} else {
		return "0B"
	}

}

// IBytes produces a human readable representation of an IEC size.
// IBytes(82854982) -> 79MiB
func (i Int) ToIBytesFormat() String {
	if i.IsPositive() {
		sizes := []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB"}
		return String(humanateBytes(uint64(i), 1024, sizes))
	} else {
		return "0B"
	}

}

// Comma produces a string form of the given number in base 10 with
// commas after every three orders of magnitude.
// e.g. Comma(834142) -> 834,142
func (i Int) ToCommaFormat() String {
	v := int64(i)
	sign := ""
	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatInt(v%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))
	return String(sign + strings.Join(parts[j:len(parts)], ","))
}

var bytesSizeTable = map[string]uint64{
	"b":   bytez,
	"kib": kiByte,
	"kb":  kByte,
	"mib": miByte,
	"mb":  mByte,
	"gib": giByte,
	"gb":  gByte,
	"tib": tiByte,
	"tb":  tByte,
	"pib": piByte,
	"pb":  pByte,
	"eib": eiByte,
	"eb":  eByte,
	// Without suffix
	"":   bytez,
	"ki": kiByte,
	"k":  kByte,
	"mi": miByte,
	"m":  mByte,
	"gi": giByte,
	"g":  gByte,
	"ti": tiByte,
	"t":  tByte,
	"pi": piByte,
	"p":  pByte,
	"ei": eiByte,
	"e":  eByte,
}

const (
	bytez  = 1
	kiByte = bytez * 1024
	miByte = kiByte * 1024
	giByte = miByte * 1024
	tiByte = giByte * 1024
	piByte = tiByte * 1024
	eiByte = piByte * 1024
)

// SI Sizes.
const (
	ibyte = 1
	kByte = ibyte * 1000
	mByte = kByte * 1000
	gByte = mByte * 1000
	tByte = gByte * 1000
	pByte = tByte * 1000
	eByte = pByte * 1000
)

// ParseBytesFormat parses a string representation of bytes into the number
// of bytes it represents.
// ParseBytesFormat("42MB") -> 42000000
// ParseBytesFormat("42mib") -> 44040192
func (s String) ParseBytesFormat() Int {
	lastDigit := 0
	for _, r := range s {
		if !(unicode.IsDigit(r) || r == '.') {
			break
		}
		lastDigit++
	}

	f, err := strconv.ParseFloat(string(s)[:lastDigit], 64)
	if err != nil {
		fmt.Printf("Error occurs at ParseBytesFormat(). %v", err)
		return 0
	}

	extra := strings.ToLower(strings.TrimSpace(string(s)[lastDigit:]))
	if m, ok := bytesSizeTable[extra]; ok {
		return Int(uint64(f * float64(m)))
	}

	fmt.Printf("Unhandled size name: %v. Error occurs at ParseBytesFormat()", extra)

	return 0
}
