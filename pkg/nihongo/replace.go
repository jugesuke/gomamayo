package nihongo

import (
	"regexp"
	"strings"
)

type (
	rep struct {
		old string
		new string
	}
	regRep struct {
		regexp *regexp.Regexp
		new    string
	}

	replacer interface {
		Replace(s string) string
	}
)

func (r *rep) Replace(s string) string {
	return strings.ReplaceAll(s, r.old, r.new)
}

func (r *regRep) Replace(s string) string {
	return r.regexp.ReplaceAllString(s, r.new)
}

