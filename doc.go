/*
Package dna implements some basic types: string, int, float, bool, array or interface such as Item.

For the most part, it focuses on redefining some of basic parts:
	* Rewriting the most popular types: Int, String, Bool... into objects
	* Adding methods to each types.
	* Rewriting the StringArray(array of string) and IntArray(array of int)
	* Adding methods to each array type
	* Adding Log method to print on screen
	* Adding some utils: ParseInt(),Map()...

Notice: The Object-like primitive types have some methods that do not check the input.
Error maybe occurs in runtime. Therefore check the input while coding.
"Array" term used here, in fact, it is slice. But for the sake of convinience, array will be called.
*/
package dna
