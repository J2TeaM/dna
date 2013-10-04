package dna

import (
	"fmt"
	"regexp"
	"strings"
)

//Redefine new text type.
//The difference between String and Text is that Text is passed by pointer to its methods
type Text string

// Replace  a substring in a text with a replacing one
func (t *Text) Replace(old, repl String) String {
	return String(strings.Replace(string(*t), string(old), string(repl), -1))
}

// Replace string with regexp string.
// ReplaceWithRegexp returns a copy of src, replacing matches of the Regexp with the replacement string repl. Inside repl, $ signs are interpreted as in Expand, so for instance $1 represents the text of the first submatch.
func (t *Text) ReplaceWithRegexp(pattern, repl String) String {
	r := regexp.MustCompile(string(pattern))
	result := r.ReplaceAllString(string(*t), string(repl))
	return String(result)
}

// Find all strings matched in terms of pattern. Return n maches
func (t *Text) FindAllString(pattern String, n Int) StringArray {
	r := regexp.MustCompile(string(pattern))
	matchedStrings := r.FindAllString(string(*t), int(n))
	return NewStringArray(matchedStrings)
}

// Return if patterns matches the text
func (t *Text) Match(pattern String) Bool {
	r := regexp.MustCompile(string(pattern))
	matched := r.MatchString(string(*t))
	return Bool(matched)
}

// Find all substring with regexp
func (t *Text) FindAllStringSubmatch(pattern String, n Int) []StringArray {
	r := regexp.MustCompile(string(pattern))
	matchedStrings := r.FindAllStringSubmatch(string(*t), int(n))
	result := make([]StringArray, len(matchedStrings))
	for i, value := range matchedStrings {
		result[i] = NewStringArray(value)
	}
	return result
}

// Remove all tags in HTML-like text. If tag param is empty ("") then it will remove all tags
func (t *Text) RemoveHtmlTags(tag String) String {
	return t.ReplaceWithRegexp(String(fmt.Sprintf("</?%v[^<>]*>", tag)), "")

}

// Get only one attribute per tag. HTML parse by using regexp is not sufficient but in some cases, it's better to use them
func (t *Text) GetTagAttributes(att String) String {
	return t.FindAllStringSubmatch(String(fmt.Sprintf("%v\\s*=[\"']([^\"']+?)[\"']", att)), -1)[0][1]

}

// Get tags of HTML-like text. NOTE: this method does not resolve the nested tags
func (t *Text) GetTags(tag String) StringArray {
	return t.FindAllString(String(fmt.Sprintf("<%v[^>]*>?(.*?)(<\\/%v>|\\/>)", tag, tag)), -1)

}
