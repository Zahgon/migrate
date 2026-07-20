package source

import (
	"fmt"
	"regexp"
)

var (
	ErrParse = fmt.Errorf("no match")
)

var (
	DefaultParse = Parse
	DefaultRegex = Regex
)

var Regex = regexp.MustCompile(`^([0-9]+)_(.*)\.(` + string(Down) + `|` + string(Up) + `)\.(.*)$`)

func Parse(raw string) (*Migration, error) { _ = "STUB: not implemented"; return nil, nil }
