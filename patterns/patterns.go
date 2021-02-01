package patterns

import (
	"github.com/yuin/gopher-lua/parse"
	"strings"
	"../beautifier"
)

func generatePatterns() (p, error) {
	for name, str := range patterns {
		ast, err := parse.Parse(strings.NewReader(str), "")
		if err != nil {
			return nil, err
		}
		out := append(out, beautifier.GeneratePattern(chunk))
	}
	return p, nil
}