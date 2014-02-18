package dna

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"html"
	"regexp"
	"strings"
	"unicode/utf8"
)

//Redefine new string type.
//This convert many functions from standards `strings` libray to methods of String Type.
//
//Notice: String is not nomalized, take a rune as a Unicode code point.
//It does not handle a character as multiple runs.
//Look more at http://blog.golang.org/normalization
type String string

// ToInt returns an integer from string.
func (s String) ToInt() Int {
	return ParseInt(s, 10)
}

// Split returns a new string array from a string given a separator.
func (s String) Split(sep String) StringArray {
	ret := strings.Split(string(s), string(sep))
	result := make([]String, len(ret))
	for index, value := range ret {
		result[index] = String(value)
	}
	return result
}

// Returns substring from indexA to indexB.
func (s String) Substring(indexA, indexB Int) String {
	ra := []rune(string(s))
	return String(string(ra[indexA:indexB]))
}

// Contains checks if a string contains a substring.
func (s String) Contains(substr String) Bool {
	return Bool(strings.Contains(string(s), string(substr)))
}

// ContainsWithRegexp checks if a string contains a given pattern.Similar with Match.
func (s String) ContainsWithRegexp(pattern String) Bool {
	return s.Match(pattern)
}

// CharAt returns a char at the specific index.
func (s String) CharAt(index Int) String {
	return String([]rune(string(s))[index])
}

// CharCodeAt returns a char code at the specific index.
func (s String) CharCodeAt(index Int) Int {
	return Int([]rune(string(s))[index])
}

// HasPrefix checks if the string begins with a prefix.
func (s String) HasPrefix(prefix String) Bool {
	return Bool(strings.HasPrefix(string(s), string(prefix)))
}

// HasSuffix checks if the string ends with a suffix.
func (s String) HasSuffix(suffix String) Bool {
	return Bool(strings.HasSuffix(string(s), string(suffix)))
}

func (s String) convertBytesToRuneIndex(substr String, bytesIndex Int) Int {
	runeArray := []rune(string(s))
	var runeIndex Int = 0
	var bytesIndexCount Int = 0

	if bytesIndex == -1 {
		runeIndex = -1
	} else {
		for bytesIndexCount < bytesIndex {
			_, sizeRune := utf8.DecodeRuneInString(string(runeArray[runeIndex]))
			bytesIndexCount += Int(sizeRune)
			runeIndex += 1
		}
	}
	return runeIndex
}

// IndexOf returns an index of a substring. It returns -1 if not found.
// IndexOf supports utf8, it is diffrent from standard library (strings.Index). It returns the position of substring, not the index in byte array.
func (s String) IndexOf(substr String) Int {
	return s.convertBytesToRuneIndex(substr, Int(strings.Index(string(s), string(substr))))
}

// Alias of Search
func (s String) IndexWithRegexp(pattern String) Int {
	return s.Search(pattern)
}

// LastIndexOf returns last index of a substring. It returns -1 if not found.
// LastIndexOf supports utf8. See more at IndexOf() method.
func (s String) LastIndexOf(substr String) Int {
	return s.convertBytesToRuneIndex(substr, Int(strings.Index(string(s), string(substr))))
}

// ToLowerCase converts a string to lower case and returns a new  string.
func (s String) ToLowerCase() String {
	return String(strings.ToLower(string(s)))
}

// ToUpperCase converts a string to upper case and returns a new  string.
func (s String) ToUpperCase() String {
	return String(strings.ToUpper(string(s)))
}

// Trim removes space, newline, newfeed, return cariage at the beginning and at the end of a string.
func (s String) Trim() String {
	return String(strings.TrimSpace(string(s)))
}

// Title turns a string into title format (Capital letter at the beginning of a word).
func (s String) Title() String {
	return String(strings.Title(string(s)))
}

// Replace replaces  a substring in a string with a replacing one.
func (s String) Replace(old, repl String) String {
	return String(strings.Replace(string(s), string(old), string(repl), -1))
}

// ReplaceWithRegexp replaces string with regexp string.
// ReplaceWithRegexp returns a copy of src, replacing matches of the Regexp with the replacement string repl. Inside repl, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first submatch.
func (s String) ReplaceWithRegexp(pattern, repl String) String {
	r := regexp.MustCompile(string(pattern))
	result := r.ReplaceAllString(string(s), string(repl))
	return String(result)
}

// Length returns the the number of characters in a string. Unicode chars such as å,ø would be counted 1.
func (s String) Length() Int {
	return Int(utf8.RuneCountInString(string(s)))
}

// TotalBytes returns the total bytes of string. EX: "ộ" will take 3 bytes.
func (s String) TotalBytes() Int {
	return Int(len(s))
}

// Count returns the number of a specific string
func (s String) Count(sep String) Int {
	return Int(strings.Count(fmt.Sprint(s), fmt.Sprint(sep)))
}

