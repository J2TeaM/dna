package terminal

import (
	. "dna"
)

// Defines ColorString Type. It inherits from String
type ColorString struct {
	String
}

func NewColorString(str String) *ColorString {
	return &ColorString{str}
}

func (cs *ColorString) getColor(prefix, suffix String) *ColorString {
	return &ColorString{prefix + String(cs.ToPrimitiveValue()) + suffix}
}

// Bold returns bold text
func (cs *ColorString) Bold() *ColorString {
	return cs.getColor("\x1B[1m", "\x1B[22m")
}

// Italic returns italic text
func (cs *ColorString) Italic() *ColorString {
	return cs.getColor("\x1B[3m", "\x1B[23m")
}

// Underline returns underline text
func (cs *ColorString) Underline() *ColorString {
	return cs.getColor("\x1B[4m", "\x1B[24m")
}

// Inverse returns inverse text
func (cs *ColorString) Inverse() *ColorString {
	return cs.getColor("\x1B[7m", "\x1B[27m")
}

// Alias of Inverse
func (cs *ColorString) Reverse() *ColorString {
	return cs.getColor("\x1B[7m", "\x1B[27m")
}

// StrikeThrough returns strikeThrough text
func (cs *ColorString) StrikeThrough() *ColorString {
	return cs.getColor("\x1B[9m", "\x1B[29m")
}

// Black returns black text
func (cs *ColorString) Black() *ColorString {
	return cs.getColor("\x1B[30m", "\x1B[39m")
}

// Red returns red text
func (cs *ColorString) Red() *ColorString {
	return cs.getColor("\x1B[31m", "\x1B[39m")
}

// Green returns green text
func (cs *ColorString) Green() *ColorString {
	return cs.getColor("\x1B[32m", "\x1B[39m")
}

// Yellow returns yellow text
func (cs *ColorString) Yellow() *ColorString {
	return cs.getColor("\x1B[33m", "\x1B[39m")
}

// Blue returns blue text
func (cs *ColorString) Blue() *ColorString {
	return cs.getColor("\x1B[34m", "\x1B[39m")
}

// Magenta returns magenta text
func (cs *ColorString) Magenta() *ColorString {
	return cs.getColor("\x1B[35m", "\x1B[39m")
}

// Cyan returns cyan text
func (cs *ColorString) Cyan() *ColorString {
	return cs.getColor("\x1B[36m", "\x1B[39m")
}

// White returns white text
func (cs *ColorString) White() *ColorString {
	return cs.getColor("\x1B[37m", "\x1B[39m")
}

// Grey returns grey text
func (cs *ColorString) Grey() *ColorString {
	return cs.getColor("\x1B[90m", "\x1B[39m")
}

// BACKGROUNDS

// BlackBackground returns black-background text
func (cs *ColorString) BlackBackground() *ColorString {
	return cs.getColor("\x1B[40m", "\x1B[49m")
}

// RedBackground returns red-background text
func (cs *ColorString) RedBackground() *ColorString {
	return cs.getColor("\x1B[41m", "\x1B[49m")
}

// GreenBackground returns green-background text
func (cs *ColorString) GreenBackground() *ColorString {
	return cs.getColor("\x1B[42m", "\x1B[49m")
}

// YellowBackground returns yellow-background text
func (cs *ColorString) YellowBackground() *ColorString {
	return cs.getColor("\x1B[43m", "\x1B[49m")
}

// BlueBackground returns blue-background text
func (cs *ColorString) BlueBackground() *ColorString {
	return cs.getColor("\x1B[44m", "\x1B[49m")
}

// MagentaBackground returns bagenta-background text
func (cs *ColorString) MagentaBackground() *ColorString {
	return cs.getColor("\x1B[45m", "\x1B[49m")
}

// CyanBackground returns cyan-background text
func (cs *ColorString) CyanBackground() *ColorString {
	return cs.getColor("\x1B[46m", "\x1B[49m")
}

// WhiteBackground returns white-background text
func (cs *ColorString) WhiteBackground() *ColorString {
	return cs.getColor("\x1B[47m", "\x1B[49m")
}

// GreyBackground returns grey-background text
func (cs *ColorString) GreyBackground() *ColorString {
	return cs.getColor("\x1B[49;5;8m", "\x1B[49m")
}

// SetTextColor returns color-defined string
func (cs *ColorString) SetTextColor(color Int) *ColorString {
	var code Int
	if color >= 0 && color <= 7 {
		code = color + 30
		prefix := String("\x1B[" + string(code.ToString()) + "m")
		return cs.getColor(prefix, "\x1B[39m")
	} else {
		return cs
	}

}

// Alias of SetTextColor
func (cs *ColorString) Color(color Int) *ColorString {
	return cs.SetTextColor(color)

}

// SetBackgroundColor returns color-defined string
func (cs *ColorString) SetBackgroundColor(color Int) *ColorString {
	var code Int
	if color >= 0 && color <= 7 {
		code = color + 40
		prefix := String("\x1B[" + string(code.ToString()) + "m")
		return cs.getColor(prefix, "\x1B[49m")
	} else {
		return cs
	}

}

// Alias of SetBackgroundColor
func (cs *ColorString) Background(color Int) *ColorString {
	return cs.SetBackgroundColor(color)
}

// SetAttributeColor returns string with attribute
func (cs *ColorString) SetAttribute(attr Int) *ColorString {
	prefix := String("\x1B[" + string(attr.ToString()) + "m")
	return cs.getColor(prefix, "\x1B[0m")
}

// Alias of SetAttributes
func (cs *ColorString) Attribute(attr Int) *ColorString {
	return cs.SetAttribute(attr)
}

// Value returns the value typed String of ColorString
func (cs *ColorString) Value() String {
	return String(cs.ToPrimitiveValue())
}
