package parser

import (
	"strconv"
	"strings"
	"unicode"
)

// func alphaNumericFields(r rune) bool {
// 	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
// }

func regularFields(r rune) bool {
	return r == ',' || r == ';' || r == '/' || r == '(' || r == ')' || r == '[' || r == ']' || r == '&'
}

func regularFieldsWithDelimiters(delimiters []rune) func(rune) bool {
	return func(r rune) bool {
		for _, delimiter := range delimiters {
			if r == delimiter {
				return true
			}
		}
		return false
	}
}

func numericFields(r rune) bool {
	return !unicode.IsNumber(r)
}

// SplitIntoRegularFields includes all alphanumeric characters.
func SplitIntoRegularFields(s string) []string {
	flds := strings.FieldsFunc(s, regularFields)
	for i := 0; i < len(flds); i++ {
		flds[i] = strings.TrimSpace(flds[i])
	}
	return flds
}

// SplitIntoRegularFieldsWithDelimiters includes all alphanumeric characters for custom delimiters.
func SplitIntoRegularFieldsWithDelimiters(s string, delimiters []rune) []string {
	flds := strings.FieldsFunc(s, regularFieldsWithDelimiters(delimiters))
	for i := 0; i < len(flds); i++ {
		flds[i] = strings.TrimSpace(flds[i])
	}
	return flds
}

// SplitIntoNumericFields includes numeric values only.
func SplitIntoNumericFields(s string) []string {
	return strings.FieldsFunc(s, numericFields)
}

// HasOneOfPrefixes searches prefixes in an input string.
func HasOneOfPrefixes(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// NaiveStringToInt преобразует строку в целочисленное значение без ошибки.
// Если значение не удается преобразовать, возвращается нулевое значение.
func NaiveStringToInt(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return ret
}