// FindAllString returns all string given a passed regular expression.
func (s String) FindAllString(pattern String, n Int) StringArray {
	r := regexp.MustCompile(string(pattern))
	matchedStrings := r.FindAllString(string(s), int(n))
	return NewStringArray(matchedStrings)
}

// FindAllStringSubmatch returns all substring with regexp.
func (s String) FindAllStringSubmatch(pattern String, n Int) []StringArray {
	r := regexp.MustCompile(string(pattern))
	matchedStrings := r.FindAllStringSubmatch(string(s), int(n))
	result := make([]StringArray, len(matchedStrings))
	for i, value := range matchedStrings {
		result[i] = NewStringArray(value)
	}
	return result
}

// FindAllStringIndex returns all indexes of substring by regexp.
func (s String) FindAllStringIndex(pattern String) []IntArray {
	r := regexp.MustCompile(string(pattern))
	matchedIndexes := r.FindAllStringIndex(string(s), -1)
	result := make([]IntArray, len(matchedIndexes))
	for i, value := range matchedIndexes {
		result[i] = NewIntArray(value)
	}
	return result
}

// Search returns the leftmost index of matched string given the pattern passed.
func (s String) Search(pattern String) Int {
	r := regexp.MustCompile(string(pattern))
	matchedIndex := r.FindStringIndex(string(s))
	if matchedIndex == nil {
		return -1
	} else {
		return Int(matchedIndex[0])
	}
}

// Match returns true if patterns matches the string
func (s String) Match(pattern String) Bool {
	r := regexp.MustCompile(string(pattern))
	matched := r.MatchString(string(s))
	return Bool(matched)
}

// SplitWithRegexp returns n elements in StringArray splitted by regexp pattern. The count determines the number of substrings to return:
//
// 		n > 0: at most n substrings; the last substring will be the unsplit remainder.
// 		n == 0: the result is nil (zero substrings)
// 		n < 0: all substrings
func (s String) SplitWithRegexp(pattern String, n Int) StringArray {
	r := regexp.MustCompile(string(pattern))
	matchedStrings := r.Split(string(s), int(n))
	return NewStringArray(matchedStrings)
}

// Concat returns a new concatenated string from target string and seperator.
func (s String) Concat(target String, sep String) String {
	return String(fmt.Sprintf("%v%v%v", s, sep, target))
}

// RemoveHtmlTags returns all tags in HTML-like string. If tag param is empty ("") then it will remove all tags.
func (s String) RemoveHtmlTags(tag String) String {
	return s.ReplaceWithRegexp(String(fmt.Sprintf("</?%v[^<>]*>", tag)), "")

}

// GetTagAttributes returns only one attribute per tag. HTML parse by using regexp is not sufficient but in some cases, it's better to use them.
func (s String) GetTagAttributes(att String) String {
	ret := s.FindAllStringSubmatch(String(fmt.Sprintf("%v\\s*=[\"']([^\"']+?)[\"']", att)), -1)
	if len(ret) > 0 {
		return ret[0][1]
	} else {
		return ""
	}
}

// GetTags returns tags of HTML-like string. NOTE: this method does not resolve the nested tags.
func (s String) GetTags(tag String) StringArray {
	return s.FindAllString(String(fmt.Sprintf("(?mis)<%v[^>]*>?(.*?)(<\\/%v>|\\/>)", tag, tag)), -1)

}

// EscapeHTML escapes special characters like "<" to become "&lt;". It escapes only five such characters: <, >, &, ' and ".
func (s String) EscapeHTML() String {
	return String(html.EscapeString(string(s)))
}

//UnescapeHTML unescapes entities like "&lt;" to become "<". It unescapes a larger range of entities than EscapeString escapes. For example, "&aacute;" unescapes to "á", as does "&#225;" and "&xE1;". UnescapeString(EscapeString(s)) == s always holds, but the converse isn't always true.
func (s String) UnescapeHTML() String {
	return String(html.UnescapeString(string(s)))
}

// OTHER MISC STRING UTILITIES.

// Clean returns a new string whose whitespaces are compressed.
func (s String) Clean() String {
	return s.Trim().ReplaceWithRegexp("\\s+", " ")
}

// ToChars returns an array of characters.
func (s String) ToChars() StringArray {
	return s.Split("")
}

// ToLines returns an array of lines which are seperated by Unix format (linefeed "\n"), Windows ("\r\n") or Mac ("\r").
func (s String) ToLines() StringArray {
	return s.SplitWithRegexp("\r\n|\r|\n", -1)
}

// ToWords returns an array of words which are seperated by "\s+".
func (s String) ToWords() StringArray {
	return s.SplitWithRegexp("\\s+", -1)
}

// IsBlank returns true if the string is blank.
func (s String) IsBlank() Bool {
	return s.Match("^\\s*$")
}

