package dna

import (
	"testing"
)

// Test cases

func TestString_ToInt(t *testing.T) {
	const x String = "123"
	if x.ToInt() != 123 {
		t.Errorf("%v cannot convert string to array", x)
	}
}

func TestString_Split(t *testing.T) {
	var x String = "a,b,c"
	var y StringArray = StringArray{"a", "b", "c"}
	tmp := x.Split(",")
	for i, v := range tmp {
		if v != y[i] {
			t.Errorf("%v cannot be splitted into StringArray", x)
		}
	}
}

func TestString_Contains(t *testing.T) {
	var x String = "This is a test string"
	if x.Contains("test") != true {
		t.Errorf("%v has to contain substring", x)
	}
	if x.Contains("test1") != false {
		t.Errorf("%v hasnt to contain substring", x)
	}

}

func TestString_ToLowerCase(t *testing.T) {
	var x String = "This is a test string"
	if x.ToLowerCase() != "this is a test string" {
		t.Errorf("%v cannot converted to lower case", x)
	}

}

func TestString_ToUpperCase(t *testing.T) {
	var x String = "This is a test string"
	if x.ToUpperCase() != "THIS IS A TEST STRING" {
		t.Errorf("%v cannot converted to upper case", x)
	}

}

func TestString_Trim(t *testing.T) {
	var x String = "  This is a test string  "
	if x.Trim() != "This is a test string" {
		t.Errorf("%v cannot be trimmed", x)
	}

}

func TestString_Title(t *testing.T) {
	var x String = "This is a test string"
	if x.Title() != "This Is A Test String" {
		t.Errorf("%v cannot be converted to title", x)
	}

}

func TestString_Replace(t *testing.T) {
	var x String = "This is a test string"
	if x.Replace("a test", "the second") != "This is the second string" {
		t.Errorf("%v cannot be replaced", x)
	}

}

// ------------------------------------------
// Example cases

func ExampleFromCharCode() {
	Logv(FromCharCode(7897))
	Logv(FromCharCode(32))
	Logv(FromCharCode(97))
	Logv(FromCharCode(65))
	// Output: "ộ"
	// " "
	// "a"
	// "A"
}

func ExampleString_ToInt() {
	const x String = "123"
	Logv(x.ToInt())
	// Output: 123
}

func ExampleString_Split() {
	var x String = "a,b,c"
	var y String = "ốộảờ"
	Logv(x.Split(","))
	Logv(y.Split(""))
	// Output: dna.StringArray{"a", "b", "c"}
	//dna.StringArray{"ố", "ộ", "ả", "ờ"}
}

func ExampleString_SplitWithRegexp() {
	var y String = "Một   hai     ba   bốn     năm"
	Logv(y.SplitWithRegexp("\\s+", -1))
	Logv(y.SplitWithRegexp("\\s+", 3))
	// Output: dna.StringArray{"Một", "hai", "ba", "bốn", "năm"}
	// dna.StringArray{"Một", "hai", "ba   bốn     năm"}
}

func ExampleString_Substring() {
	var x String = "Một hai ba bốn năm sáu bảy tám chín mười"
	Logv(x.Substring(4, 18))
	// Output: "hai ba bốn năm"
}

func ExampleString_Contains() {
	var x String = "a,b,c"
	var y String = "Một hai ba bốn năm sáu bảy tám chín mười"
	Logv(x.Contains(","))
	Logv(y.Contains("bốn"))
	Logv(y.Contains("tám mười"))
	// Output: true
	// true
	// false
}

func ExampleString_HasPrefix() {
	var x String = "http://google.com"
	Logv(x.HasPrefix("http"))
	Logv(x.HasPrefix("https"))
	// Output: true
	// false
}

func ExampleString_HasSuffix() {
	var x String = "http://google.com/my_pic.jpg"
	Logv(x.HasSuffix("jpg"))
	Logv(x.HasSuffix("png"))
	// Output: true
	// false
}

func ExampleString_CharAt() {
	var x String = "This is a test string"
	var y String = "Một hai ba bốn năm sáu bảy tám chín mười"
	Logv(x.CharAt(8))
	Logv(y.CharAt(1))
	Logv(y.CharAt(12))
	// Output: "a"
	// "ộ"
	// "ố"
}

