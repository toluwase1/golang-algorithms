package main

import (
	"fmt"
	"strings"
)

func main() {
abbr:=AbbrevName("Toluwase T")
fmt.Println(abbr)
}

func AbbrevName(name string) string{
	splitted := strings.Split(name, " ")// split string into 2 (becomes array)
	//strings.Index("", splitted)
	first:= string(splitted[0][0]) // pick splitted 1
	a:=strings.ToUpper(first) // convert splitted one to upper case
	//append(first, ".")
	second := string(splitted[1][0]) //pick splitted 2
	b:=strings.ToUpper(second) // convert to upper case
	s := a + "." + b
	return s
}

func AbbrevName1(name string) string{
	s := strings.Split(name, " ")
	return strings.ToUpper(string(s[0][0])) + "." + strings.ToUpper(string(s[1][0])) // Note 1
}

/**
Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them.

It should look like this:

Sam Harris => S.H

Patrick Feeney => P.F
 */