// ToStringArray returns a string array.
func (s String) ToStringArray() StringArray {
	return StringArray{s}
}

// Reverse returns a new reversed string.
func (s String) Reverse() String {
	return s.Split("").Reverse().Join("")
}

// Alias of HasPrefix.
func (s String) StartsWith(prefix String) Bool {
	return s.HasPrefix(prefix)
}

// Alias of HasSuffix.
func (s String) EndsWith(suffix String) Bool {
	return s.HasSuffix(suffix)
}

// Alias of Trim.
func (s String) Strip() String {
	return s.Trim()
}

// Value implements the Valuer interface in database/sql/driver package.
func (s String) Value() (driver.Value, error) {
	return driver.Value(string(s)), nil
}

// Scan implements the Scanner interface in database/sql package.
// Dedault value for nil is 0
func (s *String) Scan(src interface{}) error {
	var source String
	switch src.(type) {
	case string:
		source = String(src.(string))
	case []byte:
		source = String(string(src.([]byte)))
	case nil:
		source = ""
	default:
		return errors.New("Incompatible type for dna.String type")
	}
	*s = source
	return nil
}

// ToPrimitiveValue returns a string's primitive type.
func (s String) ToPrimitiveValue() string {
	return string(s)
}

// Value returns a string's primitive type.
func (s String) String() string {
	return string(s)
}

// ToFormattedString returns string width constant width.
// If length of string less than its width param, it will fills the string with whitespace.
// Otherwise, it will keep its original
func (s String) ToFormattedString(width Int) String {
	return String((fmt.Sprintf("%[2]*[1]s", s, (int(width)))))
}

// Repeat returns a new string repeated n times.
// If n <= 0, then return empty string.
func (s String) Repeat(count Int) String {
	if count > 0 {
		return String(strings.Repeat(string(s), int(count)))
	} else {
		return ""
	}

}

//Underscore returns converted camel cased string into a string delimited by underscores.
func (s String) Underscore() String {
	return s.ReplaceWithRegexp(`([a-z\d])([A-Z]+)`, `${1}_${2}`).ReplaceWithRegexp(`[-\s]+`, `_`).ToLowerCase()
}

// Alias of Underscore()
func (s String) ToSnakeCase() String {
	return s.Underscore()
}

// Camelize return new string which removes any underscores or dashes and convert a string into camel casing.
func (s String) Camelize() String {
	r := regexp.MustCompile(`(\-|_|\s)+(.)?`)
	result := r.ReplaceAllStringFunc(string(s), func(val string) string {
		return String(val).ToUpperCase().ReplaceWithRegexp(`(\-|_|\s)+`, ``).ToPrimitiveValue()
	})
	return String(result)
}

//Alias of Camelize()
func (s String) ToCamelCase() String {
	return s.Camelize()
}

// dasherize returns  a converted camel cased string into a string delimited by dashes.
func (s String) Dasherize() String {
	return s.ReplaceWithRegexp(`([a-z\d])([A-Z]+)`, `${1}-${2}`).ReplaceWithRegexp(`[_\s]+`, `-`).ToLowerCase()
}

// Alias of Dasherize().
func (s String) ToDashCase() String {
	return s.Dasherize()
}

// Alias of UnescapeHTML().
func (s String) DecodeHTML() String {
	return s.UnescapeHTML()
}

// IsEmpty returns true if a string is empty.
func (s String) IsEmpty() Bool {
	return s == ""
}

// IsLower returns true if a string contains all chars with lower case.
func (s String) IsLower() Bool {
	return s == s.ToLowerCase()
}

// IsLower returns true if a string contains all chars with upper case.
func (s String) IsUpper() Bool {
	return s == s.ToUpperCase()
}

// IsNumeric returns true if a string contains only digits from 0-9. Other digits not in Latin (such as Arabic) are not currently supported.
func (s String) IsNumeric() Bool {
	return s.Match(`^[0-9]+$`)
}

// IsAlpha returns true if a string contains only letters from ASCII (a-z,A-Z). Other letters from other languages are not supported.
func (s String) IsAlpha() Bool {
	return s.Match(`^[a-zA-Z]+$`)
}

// IsAlphaNumeric returns true if a string contains letters and digits.
func (s String) IsAlphaNumeric() Bool {
	return s.Match(`^[a-zA-Z0-9]+$`)
}

// Between returns a string between left and right strings.
// If multiple results found, get only the first in non-greedy method
func (s String) Between(left, right String) String {
	ret := s.FindAllStringSubmatch(left.Concat("(.+?)", "").Concat(right, ""), 1)
	if len(ret) > 0 && ret[0].Length() > 1 {
		return ret[0][1]
	} else {
		return ""
	}
}

// ToBytes returns a byte array from a string.
func (s String) ToBytes() []byte {
	return []byte(s.String())
}
