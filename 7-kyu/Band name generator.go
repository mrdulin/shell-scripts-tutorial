/*
My friend wants a new band name for her band. She like bands that use the formula: "The" + a noun with the first letter capitalized, for example:

"dolphin" -> "The Dolphin"

However, when a noun STARTS and ENDS with the same letter, she likes to repeat the noun twice and connect them together with the first and last letter, combined into one word (WITHOUT "The" in front), like this:

"alaska" -> "Alaskalaska"

Complete the function that takes a noun as a string, and returns her preferred band name written as a string.

*/
package kata

import "strings"

func bandNameGenerator(word string) string {
  s := word[0]
  e := word[len(word) - 1]
  var bandName string
  if s == e {
    bandName = strings.Title(word) + word[1:]
  } else {
    bandName = "The " + strings.Title(word)
  }
  return bandName
}