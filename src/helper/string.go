package helper

import (
	"strings"

	"github.com/mozillazg/go-unidecode"
)

func NormalizeUnicodeString(s string) string {
	return strings.ToLower(unidecode.Unidecode(s))
}