func ExampleString_CharCodeAt() {
	var x String = "This is a test string"
	var y String = "Một hai ba bốn năm sáu bảy tám chín mười"
	Logv(x.CharCodeAt(8))
	// Find index of "ộ". It takes up to 3 bytes ([225 187 153]) to store that char and has unicode of 7897
	Logv(y.CharCodeAt(1))
	Logv(y.CharCodeAt(12))
	// Output: 97
	// 7897
	// 7889
}

func ExampleString_IndexOf() {
	var x String = "This is a test string"
	var y String = "Một hai ba bốn năm sáu bảy tám chín mười"
	Logv(x.IndexOf("test"))
	Logv(x.IndexOf("test1"))
	Logv(y.IndexOf("năm"))
	Logv(y.IndexOf("năm 2"))
	Logv(y.IndexOf("ộ"))
	// Output: 10
	// -1
	// 15
	// -1
	// 1
}

func ExampleString_LastIndexOf() {
	var x String = "This is a test string. Một hai ba bốn năm sáu bảy tám chín mười"
	Logv(x.LastIndexOf("test"))
	Logv(x.LastIndexOf("bốn"))
	Logv(x.LastIndexOf("bốn2"))
	// Output: 10
	// 34
	// -1
}

func ExampleString_ToLowerCase() {
	const x String = "This is a test string"
	Logv(x.ToLowerCase())
	// Output: "this is a test string"
}

func ExampleString_ToUpperCase() {
	const x String = "This is a test string"
	Logv(x.ToUpperCase())
	// Output: "THIS IS A TEST STRING"
}

func ExampleString_Title() {
	const x String = "This is a test string.Đêm đêm đốt đèn đi đâu đó?"
	Logv(x.Title())
	// Output: "This Is A Test String.Đêm Đêm Đốt Đèn Đi Đâu Đó?"
}

func ExampleString_Trim() {
	const x String = "  \t  This is a test string Đêm đêm đốt đèn đi đâu đó?  \n \t   "
	Logv(x.Trim())
	// Output: "This is a test string Đêm đêm đốt đèn đi đâu đó?"
}

func ExampleString_Replace() {
	const x String = "This is a test string. Đêm đêm đốt đèn đi đâu đó?"
	Logv(x.Replace("a test", "the second"))
	Logv(x.Replace("đâu đó?", "đãi đỗ đen"))
	// Output: "This is the second string. Đêm đêm đốt đèn đi đâu đó?"
	// "This is a test string. Đêm đêm đốt đèn đi đãi đỗ đen"
}

func ExampleString_ReplaceWithRegexp() {
	var x String = "<html>This is a text. Đêm đêm đốt đèn đi đâu đó?</html>"
	Logv(x.ReplaceWithRegexp("h..l", "div"))
	Logv(x.ReplaceWithRegexp("<(.?ht)ml>", "<${1}ml>"))
	// Output: "<div>This is a text. Đêm đêm đốt đèn đi đâu đó?</div>"
	// "<html>This is a text. Đêm đêm đốt đèn đi đâu đó?</html>"
}

func ExampleString_Length() {
	var x String = "This is a test string"
	var y String = "Đêm đêm đốt đèn đi đâu đó?"
	Logv(x.Length())
	Logv(y.Length())
	// Output: 21
	// 26
}

func ExampleString_TotalBytes() {
	var x String = "This is a test string"
	var y String = "Đêm đêm đốt đèn đi đâu đó?"
	Logv(x.TotalBytes())
	Logv(y.TotalBytes())
	// Output: 21
	// 40
}

func ExampleString_Count() {
	const x String = "This is a test string"
	Logv(x.Count("i"))
	// Output: 3
}

func ExampleString_Concat() {
	var x String = "This is a test string"
	var y String = "This is another test string"
	Logv(x.Concat(y, ". "))
	// Output: "This is a test string. This is another test string"
}

func ExampleString_FindAllString() {
	const x String = "<html>This is a text</html>"
	Logv(x.FindAllString("<.?(ht)ml>", 1))
	Logv(x.FindAllString("<.?(ht)ml>", -1))
	// Output: dna.StringArray{"<html>"}
	// dna.StringArray{"<html>", "</html>"}
}

func ExampleString_FindAllStringSubmatch() {
	const x String = "<html>This is a text</html>"
	Logv(x.FindAllStringSubmatch("<.?(ht)ml>", -1))
	// Output: []dna.StringArray{dna.StringArray{"<html>", "ht"}, dna.StringArray{"</html>", "ht"}}
}

