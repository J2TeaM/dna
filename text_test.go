package dna

func ExampleText_Replace() {
	var x Text = "This is a test string"
	Logv(x.Replace("a test", "the second"))
	// Output: "This is the second string"
}

func ExampleText_ReplaceWithRegexp() {
	var x Text = "<html>This is a text</html>"
	Logv(x.ReplaceWithRegexp("h..l", "div"))
	Logv(x.ReplaceWithRegexp("<(.?ht)ml>", "<${1}ml>"))
	// Output: "<div>This is a text</div>"
	// "<html>This is a text</html>"
}

func ExampleText_FindAllString() {
	var x Text = "<html>This is a text</html>"
	Logv(x.FindAllString("<.?(ht)ml>", 1))
	Logv(x.FindAllString("<.?(ht)ml>", -1))
	// Output:
	// dna.StringArray{"<html>"}
	// dna.StringArray{"<html>", "</html>"}
}

func ExampleText_Match() {
	var x Text = "<html>This is a text</html>"
	Logv(x.Match("<.?(ht)ml>"))
	// Output: true
}

func ExampleText_FindAllStringSubmatch() {
	var x Text = "<html>This is a text</html>"
	Logv(x.FindAllStringSubmatch("<.?(ht)ml>", -1))
	// Output: []dna.StringArray{dna.StringArray{"<html>", "ht"}, dna.StringArray{"</html>", "ht"}}
}

// HTML parts

func ExampleText_RemoveHtmlTags() {
	var x Text = `<div class="clear-fix">This is <strong>a text</strong></div>`
	Logv(x.RemoveHtmlTags(""))
	// Output: "This is a text"
}

func ExampleText_GetTagAttributes() {
	var x Text = `<link rel="stylesheet" type="text/css" href="http://google.com" media="all" />`
	Logv(x.GetTagAttributes("href"))
	// Output: "http://google.com"
}

func ExampleText_GetTags() {
	var x Text = `<div class="clear-fix"><strong>This</strong> is <strong>a text</strong>.</div>`
	Logv(x.GetTags("strong"))
	Logv(x.GetTags("strong")[0].RemoveHtmlTags(""))
	// Output: dna.StringArray{"<strong>This</strong>", "<strong>a text</strong>"}
	// "This"
}
