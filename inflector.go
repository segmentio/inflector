/*
 * Inflector Pkg (Go)
 *
 * Copyright (c) 2013 Ivan Torres
 * Released under the MIT license
 * https://github.com/mexpolk/inflector/blob/master/LICENSE
 *
 */

package inflector

import (
	"regexp"
	"strings"
)

// Prepare replacer and Regexps used in split
var (
	repl *strings.Replacer
	rex1 *regexp.Regexp
	rex2 *regexp.Regexp
)

func init() {
	repl = strings.NewReplacer("-", " ", "_", " ")
	rex1 = regexp.MustCompile("([A-Z]+)")
	rex2 = regexp.MustCompile("(\\w+)")
}

// Shortcut to strings.ToLower()
func lower(str string) string {
	return strings.ToLower(trim(str))
}

// Prepares strings by splitting by caps, spaces, dashes, and underscore
func split(str string) (words []string) {
	str = trim(str)

	// Convert dash and underscore to spaces
	str = repl.Replace(str)

	// Split when uppercase is found (needed for Snake)
	str = rex1.ReplaceAllString(str, " $1")

	// Get the final list of words
	words = rex2.FindAllString(str, -1)

	return
}

// Inflects to camelCase
func ToCamel(str string) (out string) {
	words := split(str)
	out = words[0]

	for _, w := range split(str)[1:] {
		out += upper(w[:1]) + w[1:]
	}

	return
}

// Inflects to kebab-case
func ToDash(str string) (out string) {
	words := split(str)

	for i, w := range words {
		words[i] = lower(w)
	}

	out = strings.Join(words, "-")
	return
}

// Inflects to PascalCase
func ToPascal(str string) (out string) {
	out = ""

	for _, w := range split(str) {
		out += upper(w[:1]) + w[1:]
	}

	return
}

// Inflects to snake_case
func ToUnderscore(str string) (out string) {
	words := split(str)

	for i, w := range words {
		words[i] = lower(w)
	}

	out = strings.Join(words, "_")
	return
}

// Removes leading whitespaces
func trim(str string) string {
	return strings.Trim(str, " ")
}

// Shortcut to strings.ToUpper()
func upper(str string) string {
	return strings.ToUpper(trim(str))
}
