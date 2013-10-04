package dna

// NewString initializes and returns new String
func NewString(str string) String {
	return String(str)
}

// NewText initializes and returns new Text
func NewText(str string) *Text {
	ret := Text(str)
	return &ret
}

// NewInt initializes and returns new Int
func NewInt(in int) Int {
	return Int(in)
}

// NewFloat initializes and returns new Float
func NewFloat(f float64) Float {
	return Float(f)
}

// NewIntArray initializes and returns new IntArray
func NewIntArray(sa []int) IntArray {
	a := make(IntArray, len(sa))
	for i, v := range sa {
		a[i] = Int(v)
	}
	return a
}

//  NewStringArray initializes and returns new StringArray
func NewStringArray(sa []string) StringArray {
	a := make(StringArray, len(sa))
	for i, v := range sa {
		a[i] = String(v)
	}
	return a
}
