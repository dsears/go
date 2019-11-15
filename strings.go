Putting some garbage here, hopefully this breaks the GitHub action!

package dsears

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var reAlphaNumeric *regexp.Regexp

func init() {
	var err error
	reAlphaNumeric, err = regexp.Compile("[^A-za-z0-9]+")
	if err != nil {
		panic(err)
	}
}

// ToFuzzyString : Return a normalized version of a string suitable for fuzzy matching
func ToFuzzyString(s string) string {
	s = removeAccents(s)
	s = to7Bit(s)
	s = strings.ToLower(s)
	s = reAlphaNumeric.ReplaceAllString(s, "")
	return s
}

func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, s)
	if err != nil {
		panic(err)
	}
	return output
}

// Convert to 7 bit ASCII by replacing high codepoint characters with rune numbers
func to7Bit(s string) string {
	r := []rune(s)
	a := make([]string, len(r))
	for i, v := range r {
		if v > 127 {
			a[i] = fmt.Sprintf("x%d", v)
		} else {
			a[i] = string(v)
		}
	}
	return strings.Join(a, "")
}