func ExampleString_FindAllStringIndex() {
	const x String = "<html>This is a text</html>"
	Logv(x.FindAllStringIndex("<.?(ht)ml>"))
	// Output: []dna.IntArray{dna.IntArray{0, 6}, dna.IntArray{20, 27}}
}

func ExampleString_Search() {
	var x String = "<html>This is a text</html>"
	Logv(x.Search("<.?(ht)ml>"))
	Logv(x.Search("<.?(ht)mlx>"))
	// Output: 0
	// -1
}

func ExampleString_Match() {
	var x String = "<html>This is a text</html>"
	Logv(x.Match("<.?(ht)ml>"))
	// Output: true
}

func ExampleString_RemoveHtmlTags() {
	var x String = `<div class="clear-fix">This is <strong>a string</strong></div>`
	Logv(x.RemoveHtmlTags(""))
	// Output: "This is a string"
}

func ExampleString_GetTagAttributes() {
	var x String = `<link rel="stylesheet" type="text/css" href="http://google.com" media="all" />`
	Logv(x.GetTagAttributes("href"))
	// Output: "http://google.com"
}

func ExampleString_GetTags() {
	var x String = `<div class="clear-fix"><strong>This</strong> is <strong>a text</strong>.</div>`
	Logv(x.GetTags("strong"))
	Logv(x.GetTags("strong")[0].RemoveHtmlTags(""))
	// Output: dna.StringArray{"<strong>This</strong>", "<strong>a text</strong>"}
	// "This"
}

func ExampleString_EscapeHTML() {
	var x String = `<div>Blah blah blah</div>`
	Logv(x.EscapeHTML())
	// Output: "&lt;div&gt;Blah blah blah&lt;/div&gt;"
}

func ExampleString_UnescapeHTML() {
	var x String = `&lt;div&gt;Blah blah blah&lt;/div&gt;`
	Logv(x.UnescapeHTML())
	// Output: "<div>Blah blah blah</div>"
}

func ExampleString_Clean() {
	var x String = "    Đêm   đêm    đốt    đèn    đi   đâu    đó?   "
	Logv(x.Clean())
	// Output: "Đêm đêm đốt đèn đi đâu đó?"
}

func ExampleString_ToChars() {
	var x String = "Đêm đêm"
	Logv(x.ToChars())
	// Output: dna.StringArray{"Đ", "ê", "m", " ", "đ", "ê", "m"}
}

func ExampleString_ToLines() {
	var x String = "Hôm nay là thứ 2.\nNgày mai là thứ 3.\rThis is 3rd lines.\r\nThis is the last one"
	Logv(x.ToLines())
	// Output: dna.StringArray{"Hôm nay là thứ 2.", "Ngày mai là thứ 3.", "This is 3rd lines.", "This is the last one"}
}

func ExampleString_ToWords() {
	var x String = "Hôm nay là thứ 2"
	Logv(x.ToWords())
	// Output: dna.StringArray{"Hôm", "nay", "là", "thứ", "2"}
}

func ExampleString_IsBlank() {
	Logv(String("").IsBlank())
	Logv(String("\n").IsBlank())
	Logv(String(" ").IsBlank())
	Logv(String("a").IsBlank())
	// Output: true
	// true
	// true
	// false
}

func ExampleString_Reverse() {
	Logv(String("abc").Reverse())
	Logv(String("Hôm").Reverse())
	Logv(String("RACECAR").Reverse())
	// Output: "cba"
	// "môH"
	// "RACECAR"
}

func ExampleString_ToFormattedString() {
	Logv(String("Hello World!").ToFormattedString(15))
	Logv(String("Hello World!").ToFormattedString(-15))
	Logv(String("Hello World!").ToFormattedString(10))
	Logv(String("Hôm nay").ToFormattedString(15))
	Logv(String("Hôm nay").ToFormattedString(-15))
	// Output:
	// "   Hello World!"
	// "Hello World!   "
	// "Hello World!"
	// "        Hôm nay"
	// "Hôm nay        "
}

func ExampleString_Repeat() {
	Logv(String(" ").Repeat(5))
	Logv(String("Hôm nay.").Repeat(5))
	Logv(String("Hôm nay.").Repeat(0))
	// Output:
	// "     "
	// "Hôm nay.Hôm nay.Hôm nay.Hôm nay.Hôm nay."
	// ""
}
