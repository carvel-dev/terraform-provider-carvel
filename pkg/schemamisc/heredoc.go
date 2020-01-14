package schemamisc

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	indentRegexp = regexp.MustCompile(`^\s+`)
)

type Heredoc struct {
	Data string
}

func (h Heredoc) StripIndent() (string, error) {
	pieces := strings.Split(strings.TrimRight(h.Data, " \t\n"), "\n")
	if len(pieces) > 1 {
		firstIndent := indentRegexp.FindString(pieces[0])
		for i, piece := range pieces {
			if !strings.HasPrefix(piece, firstIndent) {
				return "", fmt.Errorf("Expected consistent heredoc indent at line %d", i+1)
			}
			pieces[i] = strings.TrimPrefix(piece, firstIndent)
		}
	}
	return strings.Join(pieces, "\n"), nil
}
