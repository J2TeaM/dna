package terminal

import (
	. "dna"
)

// Defines ColorString Type. It inherits from String
type ColorString struct {
	str        String
	attributes StringArray
	bgColor    StringArray
	color      StringArray
}

func NewColorString(str String) *ColorString {
	return &ColorString{
		str:        str,
		attributes: StringArray{"", ""},
		bgColor:    StringArray{"", ""},
		color:      StringArray{"", ""},
	}
}

func (cs *ColorString) getColor(prefix, suffix, kind String) *ColorString {
	switch kind {
	case "att":
		cs.attributes = StringArray{prefix, suffix}
	case "bg":
		cs.bgColor = StringArray{prefix, suffix}
	case "color":
		cs.color = StringArray{prefix, suffix}
	default:
		panic("Wrong kind")
	}
	return cs
}

// Bold returns bold text
func (cs *ColorString) Bold() *ColorString {
	return cs.getColor("\x1B[1m", "\x1B[22m", "att")
}

// Italic returns italic text
func (cs *ColorString) Italic() *ColorString {
	return cs.getColor("\x1B[3m", "\x1B[23m", "att")
}

// Underline returns underline text
func (cs *ColorString) Underline() *ColorString {
	return cs.getColor("\x1B[4m", "\x1B[24m", "att")
}

// Inverse returns inverse text
func (cs *ColorString) Inverse() *ColorString {
	return cs.getColor("\x1B[7m", "\x1B[27m", "att")
}

// Alias of Inverse
func (cs *ColorString) Reverse() *ColorString {
	return cs.getColor("\x1B[7m", "\x1B[27m", "att")
}

// StrikeThrough returns strikeThrough text
func (cs *ColorString) StrikeThrough() *ColorString {
	return cs.getColor("\x1B[9m", "\x1B[29m", "att")
}

// Black returns black text
func (cs *ColorString) Black() *ColorString {
	return cs.getColor("\x1B[30m", "\x1B[39m", "color")
}

// Red returns red text
func (cs *ColorString) Red() *ColorString {
	return cs.getColor("\x1B[31m", "\x1B[39m", "color")
}

// Green returns green text
func (cs *ColorString) Green() *ColorString {
	return cs.getColor("\x1B[32m", "\x1B[39m", "color")
}

// Yellow returns yellow text
func (cs *ColorString) Yellow() *ColorString {
	return cs.getColor("\x1B[33m", "\x1B[39m", "color")
}

// Blue returns blue text
func (cs *ColorString) Blue() *ColorString {
	return cs.getColor("\x1B[34m", "\x1B[39m", "color")
}

// Magenta returns magenta text
func (cs *ColorString) Magenta() *ColorString {
	return cs.getColor("\x1B[35m", "\x1B[39m", "color")
}

// Cyan returns cyan text
func (cs *ColorString) Cyan() *ColorString {
	return cs.getColor("\x1B[36m", "\x1B[39m", "color")
}

// White returns white text
func (cs *ColorString) White() *ColorString {
	return cs.getColor("\x1B[37m", "\x1B[39m", "color")
}

// Grey returns grey text
func (cs *ColorString) Grey() *ColorString {
	return cs.getColor("\x1B[90m", "\x1B[39m", "color")
}

// BgS

// BlackBg returns black-Bg text
func (cs *ColorString) BlackBg() *ColorString {
	return cs.getColor("\x1B[40m", "\x1B[49m", "bg")
}

// RedBg returns red-Bg text
func (cs *ColorString) RedBg() *ColorString {
	return cs.getColor("\x1B[41m", "\x1B[49m", "bg")
}

// GreenBg returns green-Bg text
func (cs *ColorString) GreenBg() *ColorString {
	return cs.getColor("\x1B[42m", "\x1B[49m", "bg")
}

// YellowBg returns yellow-Bg text
func (cs *ColorString) YellowBg() *ColorString {
	return cs.getColor("\x1B[43m", "\x1B[49m", "bg")
}

// BlueBg returns blue-Bg text
func (cs *ColorString) BlueBg() *ColorString {
	return cs.getColor("\x1B[44m", "\x1B[49m", "bg")
}

// MagentaBg returns bagenta-Bg text
func (cs *ColorString) MagentaBg() *ColorString {
	return cs.getColor("\x1B[45m", "\x1B[49m", "bg")
}

// CyanBg returns cyan-Bg text
func (cs *ColorString) CyanBg() *ColorString {
	return cs.getColor("\x1B[46m", "\x1B[49m", "bg")
}

// WhiteBg returns white-Bg text
func (cs *ColorString) WhiteBg() *ColorString {
	return cs.getColor("\x1B[47m", "\x1B[49m", "bg")
}

// GreyBg returns grey-Bg text
func (cs *ColorString) GreyBg() *ColorString {
	return cs.getColor("\x1B[49;5;8m", "\x1B[49m", "bg")
}

// SetTextColor returns color-defined string
func (cs *ColorString) SetTextColor(color Int) *ColorString {
	var code Int
	if color >= 0 && color <= 7 {
		code = color + 30
		prefix := String("\x1B[" + string(code.ToString()) + "m")
		return cs.getColor(prefix, "\x1B[39m", "color")
	} else {
		return cs
	}

}

// Alias of SetTextColor
func (cs *ColorString) Color(color Int) *ColorString {
	return cs.SetTextColor(color)

}

// SetBgColor returns color-defined string
func (cs *ColorString) SetBgColor(color Int) *ColorString {
	var code Int
	if color >= 0 && color <= 7 {
		code = color + 40
		prefix := String("\x1B[" + string(code.ToString()) + "m")
		return cs.getColor(prefix, "\x1B[49m", "bg")
	} else {
		return cs
	}

}

// Alias of SetBgColor
func (cs *ColorString) Background(color Int) *ColorString {
	return cs.SetBgColor(color)
}

// SetAttributeColor returns string with attribute
func (cs *ColorString) SetAttribute(attr Int) *ColorString {
	prefix := String("\x1B[" + string(attr.ToString()) + "m")
	return cs.getColor(prefix, "\x1B[0m", "att")
}

// Alias of SetAttributes
func (cs *ColorString) Attribute(attr Int) *ColorString {
	return cs.SetAttribute(attr)
}

// Value returns the value typed String of ColorString
func (cs *ColorString) Value() String {
	return cs.bgColor[0] + cs.color[0] + cs.attributes[0] + cs.str + cs.attributes[1] + cs.color[1] + cs.bgColor[1]
}

// String returns string value of ColorString. It implements Stringer interface
func (cs ColorString) String() string {
	return string(cs.bgColor[0] + cs.color[0] + cs.attributes[0] + cs.str + cs.attributes[1] + cs.color[1] + cs.bgColor[1])
}
